package zenrpc

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

var noopMiddleware = func(next InvokeFunc) InvokeFunc {
	return func(c Context, method string, params json.RawMessage) Response {
		return next(c, method, params)
	}
}

func TestGroup(t *testing.T) {
	s := NewServer(Options{})
	g := s.Group(noopMiddleware)

	assert.Len(t, g.middleware, 1)

	g = s.Group()
	assert.Len(t, g.middleware, 0)
}

func TestGroupUse(t *testing.T) {
	s := NewServer(Options{})

	g := s.Group()
	g.Use(noopMiddleware)
	assert.Len(t, g.middleware, 1)
}
