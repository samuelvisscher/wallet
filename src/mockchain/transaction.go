package mockchain

import (
	"github.com/skycoin/skycoin/src/cipher"
)

type Transaction struct {
	KittyID uint64
	From    cipher.Address
	To      cipher.Address
	Sig     cipher.Sig
}

