package main

import "testing"

func TestParseText(t *testing.T) {
	text := []byte("Hello world\nThis is a test\nAnother line")
	expected := lineData{2, 8, 39, 39, 14, "test.txt", ""}
	result := parseText(text, "test.txt")

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
	// Reset cmdFlags to default
	cmdFlags = flags{}
	if !allFlagsFalse() {
		t.Errorf("Expected allFlagsFalse to be true when no flags are set")
	}

	cmdFlags.printBytes = true
	if allFlagsFalse() {
		t.Errorf("Expected allFlagsFalse to be false when printBytes is set")
	}
}

func TestParseFile(t *testing.T) {
	fileName := "non_existent_file.txt"
	expected := lineData{0, 0, 0, 0, 0, fileName, "No such file or directory"}
	result := readFile(fileName)

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
