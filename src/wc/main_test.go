package main

import (
	"strings"
	"testing"
)

func TestParseText(t *testing.T) {
	text := []byte("Hello world\nThis is a test\nAnother line\n")
	expected := lineData{3, 8, 40, 40, 15, "test.txt", ""}

	var result lineData
	for _, line := range strings.Split(string(text), "\n") {
		if line != "" {
			parseText(&result, []byte(line+"\n"))
		}
	}

	if result.lineCount != expected.lineCount {
		t.Errorf("Expected lineCount %d, got %d", expected.lineCount, result.lineCount)
	}
	if result.wordCount != expected.wordCount {
		t.Errorf("Expected wordCount %d, got %d", expected.wordCount, result.wordCount)
	}
	if result.charCount != expected.charCount {
		t.Errorf("Expected charCount %d, got %d", expected.charCount, result.charCount)
	}
	if result.byteCount != expected.byteCount {
		t.Errorf("Expected byteCount %d, got %d", expected.byteCount, result.byteCount)
	}
	if result.maxLineLen != expected.maxLineLen {
		t.Errorf("Expected maxLineLen %d, got %d", expected.maxLineLen, result.maxLineLen)
	}
}

func TestAllFlagsFalse(t *testing.T) {
	// Reset cmdFlags to default values
	cmdFlags = flags{}

	if !allFlagsFalse() {
		t.Errorf("Expected allFlagsFalse to be true, got false")
	}
}
