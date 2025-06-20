// levenshtein.go
//
// Author: lamkuan (https://github.com/0x716)
// Created: 2025-06-18
// License: MIT
//
// Description:
// Calculates the Levenshtein edit distance between two strings.

package util

// LevenshteinDistance calculates the minimum edit distance between a and b.
func LevenshteinDistance(a, b string) int {
	arunes := []rune(a)
	brunes := []rune(b)

	lenA := len(arunes)
	lenB := len(brunes)

	// Create a 2D slice for dynamic programming
	dp := make([][]int, lenA+1)
	for i := 0; i <= lenA; i++ {
		dp[i] = make([]int, lenB+1)
	}

	// Initialize base cases
	for i := 0; i <= lenA; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= lenB; j++ {
		dp[0][j] = j
	}

	// Compute distance
	for i := 1; i <= lenA; i++ {
		for j := 1; j <= lenB; j++ {
			cost := 0
			if arunes[i-1] != brunes[j-1] {
				cost = 1
			}

			dp[i][j] = min(
				dp[i-1][j]+1,      // deletion
				dp[i][j-1]+1,      // insertion
				dp[i-1][j-1]+cost, // substitution
			)
		}
	}

	return dp[lenA][lenB]
}

func min(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
		return c
	}
	if b < c {
		return b
	}
	return c
}
