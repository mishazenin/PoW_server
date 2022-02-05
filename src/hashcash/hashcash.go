package hashcash

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"math"
	"math/big"
	"net"
	"strings"
	"time"
)

const (
	Bytes          = 8
	hashcashLength = 7
	TimeFormat     = "2006-01-02"
)

// Hashash generates and validates hashcashes.
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
func Split(s string) []string {
	vals := strings.Split(s, ":")
	if len(vals) != hashcashLength {
		return nil
	}
	return vals
}

// Challenge returns a new hashcash string with the missing nonce
// that serves as a base for Proof-of-work algorhythm, e.g.
// 1:20:060102150405:1.2.3.4::McMybZIhxKXu57jd:
//
// For details of Hashcash see https://en.wikipedia.org/wiki/Hashcash
func (h *Hashcash) Challenge(ip string) (string, error) {
	bytes := make([]byte, Bytes)

	random := base64.StdEncoding.EncodeToString(bytes)
	t := time.Now().UTC().Format(TimeFormat)
	return fmt.Sprintf("1:%d:%s:%s::%s:", h.bits, t, ip, random), nil
}

// Solve bruteforces a challenge to find the matching nonce.
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

// Validate checks if s is a valid hashcash value that meets specified target.
func (h *Hashcash) Validate(s string) bool {
	vals := Split(s)
	if vals == nil {
		return false
	}

	if !h.validateHash(s) {
		return false
	}

	if !h.validateTime(vals[2]) {
		return false
	}
	return net.ParseIP(vals[3]) != nil
}

func (h *Hashcash) validateHash(s string) bool {
	return big.NewInt(0).SetBytes([]byte(s)).Cmp(h.target) != -1
}

func (h *Hashcash) validateTime(val string) bool {
	date, err := time.Parse(TimeFormat, val)
	if err != nil {
		return false
	}
	if date.After(time.Now()) {
		return false
	}
	return date.After(time.Now().AddDate(0, 0, -1))
}
