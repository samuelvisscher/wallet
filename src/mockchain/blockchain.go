package mockchain

import (
	"github.com/skycoin/skycoin/src/cipher"
	"sync"
	"errors"
)

type BlockChainConfig struct {
	CreatorPK cipher.PubKey
	PK        cipher.PubKey
	SK        cipher.SecKey
	TxAction  TxAction
}

func (cc *BlockChainConfig) Prepare() error {
	if cc.TxAction == nil {
		cc.TxAction = func(tx *Transaction) error {
			return nil
		}
	}
	if e := cc.SK.Verify(); e != nil {
		return e
	}
	if cipher.PubKeyFromSecKey(cc.SK) != cc.PK {
		return errors.New("public and secret key does not match")
	}
	return nil
}

type BlockChain struct {
	c     *BlockChainConfig
	chain ChainDB
	state StateDB
	mux   sync.RWMutex

	wg   sync.WaitGroup
	quit chan struct{}
}

func NewBlockChain(config *BlockChainConfig, chainDB ChainDB, stateDB StateDB) (*BlockChain, error) {
	bc := &BlockChain{
		c:     config,
		chain: chainDB,
		state: stateDB,
		quit: make(chan struct{}),
	}

	if e := bc.InitState(); e != nil {
		return nil, e
	}

	bc.wg.Add(1)
	go bc.service()

	return bc, nil
}

func (bc *BlockChain) InitState() error {
	var prev *Transaction

	for i := uint64(1); i < bc.chain.Len(); i++ {

		// Get transaction.
		tx, e := bc.chain.GetTxOfSeq(i)
		if e != nil {
			return e
		}

		// Check hash, seq and sig of tx.
		if e := tx.Verify(prev); e != nil {
			return e
		}

		// If tx is to structured to create a kitty, attempt to add to state.
		// Otherwise, attempt to transfer it's ownership in the state.
		if tx.IsKittyGen(bc.c.CreatorPK) {
			if e := bc.state.AddKitty(tx.KittyID, tx.To); e != nil {
				return e
			}
		} else {
			if e := bc.state.MoveKitty(tx.KittyID, tx.From, tx.To); e != nil {
				return e
			}
		}
	}
	return nil
}

func (bc *BlockChain) Close() {
	close(bc.quit)
}

func (bc *BlockChain) service() {
	defer bc.wg.Done()

	for {
		select {
		case bc.quit:
			return

		case tx := <-bc.chain.TxChan():
			if e := bc.c.TxAction(tx); e != nil {
				panic(e)
			}
		}
	}
}

func (bc *BlockChain) GetHeadTx() (Transaction, error) {
	bc.mux.RLock()
	defer bc.mux.RUnlock()

	return bc.chain.Head()
}

func (bc *BlockChain) InjectTx(tx *Transaction) error {
	bc.mux.Lock()
	defer bc.mux.Unlock()

	var prev *Transaction

	if temp, e := bc.chain.Head(); e == nil {
		prev = &temp
	}

	if e := tx.Verify(prev); e != nil {
		return e
	}

	if tx.IsKittyGen(bc.c.CreatorPK) {
		if e := bc.state.AddKitty(tx.KittyID, tx.To); e != nil {
			return e
		}
	} else {
		if e := bc.state.MoveKitty(tx.KittyID, tx.From, tx.To); e != nil {
			return e
		}
	}

	if e := bc.chain.AddTx(*tx); e != nil {
		panic(e)
	}

	return nil
}
