package store

import (
	"github.com/skycoin/skycoin/src/cipher"
	"errors"
	"sync"
)

var (
	ErrInvalidID = errors.New("there is no kitty of such an ID")
	ErrAlreadyReserved = errors.New("already reserved")
	ErrAlreadyOwned    = errors.New("already owned")
	ErrUnknown         = errors.New("unknown error")
)

type OwnershipState struct {
	kID                uint64
	Reserved           bool
	ReservationAddress cipher.Address
	Owned              bool
	OwnerAddress       cipher.Address
}

type OwnershipStorer interface {
	EnsureCount(count int)
	SetOwnershipState(kID uint64, state OwnershipState) error
	GetOwnershipState(kID uint64) (OwnershipState, error)
	GetKittiesOfOwner(addresses ...cipher.Address) []OwnershipState
	GetKittiesOfReserver(addresses ...cipher.Address) []OwnershipState
}

type OwnershipMemoryDB struct {
	sync.RWMutex
	kitties   []OwnershipState // index (kitty ID), element (owner/reservation state of the kitty)
	owners    map[cipher.Address]map[uint64]struct{}
	reservers map[cipher.Address]uint64
}

func NewMutableMemory() *OwnershipMemoryDB {
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
			make([]OwnershipState, count-len(mm.kitties))...)
		for i := range mm.kitties {
			mm.kitties[i].kID = uint64(i)
		}
	}
}

func (mm *OwnershipMemoryDB) SetOwnershipState(kID uint64, state OwnershipState) error {
	mm.Lock()
	defer mm.Unlock()

	if kID >= uint64(len(mm.kitties)) {
		return ErrInvalidID
	}

	var originalState = mm.kitties[kID]

	if originalState.Reserved {
		delete(mm.reservers, originalState.ReservationAddress)

	} else if originalState.Owned {
		list, ok := mm.owners[originalState.OwnerAddress]
		if ok {
			delete(list, kID)
		}
	}

	mm.kitties[kID] = state

	if state.Reserved {
		mm.reservers[state.ReservationAddress] = kID

	} else if state.Owned {
		list, ok := mm.owners[state.OwnerAddress]
		if !ok {
			list = make(map[uint64]struct{})
			mm.owners[state.OwnerAddress] = list
		}
		list[kID] = struct{}{}
	}

	return nil
}

func (mm *OwnershipMemoryDB) GetOwnershipState(kID uint64) (OwnershipState, error) {
	mm.RLock()
	defer mm.RUnlock()

	if kID >= uint64(len(mm.kitties)) {
		return OwnershipState{}, ErrInvalidID
	}
	return mm.kitties[kID], nil
}

func (mm *OwnershipMemoryDB) GetKittiesOfOwner(addresses ...cipher.Address) []OwnershipState {
	mm.Lock()
	defer mm.Unlock()

	var out []OwnershipState

	for _, address := range addresses {
		if list, ok := mm.owners[address]; ok {
			for kID := range list {
				out = append(out, mm.kitties[kID])
			}
		}
	}

	return out
}

func (mm *OwnershipMemoryDB) GetKittiesOfReserver(addresses ...cipher.Address) []OwnershipState {
	mm.Lock()
	defer mm.Unlock()

	var out []OwnershipState

	for _, address := range addresses {
		if kID, ok := mm.reservers[address]; ok {
			out = append(out, mm.kitties[kID])
		}
	}

	return out
}