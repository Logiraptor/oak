package parser

import (
	"io"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v1"
)

// ID is a unique identifier for a component in a flow program
type ID string

// Program is essentially an AST for flow programs
// Use Load to verify that a program type checks.
type Program struct {
	Imports    map[string]string
	Components map[ID]string
	Flow       map[ID]ID
	Entry      ID
}

// ParseFile reads a flow Program from the specified file
func ParseFile(filename string) (Program, error) {
	f, err := os.Open(filename)
	if err != nil {
		return Program{}, err
	}
	defer f.Close()
	return ParseReader(f)
}

// ParseReader parses a flow program from the given io.Reader
func ParseReader(r io.Reader) (Program, error) {
	buf, err := ioutil.ReadAll(r)
	if err != nil {
		return Program{}, err
	}
	var prog Program
	err = yaml.Unmarshal(buf, &prog)
	if err != nil {
		return Program{}, err
	}
	return prog, nil
}
