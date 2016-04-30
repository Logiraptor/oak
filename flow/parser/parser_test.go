package parser

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/stretchr/testify/assert"

	"testing"
)

type errReader struct{}

func (e errReader) Read(b []byte) (int, error) {
	return 0, errors.New("generic error")
}

func TestParseErrReader(t *testing.T) {
	_, err := ParseReader(errReader{})
	assert.NotNil(t, err)
}

func TestParseReader(t *testing.T) {
	type testCase struct {
		input  string
		output Program
		err    bool
	}

	var tcs = []testCase{
		{
			input:  ``,
			output: Program{},
		},
		{
			input: `
	not yaml
`,
			err: true,
		},
		{
			input: `
entry: entry name
`,
			output: Program{
				Entry: "entry name",
			},
		},
	}

	for _, tc := range tcs {
		prog, err := ParseReader(strings.NewReader(tc.input))
		if tc.err {
			assert.NotNil(t, err)
			continue
		}
		assert.Nil(t, err)
		assert.EqualValues(t, tc.output, prog)
	}
}

func TestParseYAMLFile(t *testing.T) {
	f, err := ioutil.TempFile("", "flow-test")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Fprintln(f, `
entry: start
`)
	f.Close()
	prog, err := ParseFile(f.Name())
	assert.Nil(t, err)
	assert.Equal(t, Program{
		Entry: "start",
	}, prog)
}

func TestParseYAMLMissingFile(t *testing.T) {
	f, err := ioutil.TempFile("", "flow-test")
	if err != nil {
		t.Fatal(err)
	}
	f.Close()
	os.Remove(f.Name())
	_, err = ParseFile(f.Name())
	assert.NotNil(t, err)
}
