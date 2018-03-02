package http

import (
	"fmt"
	"github.com/kittycash/wallet/src/wallet"
	"mime"
	"strings"
	"net/http"
	"strconv"
)

func walletGateway(mux *http.ServeMux, g *wallet.Manager) error {

	Handle(mux, "/api/wallets/refresh",
		"GET", refreshWallets(g, "GET"))

	Handle(mux, "/api/wallets/list",
		"GET", listWallets(g, "GET"))

	Handle(mux, "/api/wallets/new",
		"POST", newWallet(g, "POST"))

	return nil
}

func refreshWallets(g *wallet.Manager, method string) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, p *Path) error {
		// TODO: implement.
		// CALLS: 'g.Refresh()'.
		// RESPONSE:
		// - 'Content-Type': 'application/json'.
		// - true & 200   : on success.
		// - string & 500 : of the error on failure.

		// Check request method
		if r.Method != method {
			sendJson(w, http.StatusBadRequest,
				fmt.Sprintf("Invalid request type. Expected %s but got %s",
						r.Method, method))
		}

		err := g.Refresh()

		// Send json response with 500 status code if error
		if err != nil {
			sendJson(w, http.StatusInternalServerError,
				fmt.Sprintf("Message: %s", err))
		}

		// Send json response with 200 status code if error is nil
		return sendJson(w, http.StatusOK, true)
	}
}

type WalletsReply struct {
	Wallets []wallet.Stat `json:"wallets"`
}

func listWallets(g *wallet.Manager, method string) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, p *Path) error {
		// TODO: implement.
		// CALLS: 'g.ListWallets()'.
		// RESPONSE:
		// - 'Content-Type': 'application/json'.
		// - json representation of 'WalletsReply'.

		// Check request method
		if r.Method != method {
			sendJson(w, http.StatusBadRequest,
				fmt.Sprintf("Invalid request type. Expected %s but got %s",
						r.Method, method))
		}
		// Get list of listWall
		var listWall WalletsReply
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

func newWallet(g *wallet.Manager, method string) HandlerFunc {
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

		if r.Method != method {
			sendJson(w, http.StatusBadRequest,
				fmt.Sprintf("Invalid request type. Expected %s but got %s",
						r.Method, method))
		}

		mimetype := "application/x-www-form-urlencoded"
		isRightContent := hasContentType(r, mimetype)
		if isRightContent {
			sendJson(w, http.StatusUnsupportedMediaType,
				fmt.Sprintf("Expecting Content-Type to be %s but got %s",
					mimetype, r.Header.Get("Content-Type")))
		}

		// Send json response if body is nil
		if r.Body == nil {
			return sendJson(w, http.StatusBadRequest,
					fmt.Sprintf("Request body missing"))
		}

		// Parse form data
		err := r.ParseForm()

		if err != nil {
			return sendJson(w, http.StatusBadRequest,
					fmt.Sprintf("Error: %s", err))
		}

		encrypted, err := strconv.ParseBool(r.PostFormValue("encrypted"))
		// Options to pass to g.NewWallet()
		opts := wallet.Options{
			Label: r.PostFormValue("label"),
			Seed: r.PostFormValue("seed"),
			Encrypted: encrypted,
			Password: r.PostFormValue("password"),
		}

		// Verify that all values are correct
		err = opts.Verify()

		// Respond if options are not correct
		if err != nil {
			sendJson(w, http.StatusBadRequest,
				fmt.Sprintf("Message: %s", err))
		}

		// Get addresses and convert it to int
		addr, err := strconv.Atoi(r.PostFormValue("addresses"))

		// Don't allow anything other than int
		if err != nil || addr < 1 {
			sendJson(w, http.StatusNotAcceptable,
				fmt.Sprintf("Error: ", err))
		}

		// Call g.NewWallet
		err = g.NewWallet(&opts, addr)

		if err != nil {
			sendJson(w, http.StatusInternalServerError,
				fmt.Sprintf("Error: %s", err))
		}

		return sendJson(w, http.StatusOK, true)
	}
}

func hasContentType(r *http.Request, mimetype string) bool {
	contentType := r.Header.Get("Content-type")
	if contentType == "" {
		return mimetype == "application/octet-stream"
	}

	for _, v := range strings.Split(contentType, ",") {
		t, _, err := mime.ParseMediaType(v)
		if err != nil {
			break
		}
		if t == mimetype {
			return true
		}
	}
	return false
}
