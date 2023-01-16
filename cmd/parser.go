package cmd

import (
	"bytes"
	"fmt"
	"gopkg.in/yaml.v3"
)

func ParseYaml(data []byte) {
	decoder := yaml.NewDecoder(bytes.NewReader(data))
	var parsedYamls []map[string]interface{}
	for {
		var parsed map[string]interface{}
		if decoder.Decode(&parsed) != nil {
			break
		}
		parsedYamls = append(parsedYamls, parsed)
	}
	fmt.Printf("%s", parsedYamls)
}
