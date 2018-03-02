package cxo

import (
	"github.com/skycoin/cxo/node"
	"github.com/skycoin/cxo/node/log"
	"github.com/skycoin/cxo/skyobject"
)

func NewCXOConfig() node.Config {
	var sc node.Config
	sc.Config = gnet.NewConfig()
	sc.Log = log.NewConfig()
	sc.Skyobject = skyobject.NewConfig()
	sc.EnableRPC = node.EnableRPC
	sc.RPCAddress = node.RPCAddress
	sc.Listen = node.Listen
	sc.EnableListener = node.EnableListener
	sc.RemoteClose = node.RemoteClose
	sc.PingInterval = node.PingInterval
	sc.InMemoryDB = node.InMemoryDB
	//sc.DataDir = node.DataDir()
	sc.DBPath = ""
	sc.ResponseTimeout = node.ResponseTimeout
	sc.PublicServer = node.PublicServer
	sc.Config.OnDial = node.OnDialFilter
	return sc
}