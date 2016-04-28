package parser

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
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

// ParseYAMLFile reads a flow Program from the specified file
func ParseYAMLFile(filename string) (Program, error) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return Program{}, err
	}
	var program Program
	err = yaml.Unmarshal(buf, &program)
	if err != nil {
		return Program{}, err
	}
	return program, nil
}
