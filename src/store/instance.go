package store

import "github.com/skycoin/skycoin/src/cipher"

type Instance struct {
	metas      MetaStorer
	ownerships OwnershipStorer
}

func NewInstance(metas MetaStorer, ownerships OwnershipStorer) *Instance {
	return &Instance{
		metas:      metas,
		ownerships: ownerships,
	}
}

func SetOwner(kID uint64, owner cipher.Address) error {
	return nil
}
