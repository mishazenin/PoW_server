package hashcash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidation(t *testing.T) {
	h := New(12)
	challenge, err := h.Challenge("1.2.3.4")
	assert.NoError(t, err)
	assert.True(t, h.Validate(h.Solve(challenge)))
}
