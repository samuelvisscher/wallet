package store

import "github.com/skycoin/skycoin/src/cipher"

type KittyStore struct {
	db DB
}

func NewKittyStore(db DB) *KittyStore {
	return &KittyStore{
		db: db,
	}
}

func SetOwner(kh cipher.SHA256, owner cipher.Address) error {
	return nil
}
