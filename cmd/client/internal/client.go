package internal

import (
	"io"
	"log"
	"net/http"
	"strconv"

	"mishazenin/PoW_server/pkg"
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
		log.Fatalf("could not get challenge: %w", err)
	}
	defer resp.Body.Close()

	challenge := resp.Header.Get("X-Hashcash")
	vals, err := pkg.Split(challenge)
	if err != nil {
		log.Fatal(err)
	}

	bits, err := strconv.ParseInt(vals[1], 10, 0)
	if err != nil {
		log.Fatalf("invalid bit length:%w", err)
	}

	solution := pkg.New(bits).Solve(challenge)
	req, err := http.NewRequest(http.MethodGet, c.addr, nil)
	if err != nil {
		log.Fatalf("new request err: %w", err)
	}

	req.Header.Add("X-Hashcash", solution)

	respSolve, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("check challenge err: %w", err)
	}
	if respSolve.StatusCode != http.StatusOK {
		log.Fatalf("check challenge not ok: %d", resp.StatusCode)
	}

	quote, err := io.ReadAll(respSolve.Body)
	if err != nil {
		log.Fatalf("read quote err: %w", err)
	}
	defer respSolve.Body.Close()

	log.Println(string(quote))
}
