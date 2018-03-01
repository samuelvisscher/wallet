package iko

import (
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

		// Change transaction previous hash to test if verify return error
		nextTrans.Prev = TxHash(cipher.SumSHA256([]byte{3, 4, 5, 6}))
		require.Errorf(t, nextTrans.Verify(prev), "Previous hash was changed!!")

		// Revert transaction previous hash to its original state and change seqence number to test if verfiy returns error
		nextTrans.Prev = prev.Hash()
		nextTrans.Seq = prev.Seq + 5
		require.Errorf(t, nextTrans.Verify(prev), "Previous Seq was changed!!")

		nextTrans.Seq = prev.Seq
		require.Errorf(t, nextTrans.Verify(prev), "Previous Seq was changed")

		// Revert transaction sequence to its original state and change TS to test if Verify will return an error
		nextTrans.Seq = prev.Seq + 1
		TS := nextTrans.TS
		nextTrans.TS = prev.TS - 10
		require.Errorf(t, nextTrans.Verify(prev), "TS was changed and should be invalid")
		fmt.Println(nextTrans.Verify(prev))

		// Revert transaction TS to its original state and test to ensure function returns nil when transaction is valid
		nextTrans.TS = TS
		require.Equal(t, nextTrans.Verify(prev), nil, "Verify should return nil for valid transactions")
	})

	t.Run("", func(t *testing.T) {

	})
}

func TestTransaction_Verify(t *testing.T) {
	stateDB := NewMemoryState()
	runTransactionVerifyTest(t, stateDB)
}
