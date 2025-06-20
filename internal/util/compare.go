// compare.go
//
// Author: lamkuan (https://github.com/0x716)
// Created: 2025-06-18
// License: MIT
//
// Description:
// This utility module provides basic text similarity metrics including:
// - Jaccard similarity based on n-gram shingling
// - Levenshtein (edit distance) similarity

package util

import "math"

// CompareJaccard computes the Jaccard similarity between two token sequences,
// using n-gram shingling.
//
// Returns a float between 0.0 (no overlap) and 1.0 (identical).
func CompareJaccard(tokensA, tokensB []string, n int) float64 {
	ngramsA := GenerateNGrams(tokensA, n)
	ngramsB := GenerateNGrams(tokensB, n)

	setA := make(map[string]bool)
	setB := make(map[string]bool)

	for _, gram := range ngramsA {
		setA[gram] = true
	}
	for _, gram := range ngramsB {
		setB[gram] = true
	}

	intersection := 0
	union := make(map[string]bool)

	for gram := range setA {
		union[gram] = true
		if setB[gram] {
			intersection++
		}
	}
	for gram := range setB {
		union[gram] = true
	}

	if len(union) == 0 {
		return 0.0
	}

	return float64(intersection) / float64(len(union))
}

// CompareLevenshtein computes a normalized similarity between two strings
// based on Levenshtein edit distance.
//
// Returns 1.0 for identical strings, and approaches 0.0 as they diverge.
func CompareLevenshtein(a, b string) float64 {
	distance := LevenshteinDistance(a, b)
	maxLen := math.Max(float64(len([]rune(a))), float64(len([]rune(b))))

	if maxLen == 0 {
		return 1.0
	}

	return 1.0 - float64(distance)/maxLen
}
