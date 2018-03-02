package iko

import (
	"github.com/skycoin/cxo/node"
	"github.com/skycoin/cxo/skyobject/registry"
	"github.com/skycoin/skycoin/src/cipher"
	"gopkg.in/sirupsen/logrus.v1"
	"sync"
)



var (
	cxoReg = registry.NewRegistry(func(r *registry.Reg) {
		r.Register("iko.Transaction", Transaction{})

	})
)

type CXOChainConfig struct {
	Dir                string
	Public             bool
	MasterPK           cipher.PubKey
	NodePK             cipher.PubKey
	NodeSK             cipher.SecKey
	MessengerAddresses []string
	CXOAddress         string
	CXORPCAddress      string
}

type CXOChain struct {
	mux      sync.Mutex
	c        *CXOChainConfig
	l        *logrus.Logger
	node     *node.Node
	wg       sync.WaitGroup
	received chan *Transaction
	accepted chan *Transaction
}

func NewCXOChain(txChecker TxChecker, config *CXOChainConfig) (*CXOChain, error) {
	chain := &CXOChain{
		c:        config,
		l:        logrus.New(),
		received: make(chan *Transaction),
		accepted: make(chan *Transaction),
	}

	if e := prepareNode(chain); e != nil {
		return nil, e
	}

	return chain, nil
}

func prepareNode(chain *CXOChain) error {

	nc := node.NewConfig()

	nc.DataDir = chain.c.Dir
	nc.Public = chain.c.Public

	nc.TCP.Listen = chain.c.CXOAddress
	nc.TCP.Discovery = node.Addresses(chain.c.MessengerAddresses)

	nc.RPC = chain.c.CXORPCAddress

	nc.OnRootReceived = func(c *node.Conn, r *registry.Root) error {
		// TODO: implement.
		return nil
	}

	nc.OnRootFilled = func(n *node.Node, r *registry.Root) {
		// TODO: push new tx to 'received chan'.
	}

	nc.OnConnect = func(c *node.Conn) error {
		// TODO: implement.
		return nil
	}

	nc.OnDisconnect = func(c *node.Conn, reason error) {
		// TODO: implement.
	}

	var e error
	chain.node, e = node.NewNode(nc)
	return e
}
