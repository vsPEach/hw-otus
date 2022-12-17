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

	sort.Strings(keys)

	sort.SliceStable(keys, func(i, j int) bool {
		return top[keys[i]] > top[keys[j]]
	})
	if len(keys) == 0 {
		return nil
	}
	return keys[:10]
}
