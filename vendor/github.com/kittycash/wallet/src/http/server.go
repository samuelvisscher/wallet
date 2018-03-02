package http

import (
	"net/http"
	"time"
)

type ServerConfig struct {
	Address     string
	EnableTLS   bool
	TLSCertFile string
	TLSKeyFile  string
}

type Server struct {
	c    *ServerConfig
	srv  *http.Server
	mux  *http.ServeMux
	api  *Gateway
	quit chan struct{}
}

func NewServer(config *ServerConfig, api *Gateway) (*Server, error) {
	var server = &Server{
		c:    config,
		mux:  http.NewServeMux(),
		api:  api,
		quit: make(chan struct{}),
	}
	if e := server.prepareMux(); e != nil {
		return nil, e
	}
	go server.serve()
	return server, nil
}

func (s *Server) serve() {
	s.srv = &http.Server{
		Addr:    s.c.Address,
		Handler: s.mux,
	}
	if s.c.EnableTLS {
		for {
			if e := s.srv.ListenAndServeTLS(s.c.TLSCertFile, s.c.TLSKeyFile); e != nil {
				time.Sleep(100 * time.Millisecond)
				continue
			} else {
				break
			}
		}
	} else {
		for {
			if e := s.srv.ListenAndServe(); e != nil {
				time.Sleep(100 * time.Millisecond)
				continue
			} else {
				break
			}
		}
	}
	s.srv = nil
}

func (s *Server) prepareMux() error {
	return s.api.host(s.mux)
}

// Close quits the http server.
func (s *Server) Close() {
	if s.quit != nil {
		close(s.quit)
		if s.srv != nil {
			s.srv.Close()
		}
	}
}
