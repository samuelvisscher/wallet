package mockchain

import (
	"errors"
	"github.com/skycoin/skycoin/src/cipher"
	"github.com/skycoin/skycoin/src/cipher/encoder"
	"gopkg.in/sirupsen/logrus.v1"
)

type Block struct {
	Head BlockHeader
	Body BlockBody
}

type BlockHeader struct {
	Time int64
	Seq  uint64

	PrevHash cipher.SHA256 // hash of previous block's header
	BodyHash cipher.SHA256 // hash of this block
}

type BlockBody struct {
	Transactions []Transaction
}

type SignedBlock struct {
	Block
	Sig cipher.Sig
}

func NewBlock(prev Block, now int64, txs []Transaction) (*Block, error) {
	if len(txs) == 0 {
		return nil, errors.New("refusing to create block with no transactions")
	}
	var body = BlockBody{Transactions: txs}
	return &Block{
		Head: NewBlockHeader(prev.Head, now, body),
		Body: body,
	}, nil
}

func NewBlockHeader(prev BlockHeader, now int64, body BlockBody) BlockHeader {
	if now <= prev.Time {
		logrus.Panic("time can only move foward")
	}
	return BlockHeader{
		Time:     now,
		Seq:      prev.Seq + 1,
		PrevHash: prev.Hash(),
		BodyHash: body.Hash(),
	}
}

func NewGenesisBlock(address cipher.Address, kittyCount uint64, ts int64) (*Block, error) {
	var txs = make([]Transaction, kittyCount)
	for i := uint64(0); i < kittyCount; i++ {
		txs[i] = Transaction{
			KittyID: i,
			To:      address,
		}
	}
	body := BlockBody{
		Transactions: txs,
	}
	return &Block{
		Head: BlockHeader{
			Time: ts,
			Seq: 0,
			PrevHash: cipher.SHA256{},
			BodyHash: body.Hash(),
		},
		Body: body,
	}, nil
}

func (bh BlockHeader) Hash() cipher.SHA256 {
	return cipher.SumSHA256(bh.Serialize())
}

func (bh BlockHeader) Serialize() []byte {
	return encoder.Serialize(bh)
}

func (bb BlockBody) Hash() cipher.SHA256 {
	return cipher.SumSHA256(bb.Serialize())
}

func (bb BlockBody) Serialize() []byte {
	return encoder.Serialize(bb)
}
