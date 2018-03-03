package iko

import (
	"errors"
	"fmt"
	"github.com/kittycash/wallet/src/util"
	"github.com/skycoin/cxo/node"
	"github.com/skycoin/cxo/skyobject"
	"github.com/skycoin/cxo/skyobject/registry"
	"github.com/skycoin/skycoin/src/cipher"
	"github.com/skycoin/skycoin/src/cipher/encoder"
	"gopkg.in/sirupsen/logrus.v1"
	"sync"
)

type MeasureChain func() uint64

type CXOStore struct {
	Meta []byte
	Txs  registry.Refs `skyobject:"schema:iko.Transaction"`
}

var (
	cxoReg = registry.NewRegistry(func(r *registry.Reg) {
		r.Register("iko.Transaction", Transaction{})
		r.Register("iko.Store", CXOStore{})
	})
)

type CXOChainConfig struct {
	Dir                string
	Public             bool
	MessengerAddresses []string
	CXOAddress         string
	CXORPCAddress      string

	Master      bool
	MasterPK    cipher.PubKey
	MasterSK    cipher.SecKey
	MasterNonce uint64 // Public
}

func (c *CXOChainConfig) Process() error {
	if e := c.MasterPK.Verify(); e != nil {
		return e
	}

	if c.Master {
		if e := c.MasterSK.Verify(); e != nil {
			return e
		}
		if c.MasterPK != cipher.PubKeyFromSecKey(c.MasterSK) {
			return errors.New("public and secret keys do not match")
		}
	}

	return nil
}

type CXOChain struct {
	mux      sync.Mutex
	c        *CXOChainConfig
	l        *logrus.Logger
	node     *node.Node
	wg       sync.WaitGroup
	received chan *Transaction
	accepted chan *Transaction

	len util.SafeInt
}

func NewCXOChain(txChecker TxChecker, config *CXOChainConfig) (*CXOChain, error) {
	if e := config.Process(); e != nil {
		return nil, e
	}

	chain := &CXOChain{
		c:        config,
		l:        logrus.New(),
		received: make(chan *Transaction),
		accepted: make(chan *Transaction),
	}

	if e := prepareNode(chain); e != nil {
		return nil, e
	}

	if e := chain.runTxService(txChecker); e != nil {
		return nil, e
	}

	return chain, nil
}

/*
	<<< PREP AND SERVICE FUNCTIONS >>>
*/

func (c *CXOChain) lock() func() {
	c.mux.Lock()
	return c.mux.Unlock
}

func prepareNode(chain *CXOChain) error {

	nc := node.NewConfig()

	nc.DataDir = chain.c.Dir
	nc.Public = chain.c.Public

	nc.TCP.Listen = chain.c.CXOAddress
	nc.TCP.Discovery = node.Addresses(chain.c.MessengerAddresses)

	nc.RPC = chain.c.CXORPCAddress

	nc.OnRootReceived = func(c *node.Conn, r *registry.Root) error {
		defer chain.lock()()

		switch {
		case r.Pub != chain.c.MasterPK:
			e := errors.New("received root is not of master public key")
			chain.l.
				WithField("master_pk", chain.c.MasterPK.Hex()).
				WithField("received_pk", r.Pub.Hex()).
				Warning(e.Error())
			return e

		case r.Nonce != chain.c.MasterNonce:
			e := errors.New("received root is not of master nonce")
			chain.l.
				WithField("master_nonce", chain.c.MasterNonce).
				WithField("received_nonce", r.Nonce).
				Warning(e.Error())
			return e

		case len(r.Refs) <= 0:
			e := errors.New("empty refs")
			chain.l.Warning(e.Error())
			return e

		default:
			return nil

		}
	}

	nc.OnRootFilled = func(n *node.Node, r *registry.Root) {
		defer chain.lock()()

		e := func(c *CXOChain, n *node.Node, r *registry.Root) error {
			var store = new(CXOStore)

			p, e := n.Container().Pack(r, cxoReg)
			if e != nil {
				return e
			}

			if e := r.Refs[0].Value(p, store); e != nil {
				return e
			}

			rLen, e := store.Txs.Len(p)
			if e != nil {
				return e
			}

			switch {
			case rLen < c.len.Val():
				return errors.New("received new root has less transactions")

			case rLen == c.len.Val():
				c.l.Info("received new root has no new transactions")
				return nil
			}

			for i := c.len.Val(); i < rLen; i++ {
				var tx = new(Transaction)
				hash, e := store.Txs.ValueByIndex(p, int(i), tx)
				if e != nil {
					return e
				}
				c.l.
					WithField("tx_hash", hash.Hex()).
					WithField("tx_seq", i).
					Info("received new transaction")
				c.received <- tx
			}

			return nil

		}(chain, n, r)

		if e != nil {
			chain.l.Error(e.Error())
			return
		}
	}

	nc.OnConnect = func(c *node.Conn) error {
		// TODO: implement.
		return nil
	}

	nc.OnDisconnect = func(c *node.Conn, reason error) {
		// TODO: implement.
	}

	var e error
	if chain.node, e = node.NewNode(nc); e != nil {
		return e
	}

	if e := chain.node.Share(chain.c.MasterPK); e != nil {
		return e
	}

	return nil
}

