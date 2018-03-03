package http

import (
	"fmt"
	"github.com/kittycash/wallet/src/wallet"
	"net/http"
	"strconv"
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

		// Send json response with 500 status code if error
		if e := g.Refresh(); e != nil {
			sendJson(w, http.StatusInternalServerError,
				fmt.Sprintf("Message: '%s'", e))
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
		var listWall WalletsReply
		listWall.Wallets = g.ListWallets()

		if listWall.Wallets != nil {
			// Send json response, status= 200 to user if content was found
			return sendJson(w, http.StatusOK, listWall)
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

		// Send json response if body is nil
		if r.Body == nil {
			return sendJson(w, http.StatusBadRequest,
					fmt.Sprint("Request body missing"))
		}

		if e := r.ParseForm(); e != nil {
			return sendJson(w, http.StatusBadRequest,
					fmt.Sprintf("Error: %s", e))
		}

		encrypted, err := strconv.ParseBool(r.PostFormValue("encrypted"))

		if err != nil {
			sendJson(w, http.StatusBadRequest,
			fmt.Sprintf("Error: %s", err))
		}
		// Options to pass to g.NewWallet()
		opts := wallet.Options{
			Label: r.PostFormValue("label"),
			Seed: r.PostFormValue("seed"),
			Encrypted: encrypted,
			Password: r.PostFormValue("password"),
		}

		/**
		 * Verify that all values are correct
		 * Respond if options are not correct
		 */
		if e := opts.Verify; e != nil {
			sendJson(w, http.StatusBadRequest,
				fmt.Sprintf("Message: %s", err))
		}

		// Get addresses and convert it to int
		addr, err := strconv.Atoi(r.PostFormValue("addresses"))

		// Don't allow anything other than int
		if err != nil || addr < 1 {
			sendJson(w, http.StatusNotAcceptable,
				fmt.Sprintf("Error: %s", err))
		}

		if e := g.NewWallet(&opts, addr); e != nil {
			sendJson(w, http.StatusInternalServerError,
				fmt.Sprintf("Error: %s", e))
		}

		return sendJson(w, http.StatusOK, true)
	}
}
