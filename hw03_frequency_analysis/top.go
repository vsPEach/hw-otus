package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

func Top10(text string) []string {
	top := make(map[string]int)
	fields := strings.Fields(text)
	for _, key := range fields {
		top[key]++
	}
	keys := make([]string, 0, len(top))
	for key := range top {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		if top[keys[i]] != top[keys[j]] {
			return top[keys[i]] > top[keys[j]]
		}
		return keys[i] < keys[j]
	})

	if len(keys) <= 10 {
		return keys
	}
	return keys[:10]
}
