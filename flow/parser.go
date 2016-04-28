package flow

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Config is essentially an AST for flow programs
// Use Load to verify that a program type checks.
type Config struct {
	Imports    map[string]string
	Components map[ComponentID]string
	Flow       map[ComponentID]ComponentID
	Entry      ComponentID
}

// ParseYAMLFile reads a flow config from the specified file
func ParseYAMLFile(filename string) (Config, error) {
	var config Config
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return Config{}, err
	}
	err = yaml.Unmarshal(buf, &config)
	if err != nil {
		return Config{}, err
	}
	return config, nil
}
