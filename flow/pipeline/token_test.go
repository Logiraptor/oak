package pipeline

import "testing"
import "github.com/stretchr/testify/assert"

func TestNewToken(t *testing.T) {
	a := NewToken("a")
	a1 := NewToken("a")

	b := NewToken("b")

	assert.NotEqual(t, a, a1)
	assert.NotEqual(t, a, b)
}