func (c *CXOChain) runTxService(txChecker TxChecker) error {
	c.wg.Add(1)
	defer c.wg.Done()

	go func() {
		for {
			select {
			case tx, ok := <-c.received:
				if !ok {
					return

				} else if e := txChecker(tx); e != nil {
					c.l.Warning(e.Error())

				} else {
					c.len.Set(int(tx.Seq) + 1)
					c.accepted <- tx
				}
			}
		}
	}()
	return nil
}

/*
	<<< PUBLIC FUNCTIONS >>>
*/

func (c *CXOChain) InitChain(sk cipher.SecKey) error {
	defer c.lock()()

	up, e := cxoUnpack(c)
	if e != nil {
		return e
	}

	s := new(CXOStore)
	sReg, e := cxoNewStore(up, s)
	if e != nil {
		return e
	}

	r := &registry.Root{
		Refs:  []registry.Dynamic{sReg},
		Reg:   cxoReg.Reference(),
		Pub:   c.c.MasterPK,
		Nonce: c.c.MasterNonce,
	}

	if e := c.node.Container().Save(up, r); e != nil {
		return e
	}

	c.node.Publish(r)

	return nil
}

func (c *CXOChain) Head() (Transaction, error) {
	defer c.lock()()
	var tx Transaction

	store, _, p, e := c.getStore(gsRead)
	if e != nil {
		return tx, e
	}
	if _, e := store.Txs.ValueByIndex(p, c.len.Val(), &tx); e != nil {
		return Transaction{}, e
	}
	return tx, nil
}

func (c *CXOChain) Len() uint64 {
	defer c.lock()()
	return uint64(c.len.Val())
}

func (c *CXOChain) AddTx(tx Transaction, check TxChecker) error {
	if c.c.Master == false {
		return errors.New("not master node")
	}
	if e := check(&tx); e != nil {
		return e
	}

	defer c.lock()()

	store, r, up, e := c.getStore(gsWrite)
	if e != nil {
		return e
	}
	if e := store.Txs.AppendValues(up, tx); e != nil {
		return e
	}
	if e := r.Refs[0].SetValue(up, store); e != nil {
		return e
	}
	if e := c.node.Container().Save(up.(*skyobject.Unpack), r); e != nil {
		return e
	}
	c.node.Publish(r)
	return nil
}

func (c *CXOChain) GetTxOfHash(hash TxHash) (Transaction, error) {
	defer c.lock()()
	var tx Transaction

	store, _, p, e := c.getStore(gsRead)
	if e != nil {
		return tx, e
	}
	if e := store.Txs.ValueByHash(p, cipher.SHA256(hash), &tx); e != nil {
		return tx, e
	}
	return tx, nil
}

