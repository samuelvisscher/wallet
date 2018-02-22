package kchain

import (
	"fmt"
	"github.com/skycoin/skycoin/src/cipher"
	"sync"
)

// StateDB records the state of the blockchain.
type StateDB interface {

	// GetAddressOfKitty obtains the address that the kitty is owned under.
	// It should return an error if kitty of specified ID does not exist.
	GetAddressOfKitty(kittyID uint64) (cipher.Address, error)

	// GetKittiesOfAddress obtains the kitties that are owned under a specified address.
	// The array of kitty IDs should be in ascending sequential order, from smallest index to highest.
	GetKittiesOfAddress(address cipher.Address) []uint64

	// AddKitty adds a kitty to the state under the specified address.
	// This should fail if:
	// 		- kitty of specified ID already exists in state.
	AddKitty(kittyID uint64, address cipher.Address) error

	// MoveKitty moves a kitty from one address to another.
	// This should fail if:
	//		- kitty of specified ID already belongs to the address ('from' and 'to' addresses are the same).
	//		- kitty of specified ID does not exist.
	//		- kitty of specified ID does not originally belong to the 'from' address.
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

// GetKittiesOfAddress GetKittiesOfAddress
// TODO (evanlinjin): return output in ascending sequential order.
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

	s.kitties[kittyID] = address

	kMap, ok := s.addresses[address]
	if !ok {
		kMap = make(map[uint64]struct{}, 1)
		s.addresses[address] = kMap
	}
	kMap[kittyID] = struct{}{}

	return nil
}

func (s *MemoryState) MoveKitty(kittyID uint64, from, to cipher.Address) error {

	if from == to {
		return fmt.Errorf("kitty of id '%d' already belongs to address '%s'",
			kittyID, from)

	} else if address, ok := s.kitties[kittyID]; !ok {
		return fmt.Errorf("kitty of id '%d' does not exist",
			kittyID)

	} else if address == from {
		return fmt.Errorf("kitty of id '%d' does not belong to address '%s'",
			kittyID, from)

	}

	s.kitties[kittyID] = to
	fromKittiesMap := s.addresses[from]
	delete(fromKittiesMap, kittyID)
	if len(fromKittiesMap) == 0 {
		delete(s.addresses, from)
	}

	kMap, ok := s.addresses[to]
	if !ok {
		kMap = make(map[uint64]struct{}, 1)
		s.addresses[to] = kMap
	}
	kMap[kittyID] = struct{}{}

	return nil
}
