package http

import (
	"encoding/json"
	"fmt"
	"github.com/kittycash/wallet/src/wallet"
	"net/http"
)

func walletGateway(mux *http.ServeMux, g *wallet.Manager) error {

	Handle(mux, "/api/wallets/refresh",
		"GET", refreshWallets(g))

	Handle(mux, "/api/wallets/list",
		"GET", listWallets(g))

	Handle(mux, "/api/wallets/new",
		"POST", newWallet(g))

	return nil
}

func refreshWallets(g *wallet.Manager) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, p *Path) error {
		// TODO: implement.
		// CALLS: 'g.Refresh()'.
		// RESPONSE:
		// - 'Content-Type': 'application/json'.
		// - true & 200   : on success.
		// - string & 500 : of the error on failure.

		err := g.Refresh()

		// Send json response with 500 status code if error
		if err != nil {
			sendJson(w, http.StatusInternalServerError,
				fmt.Sprintf("Message: '%s'", err))
		}

		// Send json response with 200 status code if error is nil
		return sendJson(w, http.StatusOK, true)
	}
}

type WalletsReply struct {
	Wallets []wallet.Stat `json:"wallets"`
}

func listWallets(g *wallet.Manager) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, p *Path) error {
		// TODO: implement.
		// CALLS: 'g.ListWallets()'.
		// RESPONSE:
		// - 'Content-Type': 'application/json'.
		// - json representation of 'WalletsReply'.

		// Get list of listWall
		listWall WalletsReply
		listWall.Wallets = g.ListWallets()

		if listWall.Wallets != nil {
			// Send json response, status= 200 to user if content was found
			return sendJson(w, http.StatusOK, listWall.Wallets)
		} else {
			// Send json response, status= 204 to user if no content was found
			return sendJson(w, http.StatusNoContent, listWall)
		}
	}
}

func newWallet(g *wallet.Manager) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, p *Path) error {
		// TODO: implement.
		// CALLS: 'g.NewWallet()'.
		// REQUEST:
		// - 'Content-Type': 'application/x-www-form-urlencoded'.
		// - Check actual function for all key-values.
		// RESPONSE:
		// - 'Content-Type': 'application/json'.
		// - true & 200   : on success.
		// - string & http status code : of the error on failure.
		return nil
	}
}

func sendJson(w http.ResponseWriter, status int, v interface{}) error {
	data, e := json.Marshal(v)
	if e != nil {
		return e
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, e = w.Write(data)
	return e
}
