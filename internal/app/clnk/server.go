package clnk

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Server ...
type Server struct {
	router *mux.Router
}

// NewServer ...
func NewServer() *Server {
	s := &Server{
		router: mux.NewRouter(),
	}
	s.configureRouter()
	return s
}

// ServeHTTP ...
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *Server) configureRouter() {
	s.router.HandleFunc("/ping", s.ping()).Methods("GET")
}

func (s *Server) ping() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.responder(w, r, http.StatusTeapot, nil)
	}
}

func (s *Server) responder(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
