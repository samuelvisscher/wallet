package iko

import (
	"fmt"
	"github.com/skycoin/skycoin/src/cipher"
	"sync"
)

// StateDB records the state of the blockchain.
type StateDB interface {

	// GetAddressOfKitty obtains the address that the kitty is owned under.
	// It should return an error if kitty of specified ID does not exist.
	GetAddressOfKitty(kittyID KittyID) (cipher.Address, error)

	// GetKittiesOfAddress obtains the kitties that are owned under a specified address.
	// The array of kitty IDs should be in ascending sequential order, from smallest index to highest.
	GetKittiesOfAddress(address cipher.Address) KittyIDs

	// AddKitty adds a kitty to the state under the specified address.
	// This should fail if:
	// 		- kitty of specified ID already exists in state.
	AddKitty(kittyID KittyID, address cipher.Address) error

	// MoveKitty moves a kitty from one address to another.
	// This should fail if:
	//		- kitty of specified ID already belongs to the address ('from' and 'to' addresses are the same).
	//		- kitty of specified ID does not exist.
	//		- kitty of specified ID does not originally belong to the 'from' address.
	MoveKitty(kittyID KittyID, from, to cipher.Address) error
}

type MemoryState struct {
	sync.Mutex
	kitties   map[KittyID]cipher.Address
	addresses map[cipher.Address]*KittyIDs
}

func NewMemoryState() *MemoryState {
	return &MemoryState{
		kitties:   make(map[KittyID]cipher.Address),
		addresses: make(map[cipher.Address]*KittyIDs),
	}
}

func (s *MemoryState) GetAddressOfKitty(kittyID KittyID) (cipher.Address, error) {
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
func (s MemoryState) GetKittiesOfAddress(address cipher.Address) KittyIDs {
	s.Lock()
	defer s.Unlock()

	kitties, ok := s.addresses[address]
	if !ok {
		return make([]KittyID, 0)
	}
	return *kitties
}

func (s *MemoryState) AddKitty(kittyID KittyID, address cipher.Address) error {
	s.Lock()
	defer s.Unlock()

	if _, ok := s.kitties[kittyID]; ok {
		return fmt.Errorf("kitty of id '%d' already exists",
			kittyID)
	}

	s.kitties[kittyID] = address

	if kitties, ok := s.addresses[address]; !ok {
		s.addresses[address] = &KittyIDs{kittyID}
	} else {
		kitties.Add(kittyID)
	}
	return nil
}

func (s *MemoryState) MoveKitty(kittyID KittyID, from, to cipher.Address) error {
	s.Lock()
	defer s.Unlock()

	if from == to {
		return fmt.Errorf("kitty of id '%d' already belongs to address '%s'",
			kittyID, from)

	} else if address, ok := s.kitties[kittyID]; !ok {
		return fmt.Errorf("kitty of id '%d' does not exist",
			kittyID)

	} else if address != from {
		return fmt.Errorf("kitty of id '%d' does not belong to address '%s'",
			kittyID, from)
	}

	s.kitties[kittyID] = to
	s.addresses[from].Remove(kittyID)

	if kitties, ok := s.addresses[to]; !ok {
		s.addresses[to] = &KittyIDs{kittyID}
	} else {
		kitties.Add(kittyID)
	}
	return nil
}
