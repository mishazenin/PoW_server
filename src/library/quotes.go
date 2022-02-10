package library

import (
	"math/rand"
	"time"
)

type Book struct {
	quotes []string
}

func NewBook(quotes []string) *Book {
	return &Book{quotes: quotes}
}

// RandomLine returns a random quote
func (b *Book) RandomLine() ([]byte, error) {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(len(b.quotes)+1) + 1
	return []byte(b.quotes[int64(n)]), nil
}

// Quotes book with quotes
var ItQuotes = &Book{
	quotes: []string{
		`First, solve the problem. Then, write the code`,
		`Any fool can write code that a computer can understand. Good programmers write code that humans can understand`,
		`Experience is the name everyone gives to their mistakes`,
		`In order to be irreplaceable, one must always be different`,
		`Java is to JavaScript what car is to Carpet`,
		`Knowledge is power`,
		`Sometimes it pays to stay in bed on Monday, rather than spending the rest of the week debugging Monday’s code`,
		`Perfection is achieved not when there is nothing more to add, but rather when there is nothing more to take away`,
		`Ruby is rubbish! PHP is pantheistic!`,
		`Code is like humor. When you have to explain it, it’s bad`,
		`Fix the cause, not the symptom`,
		`Optimism is an occupational hazard of programming: feedback is the treatment`,
		`When to use iterative development? You should use iterative development only on projects that you want to succeed`,
		`Simplicity is the soul of efficiency`,
		`Before software can be reusable it first has to be usable`,
		`Make it work, make it right, make it fast`},
}
