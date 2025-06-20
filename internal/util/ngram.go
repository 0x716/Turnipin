// shingle.go
//
// Author: lamkuan (https://github.com/0x716)
// Created: 2025-06-18
// License: MIT
//
// Description:
// This utility provides a function to generate n-gram (shingle) sequences
// from a list of tokens, for use in similarity comparison (e.g., Jaccard).

package util

import "strings"

// GenerateNGrams creates a list of n-grams from the input token slice.
// Each n-gram is a string of N concatenated tokens (joined with space).
//
// Example:
// tokens = ["this", "is", "a", "book"], n = 2
// → ["this is", "is a", "a book"]
func GenerateNGrams(tokens []string, n int) []string {
	var ngrams []string
	if len(tokens) < n {
		return ngrams
	}

	for i := 0; i <= len(tokens)-n; i++ {
		ngram := strings.Join(tokens[i:i+n], " ")
		ngrams = append(ngrams, ngram)
	}

	return ngrams
}
