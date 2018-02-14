package store

import (
	"errors"
	"github.com/skycoin/skycoin/src/cipher"
	"sync"
)

var (
	ErrInvalidID       = errors.New("there is no kitty of such an ID")
	ErrAlreadyReserved = errors.New("already reserved")
	ErrAlreadyOwned    = errors.New("already owned")
	ErrUnknown         = errors.New("unknown error")
)

type OwnershipStorer interface {
	EnsureCount(count int)
	SetOwnershipState(kID uint64, state KittyOwnership) error
	GetOwnershipState(kID uint64) (KittyOwnership, error)
	GetKittiesOfOwner(addresses ...cipher.Address) []KittyOwnership
	GetKittiesOfReserver(addresses ...cipher.Address) []KittyOwnership
}

type OwnershipMemoryDB struct {
	sync.RWMutex
	kitties   []KittyOwnership // index (kitty ID), element (owner/reservation state of the kitty)
	owners    map[cipher.Address]map[uint64]struct{}
	reservers map[cipher.Address]uint64
}

func NewOwnershipMemoryDB() *OwnershipMemoryDB {
	return &OwnershipMemoryDB{
		owners:    make(map[cipher.Address]map[uint64]struct{}),
		reservers: make(map[cipher.Address]uint64),
	}
}

func (mm *OwnershipMemoryDB) EnsureCount(count int) {
	mm.Lock()
	defer mm.Unlock()

	if len(mm.kitties) < count {
		mm.kitties = append(mm.kitties,
			make([]KittyOwnership, count-len(mm.kitties))...)
		for i := range mm.kitties {
			mm.kitties[i].KId = uint64(i)
		}
	}
}

func (mm *OwnershipMemoryDB) SetOwnershipState(kID uint64, state KittyOwnership) error {
	mm.Lock()
	defer mm.Unlock()

	if kID >= uint64(len(mm.kitties)) {
		return ErrInvalidID
	}

	originalState := mm.kitties[kID]
	switch originalState.State {
	case StateReserved:
		delete(mm.reservers, originalState.Address)
	case StateOwned:
		list, ok := mm.owners[originalState.Address]
		if ok {
			delete(list, kID)
		}
	}
	mm.kitties[kID] = state
	switch state.State {
	case StateReserved:
		mm.reservers[state.Address] = kID
	case StateOwned:
		list, ok := mm.owners[state.Address]
		if !ok {
			list = make(map[uint64]struct{})
			mm.owners[state.Address] = list
		}
		list[kID] = struct{}{}
	}
	return nil
}

func (mm *OwnershipMemoryDB) GetOwnershipState(kID uint64) (KittyOwnership, error) {
	mm.RLock()
	defer mm.RUnlock()

	if kID >= uint64(len(mm.kitties)) {
		return KittyOwnership{}, ErrInvalidID
	}
	return mm.kitties[kID], nil
}

func (mm *OwnershipMemoryDB) GetKittiesOfOwner(addresses ...cipher.Address) []KittyOwnership {
	mm.Lock()
	defer mm.Unlock()

	var out []KittyOwnership
	for _, address := range addresses {
		if list, ok := mm.owners[address]; ok {
			for kID := range list {
				out = append(out, mm.kitties[kID])
			}
		}
	}
	return out
}

func (mm *OwnershipMemoryDB) GetKittiesOfReserver(addresses ...cipher.Address) []KittyOwnership {
	mm.Lock()
	defer mm.Unlock()

	var out []KittyOwnership
	for _, address := range addresses {
		if kID, ok := mm.reservers[address]; ok {
			out = append(out, mm.kitties[kID])
		}
	}
	return out
}
