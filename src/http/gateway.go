package http

import (
	"github.com/kittycash/iko/src/kchain"
	"net/http"
	"encoding/json"
	"strconv"
)

type Gateway struct {
	BlockChain *kchain.BlockChain
}

func (g *Gateway) host(mux *http.ServeMux) error {

	mux.HandleFunc("/api/kitty",
		func(w http.ResponseWriter, r *http.Request) {
			kittyID, e := strconv.ParseUint(r.URL.Query().Get("id"), 10, 64)
			if e != nil {
				sendErr(w, e)
			}
			send(w)(g.BlockChain.GetKittyAddress(kittyID))
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

func send(w http.ResponseWriter) func(v interface{}, e error) error {
	return func(v interface{}, e error) error {
		if e != nil {
			return sendErr(w, e)
		}
		return sendOK(w, v)
	}
}

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