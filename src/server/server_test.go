package server

import (
	"github.com/LarsFox/pow-tcp/src/library"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

const (
	challenge = "MEGAchallenge"
	solution  = "MEGAsolution"
)

var (
	letters = []string{"q", "w", "e", "r", "t", "y", "u", "i", "o", "p"}
	book    = library.NewBook(letters)
)

func TestSolveChallenge(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	validator := NewMockvalidator(ctrl)
	s := NewPOWServer(book, validator)
	server := httptest.NewServer(http.HandlerFunc(s.handle))
	defer server.Close()

	// Getting a challenge.
	validator.EXPECT().Challenge(gomock.Any()).Return(challenge, nil)
	resp, err := http.Get(server.URL)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	xchallenge := resp.Header.Get("X-Hashcash")
	assert.Equal(t, challenge, xchallenge)

	// Solving the challenge.
	req, err := http.NewRequest(http.MethodGet, server.URL, nil)
	assert.NoError(t, err)
	req.Header.Add("X-Hashcash", solution)
	validator.EXPECT().Validate(solution).Return(true)

	respSolve, err := http.DefaultClient.Do(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, respSolve.StatusCode)

	quote, err := io.ReadAll(respSolve.Body)
	assert.NoError(t, err)
	defer respSolve.Body.Close()

	assert.Contains(t, letters, string(quote))
}

func TestFailChallenge(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	validator := NewMockvalidator(ctrl)
	s := NewPOWServer(book, validator)
	server := httptest.NewServer(http.HandlerFunc(s.handle))
	defer server.Close()

	// Getting a challenge.
	validator.EXPECT().Challenge(gomock.Any()).Return(challenge, nil)
	resp, err := http.Get(server.URL)
	assert.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	xchallenge := resp.Header.Get("X-Hashcash")
	assert.Equal(t, challenge, xchallenge)

	// Failing the challenge.
	req, err := http.NewRequest(http.MethodGet, server.URL, nil)
	assert.NoError(t, err)
	req.Header.Add("X-Hashcash", "WRONG_SOLUTION")
	validator.EXPECT().Validate("WRONG_SOLUTION").Return(false)

	respSolve, err := http.DefaultClient.Do(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusForbidden, respSolve.StatusCode)
}
