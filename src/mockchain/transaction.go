package mockchain

import (
	"errors"
	"github.com/skycoin/skycoin/src/cipher"
	"github.com/skycoin/skycoin/src/cipher/encoder"
	"log"
	"time"
)

type TxAction func(tx *Transaction) error

type Transaction struct {
	Prev cipher.SHA256
	Seq  uint64 // Each transaction has a sequence.
	TS   int64  // Timestamp.

	KittyID uint64
	From    cipher.Address
	To      cipher.Address
	Sig     cipher.Sig
}

func NewGenTx(prev *Transaction, kittyID uint64, sk cipher.SecKey) *Transaction {
	var (
		address = cipher.AddressFromSecKey(sk)
		ts      = time.Now().UnixNano()
		tx      *Transaction
	)
	if prev == nil {
		tx = &Transaction{
			Prev:    cipher.SHA256{},
			Seq:     0,
			TS:      ts,
			KittyID: kittyID,
			From:    address,
			To:      address,
		}
	} else {
		tx = &Transaction{
			Prev:    prev.Hash(),
			Seq:     prev.Seq + 1,
			TS:      ts,
			KittyID: kittyID,
			From:    address,
			To:      address,
		}
	}
	tx.Sig = tx.Sign(sk)
	return tx
}

func NewTransferTx(prev *Transaction, kittyID uint64, to cipher.Address, sk cipher.SecKey) *Transaction {
	tx := &Transaction{
		Prev:    prev.Hash(),
		Seq:     prev.Seq + 1,
		TS:      time.Now().UnixNano(),
		KittyID: kittyID,
		From:    cipher.AddressFromSecKey(sk),
		To:      to,
	}
	tx.Sig = tx.Sign(sk)
	return tx
}

func (tx Transaction) Serialize() []byte {
	return encoder.Serialize(tx)
}

func (tx Transaction) Hash() cipher.SHA256 {
	return cipher.SumSHA256(tx.Serialize())
}

func (tx Transaction) HashInner() cipher.SHA256 {
	tx.Sig = cipher.Sig{}
	return tx.Hash()
}

func (tx Transaction) Sign(sk cipher.SecKey) cipher.Sig {
	e := cipher.
		AddressFromSecKey(sk).
		Verify(cipher.PubKeyFromSecKey(sk))
	if e != nil {
		log.Panic(e)
	}
	return cipher.SignHash(tx.HashInner(), sk)
}

// Verify checks the hash, seq and signature of the transaction.
//		- Previous tx hash.
//		- Tx sequence.
//		- Tx signature.
// Verify does not check:
//		- Whether from address actually owns the kitty of ID.
//		- Double spending of kitties.
// TODO (evanlinjin): Write tests.
func (tx Transaction) Verify(prev *Transaction) error {
	isGenesis := prev == nil

	// Check hash.
	if isGenesis {
		if tx.Prev != (cipher.SHA256{}) {
			return errors.New("invalid prev hash")
		}
	} else {
		if tx.Prev != prev.Hash() {
			return errors.New("invalid prev hash")
		}
	}

	// Check seq.
	if isGenesis {
		if tx.Seq != 0 {
			return errors.New("invalid seq")
		}
	} else {
		if tx.Seq != prev.Seq+1 {
			return errors.New("invalid seq")
		}
	}

	return cipher.ChkSig(tx.From, tx.HashInner(), tx.Sig)
}

// IsKittyGen returns true if:
//		- Tx is of the correct structure to create a new kitty.
//		- Tx is of the right address to create a new kitty.
// TODO (evanlinjin): Write tests.
func (tx Transaction) IsKittyGen(pk cipher.PubKey) bool {
	// Check from address.
	if e := tx.From.Verify(pk); e != nil {
		return false
	}
	// Check to & from addresses are the same.
	if tx.To != tx.From {
		return false
	}
	// Accept.
	return true
}
