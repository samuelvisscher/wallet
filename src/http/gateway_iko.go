package http

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/kittycash/wallet/src/iko"
	"github.com/skycoin/skycoin/src/cipher"
	"github.com/skycoin/skycoin/src/cipher/encoder"
	"io/ioutil"
	"net/http"
	"strconv"
)

func ikoGateway(m *http.ServeMux, g *iko.BlockChain) error {
	Handle(m, "/api/iko/kitty/", "GET", getKitty(g))
	Handle(m, "/api/iko/address/", "GET", getAddress(g))
	Handle(m, "/api/iko/tx/", "GET", getTx(g))
	Handle(m, "/api/iko/head_tx", "GET", getHeadTx(g))
	Handle(m, "/api/iko/txs", "GET", getPaginatedTxs(g))
	Handle(m, "/api/iko/inject_tx", "POST", injectTx(g))
	return nil
}

type KittyReply struct {
	KittyID      iko.KittyID `json:"kitty_id"`
	Address      string      `json:"address"`
	Transactions []string    `json:"transactions"`
}

func getKitty(g *iko.BlockChain) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, p *Path) error {
		kittyID, e := iko.KittyIDFromString(p.Base)
		if e != nil {
			return sendJson(w, http.StatusBadRequest,
				e.Error())
		}
		kState, ok := g.GetKittyState(kittyID)
		if !ok {
			return sendJson(w, http.StatusNotFound,
				fmt.Sprintf("kitty of id '%d' not found", kittyID))
		}
		return SwitchTypeQuery(w, r, TqJson, TypeQueryActions{
			TqJson: func() error {
				return sendJson(w, http.StatusOK,
					KittyReply{
						KittyID:      kittyID,
						Address:      kState.Address.String(),
						Transactions: kState.Transactions.ToStringArray(),
					})
			},
			TqEnc: func() error {
				return sendBin(w, http.StatusOK,
					kState.Serialize())
			},
		})
	}
}

type AddressReply struct {
	Address      string       `json:"address"`
	Kitties      iko.KittyIDs `json:"kitties"`
	Transactions []string     `json:"transactions"`
}

func getAddress(g *iko.BlockChain) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, p *Path) error {
		address, e := cipher.DecodeBase58Address(p.Base)
		if e != nil {
			return sendJson(w, http.StatusBadRequest,
				e.Error())
		}
		aState := g.GetAddressState(address)
		return SwitchTypeQuery(w, r, TqJson, TypeQueryActions{
			TqJson: func() error {
				return sendJson(w, http.StatusOK,
					AddressReply{
						Address:      address.String(),
						Kitties:      aState.Kitties,
						Transactions: aState.Transactions.ToStringArray(),
					})
			},
			TqEnc: func() error {
				return sendBin(w, http.StatusOK,
					aState.Serialize())
			},
		})
	}
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

func NewTxReplyOfTransaction(tx iko.Transaction) TxReply {
	return TxReply{
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
	}
}

func getTx(g *iko.BlockChain) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, p *Path) error {
		var tx iko.Transaction
		ok, e := SwitchReqQuery(w, r, RqHash, ReqQueryActions{
			RqHash: func() (bool, error) {
				txHash, e := cipher.SHA256FromHex(p.Base)
				if e != nil {
					return false, sendJson(w, http.StatusBadRequest,
						e.Error())
				}
				if tx, e = g.GetTxOfHash(iko.TxHash(txHash)); e != nil {
					return false, sendJson(w, http.StatusNotFound,
						e.Error())
				}
				return true, nil
			},
			RqSeq: func() (bool, error) {
				seq, e := strconv.ParseUint(p.Base, 10, 64)
				if e != nil {
					return false, sendJson(w, http.StatusBadRequest,
						e.Error())
				}
				if tx, e = g.GetTxOfSeq(seq); e != nil {
					return false, sendJson(w, http.StatusNotFound,
						e.Error())
				}
				return true, nil
			},
		})
		if !ok {
			return e
		}
		return SwitchTypeQuery(w, r, TqJson, TypeQueryActions{
			TqJson: func() error {
				return sendJson(w, http.StatusOK, NewTxReplyOfTransaction(tx))
			},
			TqEnc: func() error {
				return sendBin(w, http.StatusOK,
					tx.Serialize())
			},
		})
	}
}

func getHeadTx(g *iko.BlockChain) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, p *Path) error {
		tx, e := g.GetHeadTx()
		if e != nil {
			return sendJson(w, http.StatusNotFound,
				e.Error())
		}
		return SwitchTypeQuery(w, r, TqJson, TypeQueryActions{
			TqJson: func() error {
				return sendJson(w, http.StatusOK,
					NewTxReplyOfTransaction(tx))
			},
			TqEnc: func() error {
				return sendBin(w, http.StatusOK,
					tx.Serialize())
			},
		})
	}
}

type InjectTxRequest struct {
	Hex string `json:"hex"`
}

func injectTx(g *iko.BlockChain) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, p *Path) error {
		txRaw, e := ioutil.ReadAll(r.Body)
		if e != nil {
			return sendJson(w, http.StatusBadRequest,
				e.Error())
		}
		var tx = new(iko.Transaction)
		switch contentType := r.Header.Get("Content-Type"); contentType {
		case "application/json":
			req := new(InjectTxRequest)
			if e := json.Unmarshal(txRaw, req); e != nil {
				return sendJson(w, http.StatusBadRequest,
					e.Error())
			}
			hexRaw, e := hex.DecodeString(req.Hex)
			if e != nil {
				return sendJson(w, http.StatusBadRequest,
					e.Error())
			}
			if e := encoder.DeserializeRaw(hexRaw, tx); e != nil {
				return sendJson(w, http.StatusBadRequest,
					e.Error())
			}
		case "application/octet-stream":
			if e := encoder.DeserializeRaw(txRaw, tx); e != nil {
				return sendJson(w, http.StatusBadRequest,
					e.Error())
			}
		default:
			return sendJson(w, http.StatusBadRequest,
				fmt.Sprintf("content type '%s' is not supported, expecting '%s'",
					contentType, []string{"application/json", "application/octet-stream"}))
		}
		if e := g.InjectTx(tx); e != nil {
			return sendJson(w, http.StatusBadRequest,
				e.Error())
		}
		return sendJson(w, http.StatusOK,
			true)
	}
}

type PaginatedTxsReply struct {
	TotalPageCount uint64    `json:"total_page_count"`
	TxReplies      []TxReply `json:"transactions"`
}

func getPaginatedTxs(g *iko.BlockChain) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, p *Path) error {
		var (
			qCurrentPage = r.URL.Query().Get("current_page")
			qPerPage     = r.URL.Query().Get("per_page")
		)
		perPage, e := strconv.ParseUint(qPerPage, 10, 64)
		if e != nil {
			return sendJson(w, http.StatusBadRequest,
				e.Error())
		}
		currentPage, e := strconv.ParseUint(qCurrentPage, 10, 64)
		if e != nil {
			return sendJson(w, http.StatusBadRequest,
				e.Error())
		}
		paginated, e := g.GetTransactionPage(currentPage, perPage)
		if e != nil {
			return sendJson(w, http.StatusBadRequest,
				e.Error())
		}
		var txReplies []TxReply
		for _, transaction := range paginated.Transactions {
			txReplies = append(txReplies, NewTxReplyOfTransaction(transaction))
		}
		paginatedTxsReply := PaginatedTxsReply{
			TotalPageCount: paginated.TotalPageCount,
			TxReplies:      txReplies,
		}
		return sendJson(w, http.StatusOK,
			paginatedTxsReply)
	}
}
