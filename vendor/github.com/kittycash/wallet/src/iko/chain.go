package iko

import (
	"errors"
	"fmt"
	"sync"
)

type ChainDB interface {
	Head() (Transaction, error)
	HeadSeq() uint64
	Len() uint64
	AddTx(tx Transaction) error
	GetTxOfHash(hash TxHash) (Transaction, error)
	GetTxOfSeq(seq uint64) (Transaction, error)
	TxChan() <-chan *Transaction
}

type MemoryChain struct {
	sync.RWMutex
	txs    []Transaction
	byHash map[TxHash]*Transaction
	txChan chan *Transaction
}

func NewMemoryChain(bufferSize int) *MemoryChain {
	return &MemoryChain{
		byHash: make(map[TxHash]*Transaction),
		txChan: make(chan *Transaction, bufferSize),
	}
}

func (c *MemoryChain) Head() (Transaction, error) {
	c.RLock()
	defer c.RUnlock()

	if len(c.txs) == 0 {
		return Transaction{}, errors.New("no transactions")
	}
	return c.txs[len(c.txs)-1], nil
}

func (c *MemoryChain) HeadSeq() uint64 {
	c.RLock()
	defer c.RUnlock()

	return uint64(len(c.txs)) - 1
}

func (c *MemoryChain) Len() uint64 {
	c.RLock()
	defer c.RUnlock()

	return uint64(len(c.txs))
}

func (c *MemoryChain) AddTx(tx Transaction) error {
	c.Lock()
	defer c.Unlock()

	c.txs = append(c.txs, tx)
	c.byHash[tx.Hash()] = &c.txs[len(c.txs)-1]
	go func() {
		c.txChan <- &c.txs[len(c.txs)-1]
	}()
	return nil
}

func (c *MemoryChain) GetTxOfHash(hash TxHash) (Transaction, error) {
	c.Lock()
	defer c.Unlock()

	tx, ok := c.byHash[hash]
	if !ok {
		return Transaction{}, fmt.Errorf("tx of hash '%s' does not exist", hash.Hex())
	}
	return *tx, nil
}

func (c *MemoryChain) GetTxOfSeq(seq uint64) (Transaction, error) {
	c.RLock()
	defer c.RUnlock()

	if seq >= uint64(len(c.txs)) {
		return Transaction{}, fmt.Errorf("block of sequence '%d' does not exist", seq)
	}
	return c.txs[seq], nil
}

func (c *MemoryChain) TxChan() <-chan *Transaction {
	return nil
}
