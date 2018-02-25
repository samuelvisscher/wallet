package http

import (
	"encoding/hex"
	"fmt"
	"github.com/kittycash/wallet/src/iko"
	"github.com/skycoin/skycoin/src/cipher"
	"net/http"
	"path"
	"strconv"
)

func ikoGateway(mux *http.ServeMux, g *iko.BlockChain) error {
	mux.HandleFunc("/api/iko/kitty/", getKitty(g))
	mux.HandleFunc("/api/iko/tx/", getTx(g))
	mux.HandleFunc("/api/iko/tx_seq/", getTxOfSeq(g))
	mux.HandleFunc("/api/iko/address/", getAddress(g))
	return nil
}

type KittyReply struct {
	KittyID      iko.KittyID `json:"kitty_id"`
	Address      string      `json:"address"`
	Transactions []string    `json:"transactions"`
}

func getKitty(g *iko.BlockChain) http.HandlerFunc {
	fmt.Println("GETTING KITTY")
	return Do(func(w http.ResponseWriter, r *http.Request) error {
		kittyID, e := iko.KittyIDFromString(path.Base(r.URL.EscapedPath()))
		if e != nil {
			return send400(w, e)
		}
		kState, ok := g.GetKittyState(kittyID)
		if !ok {
			return send404(w, fmt.Errorf("kitty of id '%s' not found", kittyID))
		}
		return send200(w, KittyReply{
			KittyID:      kittyID,
			Address:      kState.Address.String(),
			Transactions: kState.Transactions.ToStringArray(),
		})
	})
}

type TxMeta struct {
	Hash string `json:"hash"`
	Raw  string `json:"raw"`
}

type Tx struct {
	PrevHash string      `json:"prev_hash"`
	Seq      uint64      `json:"seq"`
	TS       int64       `json:"time"`
	KittyID  iko.KittyID `json:"kitty_id"`
	From     string      `json:"from"`
	To       string      `json:"to"`
	Sig      string      `json:"sig"`
}

type TxReply struct {
	Meta TxMeta `json:"meta"`
	Tx   Tx     `json:"transaction"`
}

func getTx(g *iko.BlockChain) http.HandlerFunc {
	return Do(func(w http.ResponseWriter, r *http.Request) error {
		txHash, e := cipher.SHA256FromHex(path.Base(r.URL.EscapedPath()))
		if e != nil {
			return send400(w, e)
		}
		tx, e := g.GetTxOfHash(iko.TxHash(txHash))
		if e != nil {
			return send404(w, e)
		}
		return send200(w, TxReply{
			Meta: TxMeta{
				Hash: tx.Hash().Hex(),
				Raw:  hex.EncodeToString(tx.Serialize()),
			},
			Tx: Tx{
				PrevHash: tx.Prev.Hex(),
				Seq:      tx.Seq,
				TS:       tx.TS,
				KittyID:  tx.KittyID,
				From:     tx.From.String(),
				To:       tx.To.String(),
				Sig:      tx.Sig.Hex(),
			},
		})
	})
}

func getTxOfSeq(g *iko.BlockChain) http.HandlerFunc {
	return Do(func(w http.ResponseWriter, r *http.Request) error {
		seq, e := strconv.ParseUint(path.Base(r.URL.EscapedPath()), 10, 64)
		if e != nil {
			return send400(w, e)
		}
		tx, e := g.GetTxOfSeq(seq)
		if e != nil {
			return send404(w, e)
		}
		return send200(w, TxReply{
			Meta: TxMeta{
				Hash: tx.Hash().Hex(),
				Raw:  hex.EncodeToString(tx.Serialize()),
			},
			Tx: Tx{
				PrevHash: tx.Prev.Hex(),
				Seq:      tx.Seq,
				TS:       tx.TS,
				KittyID:  tx.KittyID,
				From:     tx.From.String(),
				To:       tx.To.String(),
				Sig:      tx.Sig.Hex(),
			},
		})
	})
}

type AddressReply struct {
	Address      string       `json:"address"`
	Kitties      iko.KittyIDs `json:"kitties"`
	Transactions []string     `json:"transactions"`
}

func getAddress(g *iko.BlockChain) http.HandlerFunc {
	return Do(func(w http.ResponseWriter, r *http.Request) error {
		address, e := cipher.DecodeBase58Address(path.Base(r.URL.EscapedPath()))
		if e != nil {
			return send400(w, e)
		}
		aState := g.GetAddressState(address)
		return send200(w, AddressReply{
			Address:      address.String(),
			Kitties:      aState.Kitties,
			Transactions: aState.Transactions.ToStringArray(),
		})
	})
}