package cmd

import (
	"bytes"
	"gopkg.in/yaml.v3"
)

func ParseToYaml(data []byte) []map[string]interface{} {
	decoder := yaml.NewDecoder(bytes.NewReader(data))
	var parsedYamls []map[string]interface{}
	for {
		var parsed map[string]interface{}
		if decoder.Decode(&parsed) != nil {
			break
		}
		parsedYamls = append(parsedYamls, parsed)
	}
	return parsedYamls
}

func ParseToString(data map[string]interface{}) ([]byte, error) {
	s, err := yaml.Marshal(data)
	if err != nil {
		return nil, err
	}
	return s, nil
}
