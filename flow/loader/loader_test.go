package loader

import (
	"fmt"
	"strings"

	"github.com/Logiraptor/oak/flow/parser"
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestLoader(t *testing.T) {
	prog := parser.Program{
		Entry: "start",
		Flow:  map[parser.ID]parser.ID{"start": "end"},
		Components: map[parser.ID]string{
			"start": "base.StringCLI",
			"end":   "func(s string) {}",
		},
		Imports: map[string]string{
			"base": "github.com/Logiraptor/oak/core/base",
		},
	}

	app, err := Load(prog)
	assert.Nil(t, err)
	assert.Equal(t, app.Entry, prog.Entry)
	assert.Equal(t, app.Imports, prog.Imports)
	start := app.Flow.AddNode("start")
	end := app.Flow.AddNode("end")
	assert.Len(t, start.Children, 1)
	assert.Equal(t, start.Children[""], end)
	assert.Panics(t, func() {
		app.Component("doesnotexist")
	})
}

func TestLoaderMissingImports(t *testing.T) {
	prog := parser.Program{
		Entry: "start",
		Flow:  map[parser.ID]parser.ID{"start": "end"},
		Components: map[parser.ID]string{
			"start": "base.StringCLI",
			"end":   "fmt.Println",
		},
		Imports: map[string]string{
			"fmt":  "fmt",
			"base": "github.com/Logiraptor/oak/core/base/doesnotexist",
		},
	}

	_, err := Load(prog)
	assert.NotNil(t, err)
}

func TestLoaderCardinalityMismatch(t *testing.T) {
	prog := parser.Program{
		Entry: "start",
		Flow:  map[parser.ID]parser.ID{"start": "end"},
		Components: map[parser.ID]string{
			"start": "func() (int, int) {return 0, 0}",
			"end":   "func(int) {}",
		},
	}

	_, err := Load(prog)
	assert.Len(t, err, 1)
	assert.Contains(t, err[0].Error(), "start")
	assert.Contains(t, err[0].Error(), "end")
	assert.Contains(t, err[0].Error(), "(int, int)")
	assert.Contains(t, err[0].Error(), "(int)")
}

func TestLoaderTypeMismatch(t *testing.T) {
	prog := parser.Program{
		Entry: "start",
		Flow:  map[parser.ID]parser.ID{"start": "end"},
		Components: map[parser.ID]string{
			"start": "func() (bool) {return true}",
			"end":   "func(int) {}",
		},
	}

	_, err := Load(prog)
	assert.Len(t, err, 1)
	assert.Contains(t, err[0].Error(), "start")
	assert.Contains(t, err[0].Error(), "end")
	assert.Contains(t, err[0].Error(), "int")
	assert.Contains(t, err[0].Error(), "bool")
}

func TestLoaderVariadic(t *testing.T) {
	prog := parser.Program{
		Entry: "start",
		Flow:  map[parser.ID]parser.ID{"start": "end"},
		Components: map[parser.ID]string{
			"start": "func() (bool, bool) {return true, false}",
			"end":   "func(...bool) {}",
		},
	}

	_, err := Load(prog)
	assert.Len(t, err, 0)
}

func TestLoaderVariadicMismatch(t *testing.T) {
	prog := parser.Program{
		Entry: "start",
		Flow:  map[parser.ID]parser.ID{"start": "end"},
		Components: map[parser.ID]string{
			"start": `func() (string, bool, int) {return "", true, 0}`,
			"end":   "func(string, ...bool) {}",
		},
	}

	_, err := Load(prog)
	assert.Len(t, err, 1)
	assert.Contains(t, err[0].Error(), "start")
	assert.Contains(t, err[0].Error(), "end")
	assert.Contains(t, err[0].Error(), "int")
	assert.Contains(t, err[0].Error(), "bool")
}

func TestLoaderInvalidSyntax(t *testing.T) {
	prog := parser.Program{
		Entry: "start",
		Flow:  map[parser.ID]parser.ID{"start": "end"},
		Components: map[parser.ID]string{
			"start": "<invalid syntax>",
			"end":   "func(...bool) {}",
		},
	}

	_, err := Load(prog)
	assert.Len(t, err, 1)
}

func ExampleLoad() {
	prog, _ := parser.ParseReader(strings.NewReader(`
entry: start
components:
  start: func(s string) {}
`))
	app, _ := Load(prog)
	fmt.Println(app.Components[0].Inputs)
	// Output: (s string)
}
