package store

import "github.com/skycoin/skycoin/src/cipher"

type KittyOwnerShip struct {
	kID                uint64
	Reserved           bool
	ReservationAddress cipher.Address
	Owned              bool
	OwnerAddress       cipher.Address
}