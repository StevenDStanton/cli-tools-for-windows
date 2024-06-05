package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	// printBytes         = flag.Bool("bytes", false, "print the byte counts")
	// printBytesShort    = flag.Bool("c", false, "print the byte counts")
	// printChars         = flag.Bool("chars", false, "print the character counts")
	// printCharsShort    = flag.Bool("m", false, "print the character counts")
	// printLines         = flag.Bool("lines", false, "print the newline counts")
	// printLinesShort    = flag.Bool("l", false, "print the newline counts")
	// maxLineLength      = flag.Bool("max-line-length", false, "print the length of the longest line")
	// maxLineLengthShort = flag.Bool("L", false, "print the length of the longest line")
	// printWords         = flag.Bool("words", false, "print the word counts")
	// printWordsShort    = flag.Bool("w", false, "print the word counts")
	// helpFlag    = flag.Bool("help", false, "display this help and exit")
	versionFlag = flag.Bool("version", false, "output version information and exit")
)

func main() {
	const version = "0.0.1"
	// charCount := 0
	// newLineCount := 0
	// maxLineLen := 0
	// wordCount := 0
	flag.Parse()

	if *versionFlag {
		fmt.Printf("wc Version %s\nCopyright 2024 The Simple Dev\nLicense MIT - No Warranty\n\nWritten By Steven Stanton\nReverse Engineered by RTFM", version)
		os.Exit(0)
	}

}