func (c *CXOChain) GetTxOfSeq(seq uint64) (Transaction, error) {
	defer c.lock()()
	var tx Transaction

	store, _, p, e := c.getStore(gsRead)
	if e != nil {
		return tx, e
	}
	if _, e := store.Txs.ValueByIndex(p, int(seq), &tx); e != nil {
		return tx, e
	}
	return tx, nil
}

func (c *CXOChain) TxChan() <-chan *Transaction {
	return c.accepted
}

func (c *CXOChain) GetTxsOfSeqRange(startSeq uint64, pageSize uint64) ([]Transaction, error) {
	defer c.lock()()
	var txs []Transaction

	if pageSize == 0 {
		return txs, fmt.Errorf("invalid pageSize: %d", pageSize)
	}
	cLen := uint64(c.len.Val())
	if startSeq >= cLen {
		return txs, fmt.Errorf("invalid startSeq: %d", startSeq)
	}
	store, _, p, e := c.getStore(gsRead)
	if e != nil {
		return txs, e
	}
	refs, e := store.Txs.Slice(p, int(startSeq), int(startSeq+pageSize))
	if e != nil {
		return txs, e
	}
	refsLen, e := refs.Len(p)
	if e != nil {
		return txs, e
	}
	txs = make([]Transaction, refsLen)
	e = refs.Ascend(p, func(i int, hash cipher.SHA256) error {
		raw, _, e := c.node.Container().Get(hash, 0)
		if e != nil {
			return e
		}
		return encoder.DeserializeRaw(raw, &txs[i])
	})
	return txs, e
}

type getStoreType int

const (
	gsRead  getStoreType = iota
	gsWrite getStoreType = iota
)

func (c *CXOChain) getStore(t getStoreType) (*CXOStore, *registry.Root, registry.Pack, error) {
	r, e := cxoRoot(c)
	if e != nil {
		return nil, nil, nil, e
	}
	var p registry.Pack
	switch t {
	case gsRead:
		if p, e = cxoPack(c, r); e != nil {
			return nil, nil, nil, e
		}
	case gsWrite:
		if p, e = cxoUnpack(c); e != nil {
			return nil, nil, nil, e
		}
	default:
		panic("invalid getStoreType")
	}
	store, e := cxoGetStore(r, p)
	if e != nil {
		return nil, nil, nil, e
	}
	return store, r, p, nil
}

/*
	<<< HELPER FUNCTIONS >>>
*/

func cxoRoot(c *CXOChain) (*registry.Root, error) {
	r, e := c.node.Container().LastRoot(c.c.MasterPK, c.c.MasterNonce)
	if e != nil {
		return nil, e
	}
	return r, nil
}

func cxoGetStore(r *registry.Root, p registry.Pack) (*CXOStore, error) {
	if len(r.Refs) < 1 {
		return nil, errors.New("corrupt root, invalid ref count")
	}
	store := new(CXOStore)
	if e := r.Refs[0].Value(p, store); e != nil {
		return nil, e
	}
	return store, nil
}

func cxoPack(c *CXOChain, r *registry.Root) (*skyobject.Pack, error) {
	p, e := c.node.Container().Pack(r, cxoReg)
	if e != nil {
		return nil, e
	}

	return p, nil
}

func cxoUnpack(c *CXOChain) (*skyobject.Unpack, error) {
	if c.c.Master == false {
		return nil, errors.New("not master")
	}
	return c.node.Container().Unpack(c.c.MasterSK, cxoReg)
}

func cxoNewStore(up *skyobject.Unpack, s *CXOStore) (registry.Dynamic, error) {
	raw := encoder.Serialize(s)
	hash, e := up.Add(raw)
	if e != nil {
		return registry.Dynamic{}, e
	}
	schema, e := up.Registry().SchemaByName("iko.Store")
	if e != nil {
		return registry.Dynamic{}, e
	}
	return registry.Dynamic{
		Hash:   hash,
		Schema: schema.Reference(),
	}, nil
}
