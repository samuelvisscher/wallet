package http

import (
	"encoding/json"
	"fmt"
	"github.com/kittycash/wallet/src/iko"
	"net/http"
)

type Gateway struct {
	IKO *iko.BlockChain
}

func (g *Gateway) host(mux *http.ServeMux) error {

	if g.IKO != nil {
		if e := ikoGateway(mux, g.IKO); e != nil {
			return e
		}
	}

	return nil
}

/*
	<<< ACTION >>>
*/

type httpAction func(w http.ResponseWriter, r *http.Request) error

func Do(action httpAction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if e := action(w, r); e != nil {
			fmt.Println(e)
		}
	}
}

/*
	<<< RETURN SPECIFICATIONS >>>
*/

type Error struct {
	Msg string `json:"message"`
}

type Response struct {
	Data  interface{} `json:"data,omitempty"`
	Error *Error      `json:"error,omitempty"`
}

// send200 - OK.
func send200(w http.ResponseWriter, v interface{}) error {
	response := Response{Data: v}
	return sendWithStatus(w, response, http.StatusOK)
}

// send400 bad request.
func send400(w http.ResponseWriter, e error) error {
	return sendErrWithStatus(w, e, http.StatusBadRequest)
}

// send404 not found.
func send404(w http.ResponseWriter, e error) error {
	return sendErrWithStatus(w, e, http.StatusNotFound)
}

// send405 method not allowed.
func send405(w http.ResponseWriter, got, exp string) error {
	e := fmt.Errorf("got method '%s' while expecting '%s'", got, exp)
	return sendErrWithStatus(w, e, http.StatusNotFound)
}

func sendErrWithStatus(w http.ResponseWriter, e error, status int) error {
	response := Response{
		Error: &Error{
			Msg: e.Error(),
		},
	}
	return sendWithStatus(w, response, status)
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

func sendBin(w http.ResponseWriter, data []byte, status int) error {
	w.Header().Set("Content-Type", "application/octet-stream")
	w.WriteHeader(status)
	w.Write(data)
	return nil
}