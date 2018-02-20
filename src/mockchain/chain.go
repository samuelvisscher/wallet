package mockchain

import (
	"github.com/skycoin/skycoin/src/cipher"
	"sync"
	"fmt"
)

type ChainDB interface {
	Head() (*SignedBlock, error)
	HeadSeq() uint64
	Len() uint64
	AddBlock(block *SignedBlock) error
	GetBlockOfHash(hash cipher.SHA256) (*SignedBlock, error)
	GetBlockOfSeq(seq uint64) (*SignedBlock, error)
}

type MemoryChain struct {
	sync.RWMutex
	blocks []SignedBlock
	byHash map[cipher.SHA256]*SignedBlock
}

func (c *MemoryChain) Head() (*SignedBlock, error) {
	c.RLock()
	defer c.RUnlock()

	block := c.blocks[len(c.blocks)-1]
	return &block, nil
}

func (c *MemoryChain) HeadSeq() uint64 {
	c.RLock()
	defer c.RUnlock()

	return uint64(len(c.blocks))-1
}

func (c *MemoryChain) Len() uint64 {
	c.RLock()
	defer c.RUnlock()

	return uint64(len(c.blocks))
}

func (c *MemoryChain) AddBlock(block *SignedBlock) error {
	c.Lock()
	defer c.Unlock()

	// Some checks?

	c.blocks = append(c.blocks, *block)
	c.byHash[block.GetHeaderHash()] = &c.blocks[len(c.blocks)-1]
	return nil
}

func (c *MemoryChain) GetBlockOfHash(hash cipher.SHA256) (*SignedBlock, error) {
	c.Lock()
	defer c.Unlock()

	block, ok := c.byHash[hash]
	if !ok {
		return nil, fmt.Errorf("block of hash '%s' does not exist", hash.Hex())
	}

	return &(*block), nil
}

func (c *MemoryChain) GetBlockOfSeq(seq uint64) (*SignedBlock, error) {
	c.RLock()
	defer c.RUnlock()

	if seq >= uint64(len(c.blocks)) {
		return nil, fmt.Errorf("block of sequence '%d' does not exist", seq)
	}

	return &c.blocks[seq], nil
}