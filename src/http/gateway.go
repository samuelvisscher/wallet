package http

import (
	"encoding/hex"
	"encoding/json"
	"github.com/kittycash/iko/src/kchain"
	"github.com/skycoin/skycoin/src/cipher"
	"net/http"
	"path"
	"strconv"
)

type Gateway struct {
	BlockChain *kchain.BlockChain
}

func (g *Gateway) host(mux *http.ServeMux) error {

	type KittyReply struct {
		KittyID      uint64   `json:"kitty_id"`
		Address      string   `json:"address"`
		Transactions []string `json:"transactions"`
	}

	mux.HandleFunc("/api/kitty/",
		func(w http.ResponseWriter, r *http.Request) {
			kittyID, e := strconv.ParseUint(path.Base(r.URL.EscapedPath()), 10, 64)
			if e != nil {
				sendErr(w, e)
				return
			}
			address, e := g.BlockChain.GetKittyAddress(kittyID)
			if e != nil {
				sendErr(w, e)
				return
			}
			sendOK(w, KittyReply{
				KittyID:      kittyID,
				Address:      address.String(),
				Transactions: []string{"TO_BE_IMPLEMENTED"},
			})
		})

	type TxMeta struct {
		Hash string `json:"hash"`
		Raw  string `json:"raw"`
	}

	type Tx struct {
		PrevHash string `json:"prev_hash"`
		Seq      uint64 `json:"seq"`
		TS       int64  `json:"time"`
		KittyID  uint64 `json:"kitty_id"`
		From     string `json:"from"`
		To       string `json:"to"`
		Sig      string `json:"sig"`
	}

	type TxReply struct {
		Meta TxMeta `json:"meta"`
		Tx   Tx     `json:"transaction"`
	}

	mux.HandleFunc("/api/tx/",
		func(w http.ResponseWriter, r *http.Request) {
			txHash, e := cipher.SHA256FromHex(path.Base(r.URL.EscapedPath()))
			if e != nil {
				sendErr(w, e)
				return
			}
			tx, e := g.BlockChain.GetTxOfHash(txHash)
			if e != nil {
				sendErr(w, e)
				return
			}
			sendOK(w, TxReply{
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

	mux.HandleFunc("/api/tx_seq/",
		func(w http.ResponseWriter, r *http.Request) {
			seq, e := strconv.ParseUint(path.Base(r.URL.EscapedPath()), 10, 64)
			if e != nil {
				sendErr(w, e)
				return
			}
			tx, e := g.BlockChain.GetTxOfSeq(seq)
			if e != nil {
				sendErr(w, e)
				return
			}
			sendOK(w, TxReply{
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

	type AddressReply struct {
		Address      string   `json:"address"`
		Kitties      []uint64 `json:"kitties"`
		Transactions []string `json:"transactions"`
	}

	mux.HandleFunc("/api/address/",
		func(w http.ResponseWriter, r *http.Request) {
			address, e := cipher.DecodeBase58Address(path.Base(r.URL.EscapedPath()))
			if e != nil {
				sendErr(w, e)
				return
			}
			info := g.BlockChain.GetAddressInfo(address)
			sendOK(w, AddressReply{
				Address:      address.String(),
				Kitties:      info.Kitties,
				Transactions: []string{"TO_BE_IMPLEMENTED"},
			})
		})

	return nil
}

type Error struct {
	Msg string `json:"message"`
}

type Response struct {
	Data  interface{} `json:"data,omitempty"`
	Error *Error      `json:"error,omitempty"`
}

//func send(w http.ResponseWriter) func(v interface{}, e error) error {
//	return func(v interface{}, e error) error {
//		if e != nil {
//			return sendErr(w, e)
//		}
//		return sendOK(w, v)
//	}
//}

func sendOK(w http.ResponseWriter, v interface{}) error {
	response := Response{Data: v}
	return sendWithStatus(w, response, http.StatusOK)
}

func sendErr(w http.ResponseWriter, e error) error {
	// TODO (evanlinjin): Implement way to determine http status approprite for error.
	response := Response{
		Error: &Error{
			Msg: e.Error(),
		},
	}
	return sendWithStatus(w, response, http.StatusBadRequest)
}

func sendWithStatus(w http.ResponseWriter, v interface{}, status int) error {
	data, e := json.Marshal(v)
	if e != nil {
		return e
	}
	sendRaw(w, data, status)
	return nil
}

func sendRaw(w http.ResponseWriter, data []byte, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(data)
}
