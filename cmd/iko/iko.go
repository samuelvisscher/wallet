package main

import (
	"gopkg.in/urfave/cli.v1"
	"os"
	"log"
	"github.com/kittycash/iko/src/kchain"
	"github.com/skycoin/skycoin/src/cipher"
)

const (
	MemoryMode      = "memory"
	TestMode        = "test"
	MasterPublicKey = "master-public-key"
)

func Flag(flag, short string) string {
	return flag + ", " + short
}

var app = cli.NewApp()

func init() {
	app.Name = "iko"
	app.Description = "kittycash initial coin offering service"
	app.Flags = cli.FlagsByName{
		cli.BoolFlag{
			Name: Flag(MemoryMode, "m"),
			Usage: "whether to run in memory-only mode",
		},
		cli.BoolFlag{
			Name: Flag(TestMode, "t"),
			Usage: "whether to use test data for run",
		},
		cli.StringFlag{
			Name: Flag(MasterPublicKey, "pk"),
			Usage: "public key to trust as master decision maker",
		},
	}
	app.Action = cli.ActionFunc(action)
}

func action(ctx *cli.Context) error {

	bc, e := kchain.NewBlockChain(
		&kchain.BlockChainConfig{
			CreatorPK: cipher.MustPubKeyFromHex(ctx.String(MasterPublicKey)),
			TxAction: func(tx *kchain.Transaction) error {
				return nil
			},
		},
		kchain.NewMemoryChain(10),
		kchain.NewMemoryState(),
	)
	if e != nil {
		return e
	}
	for i := 0; i < 10; i++ {
		e := bc.InjectTx(&kchain.Transaction{

		})
		if e != nil {
			return e
		}
	}
	return nil
}

func main() {
	if e := app.Run(os.Args); e != nil {
		log.Println(e)
	}
}
