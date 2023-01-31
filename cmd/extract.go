package cmd

import "fmt"

// Extract keyLists example: [metadata, name]
func Extract(keyLists []string, value string, yamls []map[string]interface{}) [][]byte {
	var extracted [][]byte
	for _, y := range yamls {
		t := y
		for i, k := range keyLists {
			if i < (len(keyLists) - 1) {
				w := t[k]
				t = w.(map[string]interface{})
				continue
			} else {
				if t[k] == value {
					s, err := ParseToString(y)
					if err != nil {
						fmt.Errorf("%s", err)
						return nil
					}
					extracted = append(extracted, s)
					break
				}
			}
		}
	}
	return extracted
}
