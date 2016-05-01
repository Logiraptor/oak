package codegen

import (
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestGoFmt(t *testing.T) {

	assert.Panics(t, func() {
		formatString(`
not a go file
`)
	})

	assert.Equal(t, formatString(`
package main

func main  (  ) {
	i := 0
	_ = i
	}

`), `package main

func main() {
	i := 0
	_ = i
}
`)

}
