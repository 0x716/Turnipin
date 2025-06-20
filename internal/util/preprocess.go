// preprocess.go
//
// Author: lamkuan (https://github.com/0x716)
// Created: 2025-06-18
// License: MIT
//
// Description:
// This module provides text preprocessing utilities, including lowercasing,
// punctuation removal, tokenization, and optional stopword filtering.
// It prepares raw text for n-gram generation and similarity analysis.

package util

import (
	"regexp"
	"strings"
)

var (
	// Only keep letters, numbers, and spaces
	rePunct      = regexp.MustCompile(`[^\w\s]`)
	reMultiSpace = regexp.MustCompile(`\s+`)
)

// PreprocessText applies normalization, punctuation removal, and tokenization.
func PreprocessText(text string) []string {
	// Lowercase
	text = strings.ToLower(text)

	// Remove punctuation
	text = rePunct.ReplaceAllString(text, "")

	// Normalize whitespace
	text = reMultiSpace.ReplaceAllString(text, " ")

	// Trim leading/trailing space
	text = strings.TrimSpace(text)

	// Tokenize
	tokens := strings.Fields(text)

	return tokens
}

// Optional: Stopword filtering (basic list)
var stopwords = map[string]bool{
	"the": true, "is": true, "a": true, "an": true, "and": true,
	"or": true, "in": true, "on": true, "of": true, "to": true,
	"with": true, "for": true, "by": true, "as": true, "at": true,
}

// FilterStopwords removes common English stopwords.
func FilterStopwords(tokens []string) []string {
	var filtered []string
	for _, token := range tokens {
		if !stopwords[token] {
			filtered = append(filtered, token)
		}
	}
	return filtered
}

// PreprocessWithStopwords optionally removes stopwords after tokenizing
func PreprocessWithStopwords(text string) []string {
	tokens := PreprocessText(text)
	return FilterStopwords(tokens)
}
