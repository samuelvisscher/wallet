package store

import "github.com/skycoin/skycoin/src/cipher"

type OwnershipState byte

const (
	StateNoOwner  OwnershipState = iota
	StateReserved OwnershipState = iota
	StateOwned    OwnershipState = iota
)

type KittyOwnership struct {
	KId     uint64
	State   OwnershipState
	Address cipher.Address
}
