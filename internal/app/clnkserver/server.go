package clnkserver

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/imflop/clnk/internal/app/converter"
	"github.com/imflop/clnk/internal/app/store"
)

// server ...
type server struct {
	router *mux.Router
	store  store.Store
}

// ServeHTTP ...
func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

// NewServer ...
func NewServer(store store.Store) *server {
	s := &server{
		router: mux.NewRouter(),
		store:  store,
	}
	s.configureRouter()
	return s
}

func (s *server) configureRouter() {
	s.router.HandleFunc("/teapot", s.teapot()).Methods(http.MethodGet)
	s.router.HandleFunc("/links/{id:[0-9]+}", s.getLink()).Methods(http.MethodGet)
	s.router.HandleFunc("/", s.shortenLink()).Methods(http.MethodPost)
	s.router.HandleFunc("/{encodedURL}", s.unshortenLink()).Methods(http.MethodGet, http.MethodHead)
}

func (s *server) teapot() http.HandlerFunc {
	type Teapot struct {
		IsTeapot bool `json:"is_teapot"`
	}
	t := &Teapot{}
	t.IsTeapot = true
	return func(w http.ResponseWriter, r *http.Request) {
		s.responder(w, r, http.StatusTeapot, &t)
	}
}

func (s *server) shortenLink() http.HandlerFunc {
	type request struct {
		OriginalURL string `json:"original_url"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		l, err := s.store.Link().Create(req.OriginalURL)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}
		id := l.ID
		encoded, _ := converter.Encode(id)

		shortURL := s.checkScheme(r).URL.Scheme + "://" + r.Host + "/" + encoded
		l, err = s.store.Link().Update(id, shortURL)
		if err != nil {
			s.responder(w, r, http.StatusInternalServerError, err)
			return
		}

		s.responder(w, r, http.StatusOK, l)
	}
}

func (s *server) unshortenLink() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uri := mux.Vars(r)["encodedURL"]
		linkID, err := converter.Decode(uri)
		if err != nil {
			s.error(w, r, http.StatusNotFound, err)
			return
		}
		l, err := s.store.Link().Find(linkID)
		if err != nil {
			s.error(w, r, http.StatusNotFound, err)
			return
		}
		http.Redirect(w, r, l.OriginalURL, http.StatusMovedPermanently)
	}
}

func (s *server) getLink() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
		}
		l, err := s.store.Link().Find(id)
		if err != nil {
			s.error(w, r, http.StatusNotFound, err)
			return
		}
		s.responder(w, r, http.StatusOK, l)
	}
}

func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.responder(w, r, code, map[string]string{"error": err.Error()})
}

func (s *server) responder(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

func (s *server) checkScheme(r *http.Request) *http.Request {
	if r.TLS != nil {
		r.URL.Scheme = "https"
	} else {
		r.URL.Scheme = "http"
	}
	return r
}
