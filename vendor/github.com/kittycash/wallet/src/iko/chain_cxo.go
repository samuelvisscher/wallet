package iko

import (
	"github.com/skycoin/cxo/node"
	"gopkg.in/sirupsen/logrus.v1"
	"sync"
)

type CXOChainConfig struct {
	Dir           string
	Public        bool
	CXOAddress    string
	CXORPCEnable  bool
	CXORPCAddress string
}

type CXOChain struct {
	mux  sync.Mutex
	c    *CXOChainConfig
	l    *logrus.Logger
	node *node.Node
}

func NewCXOChain() (*CXOChain, error) {
	node.NewConfig()

	return nil, nil
}
