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
	helpFlag    = flag.Bool("help", false, "display this help and exit")
	versionFlag = flag.Bool("version", false, "output version information and exit")
)

func main() {
	const version = "0.0.2"
	// charCount := 0
	// newLineCount := 0
	// maxLineLen := 0
	// wordCount := 0
	flag.Parse()

	if *versionFlag {
		fmt.Printf("go Version %s\nCopyright 2024 The Simple Dev\nLicense MIT - No Warranty\n\nWritten By Steven Stanton\nReverse Engineered by RTFM", version)
		os.Exit(0)
	}

	if *helpFlag {
		help := `Usage: wc [OPTION]... [FILE]...
Multiple Files: wc [OPTION]... --files0-from=F

Prints new line, word, and byte counts for each FILE and a total line if more file is specified.
A word is a non-zero length sequence of characters delimited by white space.

When no FILE or when FILE is -, read from standard input.

The options below may be used to select which counts are displayed. Always in the following order:
newLine, word, character, byte, maximum line length.

-c, --bytes            print the byte counts
-m, --chars            print the character counts
-l, --lines            print the newline counts
-L, --max-line-length  print the length of the longest line
-w, --words            print the word counts

--files0-from=F        read input from the files specified by NUL-terminated names in file F; If F is - then read names from standard input

--help                 display this help and exit
--version              output version information and exit
--about 			  display information about the program and exit`

		fmt.Println(help)
	}

}
