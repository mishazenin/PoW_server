package server

import (
	"github.com/LarsFox/pow-tcp/src/library"
	"log"
	"net"
	"net/http"
)

const (
	hashcashHeader = "X-Hashcash"
)

////go:generate mockgen -destination=server_validator_mock_test.go -source=server.go -package=server
type validator interface {
	Challenge(ip string) (string, error)
	Validate(solution string) bool
}

// POWServer is a simple Proof-of-Work server implementation.
type POWServer struct {
	book      *library.Library
	validator validator
}

// NewPOWServer returns new Proof-of-Work server.
func NewPOWServer(book *library.Library, validator validator) *POWServer {
	return &POWServer{
		book:      book,
		validator: validator,
	}
}

// Listen listens on TCP network.
func (s *POWServer) Listen(addr string) {
	log.Fatal(http.ListenAndServe(addr, http.HandlerFunc(s.PoWHandler)))
}

func (s *POWServer) GetIP(r *http.Request) string {
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "9.9.9.9"
	}
	return ip
}

func (s *POWServer) PoWHandler(w http.ResponseWriter, r *http.Request) {
	hashcash := r.Header.Get(hashcashHeader)
	if hashcash == "" {
		s.handleNewChallenge(w, r)
		return
	}

	val := s.validator.Validate(hashcash)
	if !val {
		s.sendErr(w, http.StatusForbidden)
		return
	}

	quote, err := s.book.RandomLine()
	if err != nil {
		s.sendErr(w, http.StatusInternalServerError)
		return
	}

	log.Println("serving quote", string(quote))
	w.Write(quote)
}

func (s *POWServer) handleNewChallenge(w http.ResponseWriter, r *http.Request) {
	ip := s.GetIP(r)
	challenge, err := s.validator.Challenge(ip)
	if err != nil {
		s.sendErr(w, http.StatusInternalServerError)
		return
	}

	log.Println("serving challenge", challenge)
	w.Header().Set(hashcashHeader, challenge)
	w.WriteHeader(http.StatusOK)
}

func (s *POWServer) sendErr(w http.ResponseWriter, code int) {
	w.WriteHeader(code)
}
