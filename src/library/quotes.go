package library

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

// Book is a storage for quotes.
type Book struct {
	quotes []string
}

// NewBook returns a new book with quotes.
func NewBook(quotes []string) *Book {
	return &Book{quotes: quotes}
}

// Random returns a random quote from the book.
func (b *Book) RandomLine() ([]byte, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(len(b.quotes))))
	if err != nil {
		return nil, fmt.Errorf("random quote getting err: %w", err)
	}
	return []byte(b.quotes[n.Int64()]), nil
}

// ITQuotes book with quotes
var ITQuotes = &Book{
	quotes: []string{
		`First, solve the problem. Then, write the code`,
		`Any fool can write code that a computer can understand. Good programmers write code that humans can understand`,
		`Experience is the name everyone gives to their mistakes`,
		`In order to be irreplaceable, one must always be different`,
		`Java is to JavaScript what car is to Carpet`,
		`Knowledge is power`,
		`Sometimes it pays to stay in bed on Monday, rather than spending the rest of the week debugging Monday’s code`,
		`Perfection is achieved not when there is nothing more to add, but rather when there is nothing more to take away`,
		`Ruby is rubbish! PHP is phpantastic!`,
		`Code is like humor. When you have to explain it, it’s bad`,
		`Fix the cause, not the symptom`,
		`Optimism is an occupational hazard of programming: feedback is the treatment`,
		`When to use iterative development? You should use iterative development only on projects that you want to succeed`,
		`Simplicity is the soul of efficiency`,
		`Before software can be reusable it first has to be usable`,
		`Make it work, make it right, make it fast`},
}
