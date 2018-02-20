package mockchain

import (
	"github.com/skycoin/skycoin/src/cipher"
	"sync"
	"fmt"
)

type StateDB interface {
	GetAddressOfKitty(kittyID uint64) (cipher.Address, error)
	GetKittiesOfAddress(address cipher.Address) []uint64
	AddKitty(kittyID uint64, address cipher.Address) error
	MoveKitty(kittyID uint64, from, to cipher.Address) error
}

type MemoryState struct {
	sync.Mutex
	kitties   map[uint64]cipher.Address
	addresses map[cipher.Address]map[uint64]struct{}
}

func NewMemoryState() *MemoryState {
	return &MemoryState{
		kitties:   make(map[uint64]cipher.Address),
		addresses: make(map[cipher.Address]map[uint64]struct{}),
	}
}

func (s *MemoryState) GetAddressOfKitty(kittyID uint64) (cipher.Address, error) {
	s.Lock()
	defer s.Unlock()

	address, ok := s.kitties[kittyID]
	if !ok {
		return address, fmt.Errorf("kitty of id '%d' is not recorded in state",
			kittyID)
	}
	return address, nil
}

func (s *MemoryState) GetKittiesOfAddress(address cipher.Address) []uint64 {
	s.Lock()
	defer s.Unlock()

	kMap, ok := s.addresses[address]
	if !ok {
		return make([]uint64, 0)
	}

	kittyIDs, i := make([]uint64, len(kMap)), 0
	for id := range kMap {
		kittyIDs[i], i = id, i+1
	}
	return kittyIDs
}

func (s *MemoryState) AddKitty(kittyID uint64, address cipher.Address) error {
	s.Lock()
	defer s.Unlock()

	if _, ok := s.kitties[kittyID]; ok {
		return fmt.Errorf("kitty of id '%d' already exists",
			kittyID)
	}

	kMap, ok := s.addresses[address]
	if !ok {
		kMap = make(map[uint64]struct{}, 1)
		s.addresses[address] = kMap
	}
	kMap[kittyID] = struct{}{}

	return nil
}

func (s *MemoryState) MoveKitty(kittyID uint64, from, to cipher.Address) error {

	if address, ok := s.kitties[kittyID]; !ok {
		return fmt.Errorf("kitty of id '%d' does not exist",
			kittyID)
	} else if address == from {
		return fmt.Errorf("kitty of id '%d' does not belong to address '%s'",
			kittyID, from)
	}

	s.kitties[kittyID] = to
	delete(s.addresses[from], kittyID)

	kMap, ok := s.addresses[to]
	if !ok {
		kMap = make(map[uint64]struct{}, 1)
		s.addresses[to] = kMap
	}
	kMap[kittyID] = struct{}{}

	return nil
}