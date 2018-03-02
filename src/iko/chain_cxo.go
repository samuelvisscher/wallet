package iko

import (
	"github.com/skycoin/cxo/node"
	"github.com/skycoin/skycoin/src/cipher"
	"gopkg.in/sirupsen/logrus.v1"
	"sync"
)

type CXOChainConfig struct {
	Dir                string
	Public             bool
	MasterPK           cipher.PubKey
	NodePK             cipher.PubKey
	NodeSK             cipher.SecKey
	MessengerAddresses []string
	CXOAddress         string
	CXORPCEnable       bool
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



	return nil
}

