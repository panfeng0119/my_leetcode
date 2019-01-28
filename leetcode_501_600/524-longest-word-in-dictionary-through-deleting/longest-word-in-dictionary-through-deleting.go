package problem524

import (
	"sort"
)
func findLongestWord(s string, d []string) string {
	sort.Slice(d, func(i, j int) bool {
		return d[i] < d[j]
	})
	res := ""
	for _, word := range d {
		if canBeFormed(s, word) && len(word) > len(res) {
			res = word
		}
	}
	return res
}

func canBeFormed(s, target string) bool {
	i, j := 0, 0
	for i < len(target) {
		for j < len(s) && s[j] != target[i] {
			j++
		}
		if j == len(s) {
			return false
		}
		i++
		j++
	}
	return true
}