package common

import "sort"

func Map_string(m map[string]string) (result map[string]string, keys []string) {
	var keyst []string
	for key := range m {
		keyst = append(keyst, key)
	}
	sort.Strings(keyst)
	return m, keyst
}
