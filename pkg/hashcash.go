package pkg

import (
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"math"
	"math/big"
	"net"
	"strings"
	"time"
)

const (
	Bytes          = 80
	hashcashLength = 7
)

// Hashcash generates and validates hashcashes
type Hashcash struct {
	bits   int64
	target *big.Int
}

// New returns a new Hashcash generator and validator.
func New(bits int64) *Hashcash {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-bits))
	return &Hashcash{
		bits:   bits,
		target: target,
	}
}

// Split returns values of a hashcash challenge as an array.
// In case of array length mismatch, returns nil.
func Split(s string) ([]string, error) {
	vals := strings.Split(s, ":")
	if len(vals) != hashcashLength {
		return nil, errors.New("invalid hashcash length")
	}
	return vals, nil
}

// Constructor returns new hashcash string with the missing nonce
func (h *Hashcash) Constructor(host string) (string, error) {
	bytes := make([]byte, Bytes)
	random := base64.StdEncoding.EncodeToString(bytes)
	return fmt.Sprintf("1:%d:%s:%s::%s:",
		h.bits,
		time.Now().UTC().Format("2006-01-02"),
		host,
		random), nil
}

// Solve bruteforces a challenge to find the matching nonce.
func (h *Hashcash) Solve(challenge string) string {

	for nonce := 0; nonce < math.MaxInt64; nonce++ {
		solution := fmt.Sprintf("%s%d", challenge, nonce)
		hash := sha256.Sum256([]byte(solution))
		if big.NewInt(0).SetBytes(hash[:]).Cmp(h.target) == -1 {
			return solution
		}
	}
	return ""
}

// Validate checks the hashcash is valid
func (h *Hashcash) Validate(s string) bool {
	vals, err := Split(s)
	if err != nil {
		log.Fatal(err)
	}

	if vals == nil || !h.validateHash(s) || !h.validateTime(vals[2]) || net.ParseIP(vals[3]) == nil {
		return false
	}

	return true
}

func (h *Hashcash) validateHash(s string) bool {
	return big.NewInt(0).SetBytes([]byte(s)).Cmp(h.target) != -1
}

func (h *Hashcash) validateTime(val string) bool {
	date, err := time.Parse("2006-01-02", val)
	if err != nil {
		return false
	}
	return date.After(time.Now().AddDate(0, 0, -1))
}
