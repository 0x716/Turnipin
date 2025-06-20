// extract.go
//
// Author: lamkuan (https://github.com/0x716)
// Created: 2025-06-18
// License: MIT
//
// Description:
// This file provides utility functions to extract plain text content
// from PDF and Word (.docx) files. It supports parsing text for further
// use in plagiarism detection and similarity analysis.
//
// Dependencies:
// - github.com/ledongthuc/pdf (for PDF extraction)
// - github.com/unidoc/unioffice/document (for DOCX extraction)

package util

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/ledongthuc/pdf"
	"github.com/unidoc/unioffice/document"
)

// ExtractText determines the file type by extension and routes it to
// the appropriate text extraction function.
// Supported formats: .pdf, .docx
func ExtractText(filePath string) (string, error) {
	if strings.HasSuffix(filePath, ".pdf") {
		return extractFromPDF(filePath)
	} else if strings.HasSuffix(filePath, ".docx") {
		return extractFromWord(filePath)
	}
	return "", errors.New("unsupported format")
}

// extractFromPDF extracts text content from a PDF file using the ledongthuc/pdf package.
// It loops through all pages and appends plain text into a buffer.
func extractFromPDF(filePath string) (string, error) {
	f, r, err := pdf.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to open PDF: %w", err)
	}
	defer f.Close()

	var buf bytes.Buffer

	totalPage := r.NumPage()
	for pageIndex := 1; pageIndex <= totalPage; pageIndex++ {
		page := r.Page(pageIndex)
		if page.V.IsNull() {
			continue
		}
		content, err := page.GetPlainText(nil)
		if err != nil {
			return "", fmt.Errorf("failed to extract page %d: %w", pageIndex, err)
		}
		io.WriteString(&buf, content)
	}

	return buf.String(), nil
}

// extractFromWord extracts plain text from a .docx file using the unioffice package.
// It iterates through paragraphs and text runs to reconstruct the document content.
func extractFromWord(filePath string) (string, error) {
	doc, err := document.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to open Word file: %w", err)
	}

	var buf bytes.Buffer
	for _, para := range doc.Paragraphs() {
		for _, run := range para.Runs() {
			buf.WriteString(run.Text())
		}
		buf.WriteString("\n")
	}

	return buf.String(), nil
}
