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

func ikoGateway(mux *http.ServeMux, g *iko.BlockChain) error {

	MultiHandle(mux, Do(3, "GET", getKitty(g)),
		[]string{
			"/api/iko/kitty/",
			"/api/iko/kitty.json/",
			"/api/iko/kitty.bin/",
		})
	MultiHandle(mux, Do(3, "GET", getAddress(g)),
		[]string{
			"/api/iko/address/",
			"/api/iko/address.json/",
			"/api/iko/address.bin/",
		})
	MultiHandle(mux, Do(3, "GET", getTx(g)),
		[]string{
			"/api/iko/tx/",
			"/api/iko/tx.json/",
			"/api/iko/tx.bin/",
		})
	MultiHandle(mux, Do(3, "GET", getHeadTx(g)),
		[]string{
			"/api/iko/head_tx",
			"/api/iko/head_tx.json",
			"/api/iko/head_tx.bin",
		})
	MultiHandle(mux, Do(3, "GET", getHeadTx(g)),
		[]string{
			"/api/iko/inject_tx",
			"/api/iko/inject_tx.json",
			"/api/iko/inject_tx.bin",
		})

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
				fmt.Sprintf("kitty of id '%s' not found", kittyID))
		}
		return SwitchExtension(w, p,
			func() error {
				return sendJson(w, http.StatusOK,
					KittyReply{
						KittyID:      kittyID,
						Address:      kState.Address.String(),
						Transactions: kState.Transactions.ToStringArray(),
					})
			},
			func() error {
				return sendBin(w, http.StatusOK,
					kState.Serialize())
			},
		)
	}
}

type AddressReply struct {
	Address      string       `json:"address"`
	Kitties      iko.KittyIDs `json:"kitties"`
	Transactions []string     `json:"transactions"`
}

func getAddress(g *iko.BlockChain) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, p *Path) error {
		address, e := cipher.DecodeBase58Address(p.Segment(p.BasePos + 1))
		if e != nil {
			return sendJson(w, http.StatusBadRequest,
				e.Error())
		}
		aState := g.GetAddressState(address)
		return SwitchExtension(w, p,
			func() error {
				return sendJson(w, http.StatusOK,
					AddressReply{
						Address:      address.String(),
						Kitties:      aState.Kitties,
						Transactions: aState.Transactions.ToStringArray(),
					})
			},
			func() error {
				return sendBin(w, http.StatusOK,
					aState.Serialize())
			},
		)
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

func getTx(g *iko.BlockChain) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, p *Path) error {
		var tx iko.Transaction
		switch reqVal := r.URL.Query().Get("request"); reqVal {
		case "", "hash":
			txHash, e := cipher.SHA256FromHex(p.Segment(p.BasePos + 1))
			if e != nil {
				return sendJson(w, http.StatusBadRequest,
					e.Error())
			}
			if tx, e = g.GetTxOfHash(iko.TxHash(txHash)); e != nil {
				return sendJson(w, http.StatusNotFound,
					e.Error())
			}
		case "seq":
			seq, e := strconv.ParseUint(p.Segment(p.BasePos+1), 10, 64)
			if e != nil {
				return sendJson(w, http.StatusBadRequest,
					e.Error())
			}
			if tx, e = g.GetTxOfSeq(seq); e != nil {
				return sendJson(w, http.StatusNotFound,
					e.Error())
			}
		default:
			return sendJson(w, http.StatusBadRequest,
				fmt.Sprintf("invalid request query value of '%s', expected '%s'",
					reqVal, []string{"", "hash", "seq"}))
		}
		return SwitchExtension(w, p,
			func() error {
				return sendJson(w, http.StatusOK,
					TxReply{
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
			},
			func() error {
				return sendBin(w, http.StatusOK,
					tx.Serialize())
			},
		)
	}
}

type HeadHashReply struct {
	Seq  uint64 `json:"seq"`
	Hash string `json:"hash"`
}

func getHeadTx(g *iko.BlockChain) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, p *Path) error {
		tx, e := g.GetHeadTx()
		if e != nil {
			return sendJson(w, http.StatusNotFound,
				e.Error())
		}
		return SwitchExtension(w, p,
			func() error {
				return sendJson(w, http.StatusOK, TxReply{
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
			},
			func() error {
				return sendBin(w, http.StatusOK,
					tx.Serialize())
			},
		)
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
		return SwitchExtension(w, p,
			func() error {
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
				tx := new(iko.Transaction)
				if e := encoder.DeserializeRaw(hexRaw, tx); e != nil {
					return sendJson(w, http.StatusBadRequest,
						e.Error())
				}
				if e := g.InjectTx(tx); e != nil {
					return sendJson(w, http.StatusBadRequest,
						e.Error())
				}
				return sendJson(w, http.StatusOK,
					true)
			},
			func() error {
				tx := new(iko.Transaction)
				if e := encoder.DeserializeRaw(txRaw, tx); e != nil {
					return sendJson(w, http.StatusBadRequest,
						e.Error())
				}
				if e := g.InjectTx(tx); e != nil {
					return sendJson(w, http.StatusBadRequest,
						e.Error())
				}
				return sendJson(w, http.StatusOK,
					true)
			},
		)
	}
}
