package client

import (
	"io"
	"log"
	"mishazenin/PoW_server/src/hashcash"
	"net/http"
	"strconv"
)

type HttpClient struct {
	addr string
}

// New returns a new solver client.
func New(addr string) *HttpClient {
	return &HttpClient{addr: addr}
}

// GetQuote returns the quote after solving challenge.
func (c *HttpClient) GetQuote() {
	resp, err := http.Get(c.addr)
	if err != nil {
		log.Fatal("could not get challenge:", err)
	}
	defer resp.Body.Close()

	challenge := resp.Header.Get("X-Hashcash")
	vals, err := hashcash.Split(challenge)
	if err != nil {
		log.Fatal(err)
	}

	bits, err := strconv.ParseInt(vals[1], 10, 0)
	if err != nil {
		log.Fatal("invalid bit length:", err)
	}

	solution := hashcash.New(bits).Solve(challenge)
	req, err := http.NewRequest(http.MethodGet, c.addr, nil)
	if err != nil {
		log.Fatal("new request err:", err)
	}

	req.Header.Add("X-Hashcash", solution)

	respSolve, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("check challenge err:", err)
	}
	if respSolve.StatusCode != http.StatusOK {
		log.Fatal("check challenge not ok:", resp.StatusCode)
	}

	quote, err := io.ReadAll(respSolve.Body)
	if err != nil {
		log.Fatal("read quote err:", err)
	}
	defer respSolve.Body.Close()

	log.Println(string(quote))
}
