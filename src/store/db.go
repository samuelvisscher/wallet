package store

import "github.com/skycoin/skycoin/src/cipher"

type DB interface {
	/*
		<<< ADMIN >>>
	*/

	// AddKitty adds a kitty to DB.
	AddKitty(image []byte, meta *KittyMeta) (cipher.SHA256, error)

	// RemoveKitty removes a kitty from DB.
	RemoveKitty(kID uint64)

	/*
		<<< OWNERSHIP >>>
	*/

	// SetKittyOwner changes the ownership of a kitty of id.
	SetKittyOwner(kID uint64, address cipher.Address) error

	// ReserveKitty reserves a kitty for a given address.
	ReserveKitty(kID uint64, address cipher.Address) error

	// RemoveReservation removes a kitty's reservation.
	RemoveReservation(kID uint64)

	// IsReserved checks whether a kitty of id is reserved.
	IsReserved(kID uint64) bool

	/*
		<<< VIEW >>>
	*/

	// GetKittyOfID obtains a kitty of given kitty id.
	GetKittyOfID(kID uint64) (*KittyMeta, error)

	// ListAllKitties lists all kitties, claimed or unclaimed.
	ListAllKitties() []*KittyMeta

	// ListClaimedKitties lists only the kitties that are claimed.
	// If len(owners) > 0, only kitties owned by given addresses will be listed.
	ListClaimedKitties(addresses ...cipher.Address) []*KittyMeta

	// ListUnclaimedKitties lists only the kitties that are unclaimed.
	ListUnclaimedKitties() []*KittyMeta
}
