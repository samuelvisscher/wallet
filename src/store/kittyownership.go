package store

import "github.com/skycoin/skycoin/src/cipher"

type OwnershipStatus byte

const (
	OwnershipNone     OwnershipStatus = iota // No owner.
	OwnershipReserved OwnershipStatus = iota // Someone attempting to buy.
	OwnershipSold     OwnershipStatus = iota // Kitty Sold.
)

type KittyOwnership struct {
	Status  OwnershipStatus `json:"status"`
	Reserve cipher.Address  `json:"reserve_address,omitempty"`
	Owner   cipher.Address  `json:"owner_address,omitempty"`
}
