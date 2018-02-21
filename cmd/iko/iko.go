package main

import (
	"gopkg.in/urfave/cli.v1"
	"os"
	"log"
	"github.com/kittycash/iko/src/kchain"
	"github.com/skycoin/skycoin/src/cipher"
	"os/signal"
	"github.com/kittycash/iko/src/http"
)

const (
	MasterPublicKey = "master-public-key"

	MemoryMode      = "memory"

	TestMode        = "test"
	TestSecretKey   = "test-secret-key"
	TestInjectionCount = "test-injection-count"
)

func Flag(flag, short string) string {
	return flag + ", " + short
}

var app = cli.NewApp()

func init() {
	app.Name = "iko"
	app.Description = "kittycash initial coin offering service"
	app.Flags = cli.FlagsByName{

		cli.StringFlag{
			Name: Flag(MasterPublicKey, "pk"),
			Usage: "public key to trust as master decision maker",
		},

		cli.BoolFlag{
			Name: Flag(MemoryMode, "m"),
			Usage: "whether to run in memory-only mode",
		},

		cli.BoolFlag{
			Name: Flag(TestMode, "t"),
			Usage: "whether to use test data for run",
		},
		cli.StringFlag{
			Name: Flag(TestSecretKey, "sk"),
			Usage: "only valid in test mode, used for injecting transactions",
		},
		cli.IntFlag{
			Name: Flag(TestInjectionCount, "tc"),
			Usage: "only valid in test mode, injects a number of initial transactions for testing",
		},
	}
	app.Action = cli.ActionFunc(action)
}

func action(ctx *cli.Context) error {
	quit := CatchInterrupt()

	var (
		masterPK   = cipher.MustPubKeyFromHex(ctx.String(MasterPublicKey))
		memoryMode = ctx.Bool(MemoryMode)
		testMode  = ctx.Bool(TestMode)
		testSK    = cipher.MustSecKeyFromHex(ctx.String(TestSecretKey))
		testCount = ctx.Int(TestInjectionCount)
	)

	var (
		chainDB kchain.ChainDB
		stateDB kchain.StateDB
	)

	// Prepare ChainDB.
	switch {
	case memoryMode:
		chainDB = kchain.NewMemoryChain(10)
	}

	// Prepare StateDB.
	stateDB = kchain.NewMemoryState()

	// Prepare blockchain config.
	bcConfig := &kchain.BlockChainConfig{
		CreatorPK: masterPK,
		TxAction: func(tx *kchain.Transaction) error {
			return nil
		},
	}

	// Prepare blockchain.
	bc, e := kchain.NewBlockChain(bcConfig, chainDB, stateDB)
	if e != nil {
		return e
	}
	defer bc.Close()

	// Prepare test data.
	if testMode {
		var tx *kchain.Transaction
		for i := 0; i < testCount; i++ {
			tx = kchain.NewGenTx(tx, uint64(i), testSK)
			if e := bc.InjectTx(tx); e != nil {
				return e
			}
		}
	}

	// Prepare http server.
	httpServer, e := http.NewServer(
		&http.ServerConfig{
			Address: "127.0.0.1:1234",
			EnableTLS: false,
		},
		&http.Gateway{
			BlockChain: bc,
		},
	)
	if e != nil {
		return e
	}
	defer httpServer.Close()

	<-quit
	return nil
}

func main() {
	if e := app.Run(os.Args); e != nil {
		log.Println(e)
	}
}

// CatchInterrupt catches Ctrl+C behaviour.
func CatchInterrupt() chan int {
	quit := make(chan int)
	go func(q chan<- int) {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, os.Interrupt)
		<-sigChan
		signal.Stop(sigChan)
		q <- 1
	}(quit)
	return quit
}