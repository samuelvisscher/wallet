package store

import "github.com/skycoin/skycoin/src/cipher"

type Transaction struct {
	Length uint32         // Length prefix.
	InnerHash cipher.SHA256 // Only hashing input and outputs.

	Sig    cipher.Sig
	In     cipher.Address // Output being spent.
	Out    TransactionOutput
}

type TransactionOutput struct {
	Address cipher.Address
	KIDs    []uint64 // Kitties (represented as ID) to be sent.
}
