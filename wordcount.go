package main

import (
	"strings"
)

func countwords(text string) map[string]int {
	words := strings.Fields(strings.ToLower(text))
	result := map[string]int{}
	for _, wrd := range words {
		_, ok := result[wrd]
		if ok {
			result[wrd] += 1
		} else {
			result[wrd] = 1
		}
	}
	return result
}
