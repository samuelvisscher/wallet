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

	t.Run("TransactionCreated_NoTransactionAvailable", func(t *testing.T) {
		// Add a kitty
		txHash := TxHash(cipher.SumSHA256([]byte{3, 4, 5, 6}))
		kID := KittyID(3)

		err := stateDB.AddKitty(txHash, kID, anAddress)
		if err == nil {
			prev := NewGenTx(nil, kID, sk)
			// Since there is no previous transactions then an error should be thrown
			require.EqualError(t, prev.Verify(prev), "invalid prev hash")
		} else {
			fmt.Println("StateDB error, failed to create Kitty")
		}
	})

	fmt.Println(sk)
}
func TestTransaction_Verify(t *testing.T) {
	stateDB := NewMemoryState()
	runTransactionVerifyTest(t, stateDB)
}
