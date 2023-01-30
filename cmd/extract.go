package cmd

import "fmt"

func Extract(key string, value string, yamls []map[string]interface{}) [][]byte {
	var extracted [][]byte
	for _, y := range yamls {
		if y[key] == nil || y[key] != value {
			continue
		}
		s, err := ParseToString(y)
		if err != nil {
			fmt.Errorf("%s", err)
			return nil
		}
		extracted = append(extracted, s)
	}
	return extracted
}
