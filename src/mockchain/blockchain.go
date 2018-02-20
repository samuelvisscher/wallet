package mockchain

import (
	"github.com/skycoin/skycoin/src/cipher"
)

type BlockChainConfig struct {
	Master bool
	PK     cipher.PubKey
	SK     cipher.SecKey
}

type BlockChain struct {
	c     *BlockChainConfig
	chain ChainDB
	state StateDB
}

func NewBlockChain(config *BlockChainConfig, genesis *SignedBlock, chainDB ChainDB, stateDB StateDB) (*BlockChain, error) {
	bc := &BlockChain{
		c:     config,
		chain: chainDB,
		state: stateDB,
	}

	if bc.chain.Len() == 0 {
		if e := bc.chain.AddBlock(genesis); e != nil {
			return nil, e
		}
	}
	return bc, nil
}

func (bc *BlockChain) InitState() error {
	return nil
}