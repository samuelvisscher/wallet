package iko

import (
	/*	"errors"*/
	"fmt"
	"github.com/skycoin/skycoin/src/cipher"
	"github.com/stretchr/testify/require"
	"testing"
)

func runTransactionVerifyTest(t *testing.T, stateDB StateDB) {
	sk := cipher.SecKey([32]byte{
		3, 4, 5, 6,
		3, 4, 5, 6,
		3, 4, 5, 6,
		3, 4, 5, 6,
		3, 4, 5, 6,
		3, 4, 5, 6,
		3, 4, 5, 6,
		3, 4, 5, 6,
	})
	anAddress := cipher.AddressFromSecKey(sk)

	t.Run("TransactionCreated_InvalidPrevTransaction", func(t *testing.T) {
		// Add a kitty
		txHash := TxHash(cipher.SumSHA256([]byte{3, 4, 5, 6}))
		kID := KittyID(3)

		err := stateDB.AddKitty(txHash, kID, anAddress)

		// If there's an error creating kitty, then deviate testing transaction -- no kitty means no transaction
		if err == nil {
			prev := NewGenTx(nil, kID, sk)
			// Since there is no previous transactions then an error should be thrown that previous hash is invalid
			require.Errorf(t, prev.Verify(prev), "There is no previous tansaction to verify")
		} else {
			fmt.Println("StateDB error, failed to create Kitty")
		}
	})

	t.Run("TransactionCreated_InvalidPrevHash", func(t *testing.T) {
		kID := KittyID(3)

		// New secret key for toAddress
		sk2 := cipher.SecKey([32]byte{
			3, 4, 5, 6,
			3, 4, 5, 6,
			3, 4, 5, 6,
			3, 4, 5, 6,
			5, 4, 3, 6,
			3, 4, 5, 6,
			4, 4, 5, 6,
			3, 4, 5, 6,
		})
		toAddress := cipher.AddressFromSecKey(sk2)
		prev := NewGenTx(nil, kID, sk)
		nextTrans := NewTransferTx(prev, kID, toAddress, sk)

		fmt.Println(nextTrans.Seq)
		nextTrans.Prev = TxHash(cipher.SumSHA256([]byte{3, 4, 5, 6}))
		require.Errorf(t, nextTrans.Verify(prev), "Previous hash was changed!!")
	})
}

func TestTransaction_Verify(t *testing.T) {
	stateDB := NewMemoryState()
	runTransactionVerifyTest(t, stateDB)
}
