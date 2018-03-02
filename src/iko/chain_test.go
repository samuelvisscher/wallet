package iko

import (
	"github.com/stretchr/testify/require"
	"github.com/skycoin/skycoin/src/cipher"
	"fmt"
	"testing"
)

func runChainDBTest(t *testing.T, chainDB ChainDB) {
	t.Run("Head_NoTransactions", func(t *testing.T) {
		_, err := chainDB.Head()

		require.NotNil(t, err, "Should give us an error because there are no transactions yet")
	})

	nonexistentHash := TxHash(cipher.SumSHA256([]byte{3, 4, 5, 6}))

	t.Run("GetTxOfHash_NonexistentHash_01", func(t *testing.T) {
		_, err := chainDB.GetTxOfHash(nonexistentHash)

		require.NotNil(t, err, "Should give us an error because there are no transactions yet")
	})

	t.Run("GetTxOfSeq_NonexistentSeq", func(t *testing.T) {
		_, err := chainDB.GetTxOfSeq(0)

		require.NotNil(t, err, "Should give us an error because there are no transactions yet")
	})

	t.Run("withTransactions", func(t *testing.T) {
		secKey := cipher.SecKey([32]byte{
			3, 4, 5, 6,
			3, 4, 5, 6,
			3, 4, 5, 6,
			3, 4, 5, 6,
			3, 4, 5, 6,
			3, 4, 5, 6,
			3, 4, 5, 6,
			3, 4, 5, 6,
		})

		secondOwnerAddress := cipher.AddressFromSecKey(
			cipher.SecKey([32]byte{
				7, 8, 9, 10,
				7, 8, 9, 10,
				7, 8, 9, 10,
				7, 8, 9, 10,

				7, 8, 9, 10,
				7, 8, 9, 10,
				7, 8, 9, 10,
				7, 8, 9, 10,
			}))

		firstTransaction := NewGenTx(nil, KittyID(5), secKey)

		err := chainDB.AddTx(*firstTransaction)

		require.Nil(t, err, "We should be able to successfully add our first transaction")

		t.Run("Head_Success_01", func(t *testing.T) {
			transaction, err := chainDB.Head()

			require.Nil(t, err, "Should not give us an error")
			require.Equal(t, transaction, *firstTransaction, "Should correctly return the first transaction")
		})

		secondTransaction := NewTransferTx(firstTransaction, KittyID(5), secondOwnerAddress, secKey)

		err = chainDB.AddTx(*secondTransaction)

		require.Nil(t, err, "We should be able to successfully add our second transaction")

		transactions := []Transaction{*firstTransaction, *secondTransaction}

		t.Run("Head_Success_02", func(t *testing.T) {
			transaction, err := chainDB.Head()

			require.Nil(t, err, "Should not give us an error")
			require.Equal(t, transaction, *secondTransaction, "Should correctly return the second transaction")
		})

		t.Run("Len", func(t *testing.T) {
			require.Equal(t, chainDB.Len(), uint64(2), "We should have two transactions by now")
		})

		t.Run("GetTxOfHash_NonexistentHash_02", func(t *testing.T) {
			_, err := chainDB.GetTxOfHash(nonexistentHash)

			require.NotNil(t, err, "Should still give us an error because there are no transactions by that hash")
		})

		for idx, transaction := range transactions {
			// our test label is GetTxOfHash_Success_XX, where 01 is firstTransaction, 02 is secondTransaction, etc
			testLabel := fmt.Sprintf("GetTxOfHash_Success_%2.2d", idx + 1)

			t.Run(testLabel, func(t *testing.T) {
				reqTransaction, err := chainDB.GetTxOfHash(transaction.Hash())

				require.Nil(t, err, "Shouldn't return an error for a valid hash")
				require.Equal(t, transaction, reqTransaction, "Should correctly return the right transaction")
			})
		}

		for idx, transaction := range transactions {
			// same as above
			testLabel := fmt.Sprintf("GetTxOfSeq_Success_%2.2d", idx + 1)

			t.Run(testLabel, func(t *testing.T) {
				reqTransaction, err := chainDB.GetTxOfSeq(transaction.Seq)

				require.Nil(t, err, "Shouldn't return an error for a valid sequence index")
				require.Equal(t, transaction, reqTransaction, "Should correctly return the right transaction")
			})
		}

		t.Run("HeadSeq", func(t *testing.T) {
			require.Equal(t, chainDB.Len() - 1, chainDB.HeadSeq(), "HeadSeq() should be Len() - 1")
		})

		t.Run("TxChan", func(t *testing.T) {
			txChan := chainDB.TxChan()

			require.NotNil(t, txChan, "Should return a valid receiving channel for transactions")

			for _, transaction := range transactions {
				channelTransaction := <-txChan
				require.Equal(t, *channelTransaction, transaction, "Should return our transactions through the TxChan() channel in order they were added")
			}
		})
	})
}

func TestChainDB_MemoryChain(t *testing.T) {
	chainDB := NewMemoryChain(0)

	require.NotNil(t, chainDB, "We should be able to create an empty MemoryChain")

	runChainDBTest(t, chainDB)
}
