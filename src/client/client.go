package client

import (
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/LarsFox/pow-tcp/src/hashcash"
)

type HttpClient struct {
	addr string
}

// New returns a new challenge getter-solver client.
func New(addr string) *HttpClient {
	return &HttpClient{addr: addr}
}

// Quote returns solves the challenge and returns the quote.
func (c *HttpClient) Quote() {
	// Getting the challenge.
	resp, err := http.Get(c.addr)
	if err != nil {
		log.Fatal("some error in challenge:", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatal("get challenge not ok:", resp.StatusCode)
	}

	challenge := resp.Header.Get("X-Hashcash")
	vals := hashcash.Split(challenge)
	if vals == nil {
		log.Fatal("invalid challenge value:", challenge)
	}

	bits, err := strconv.ParseInt(vals[1], 10, 0)
	if err != nil {
		log.Fatal("invalid bit length:", err)
	}

	// Checking the solution.
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
