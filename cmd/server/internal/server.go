package internal

import (
	"log"
	"net"
	"net/http"

	"mishazenin/PoW_server/pkg"
)

const hashcashHeader = "X-Hashcash"

// POWServer is a simple Proof-of-Work server implementation
type POWServer struct {
	book      *Book
	validator pkg.Hashcash
}

// NewPOWServer returns new Proof-of-Work server.
func NewPOWServer(book *Book, hc pkg.Hashcash) *POWServer {
	return &POWServer{
		book:      book,
		validator: hc,
	}
}

// Listen listens to TCP
func (s *POWServer) Listen(addr string) {
	log.Fatal(http.ListenAndServe(addr, http.HandlerFunc(s.PoWHandler)))
}

func (s *POWServer) PoWHandler(w http.ResponseWriter, r *http.Request) {
	hashcash := r.Header.Get(hashcashHeader)
	if hashcash == "" {
		s.handleNewChallenge(w, r)
		return
	}

	val := s.validator.Validate(hashcash)
	if !val {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	quote, err := s.book.RandomLine()
	if err != nil {
		log.Print("Could not get quote")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Println("generating quote", string(quote))
	w.Write(quote)
}

func (s *POWServer) handleNewChallenge(w http.ResponseWriter, r *http.Request) {
	host := GetHost(r)
	challenge, err := s.validator.Constructor(host)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Println("generating challenge", challenge)
	w.Header().Set(hashcashHeader, challenge)
	w.WriteHeader(http.StatusOK)
}

func GetHost(r *http.Request) string {
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "7.7.7.7"
	}
	return host
}
