package flow

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type ComponentID string

type Config struct {
	Imports    map[string]string
	Components map[ComponentID]string
	Flow       map[ComponentID]ComponentID
	Entry      ComponentID
}

func ParseYAMLFile(filename string) (*Config, error) {
	var Config = new(Config)
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(buf, &Config)
	if err != nil {
		return nil, err
	}
	return Config, nil
}
