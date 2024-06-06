package main

import (
	"fmt"
	"os"
	"strings"
)

type lineData struct {
	lineCount  int
	wordCount  int
	charCount  int
	byteCount  int
	maxLineLen int
	name       string
	err        string
}

type lineSize struct {
	lineCount  int
	wordCount  int
	charCount  int
	byteCount  int
	maxLineLen int
	name       int
}

type flags struct {
	printBytes    bool
	printChars    bool
	printLines    bool
	maxLineLength bool
	printWords    bool
	helpFlag      bool
	versionFlag   bool
	aboutFlag     bool
}

var (
	fileNames       []string
	totalBytes      = 0
	totalChars      = 0
	totalLines      = 0
	totalWords      = 0
	totalMaxLineLen = 0
	cmdFlags        flags
)

func main() {
	const version = "0.8.0"
	parseArgs()

	debug := true

	if debug {

		fmt.Printf(`Debugging
printBytes: %t
printChars: %t
printLines: %t
maxLineLength: %t
printWords: %t
helpFlag: %t
versionFlag: %t
aboutFlag: %t
IsAllFalse: %t

`, cmdFlags.printBytes, cmdFlags.printChars, cmdFlags.printLines, cmdFlags.maxLineLength, cmdFlags.printWords, cmdFlags.helpFlag, cmdFlags.versionFlag, cmdFlags.aboutFlag, allFlagsFalse())
	}

	if cmdFlags.versionFlag {
		printVersion(version)
	}

	if cmdFlags.helpFlag {
		printHelp()
	}

	if cmdFlags.aboutFlag {
		printAbout()
	}

	lines := []lineData{}
	for _, filename := range fileNames {
		lines = append(lines, parseFile(filename))
	}

	if len(fileNames) > 1 {
		lines = append(lines, lineData{totalLines, totalWords, totalChars, totalBytes, totalMaxLineLen, "total", ""})
	}

	lineSize := getMaxWidths(lines)

	printLinesToConsole(lines, lineSize)

}

func allFlagsFalse() bool {
	return !cmdFlags.printBytes &&
		!cmdFlags.printChars &&
		!cmdFlags.printLines &&
		!cmdFlags.maxLineLength &&
		!cmdFlags.printWords
}

func parseArgs() {
	//I am aware of the flags package. However as I am trying to replicate how wc works on linux it proved to limited for my needs.
	args := os.Args[1:]

	for _, arg := range args {
		switch arg {
		case "--byte", "-c":
			cmdFlags.printBytes = true
		case "--chars", "-m":
			cmdFlags.printChars = true
		case "--lines", "-l":
			cmdFlags.printLines = true
		case "--max-line-length", "-L":
			cmdFlags.maxLineLength = true
		case "--words", "-w":
			cmdFlags.printWords = true
		case "--help":
			cmdFlags.helpFlag = true
		case "--version":
			cmdFlags.versionFlag = true
		case "--about":
			cmdFlags.aboutFlag = true
		default:
			if strings.HasPrefix(arg, "-") {
				continue
			}
			fileNames = append(fileNames, arg)
		}
	}

}

func parseFile(fileName string) lineData {
	fileData, err := os.ReadFile(fileName)
	if err != nil {
		return lineData{0, 0, 0, 0, 0, fileName, "No such file or directory"}
	}
	charCount := 0
	lineCount := 0
	maxLineLen := 0
	wordCount := 0
	byteCount := len(fileData)

	stringData := string(fileData)
	inWord := false
	currentLineLen := 0
	for _, char := range stringData {
		charCount++

		if char == '\n' {
			lineCount++
			if currentLineLen > maxLineLen {
				maxLineLen = currentLineLen
			}
			currentLineLen = 0
		} else {
			currentLineLen++
		}

		if inWord && (char == ' ' || char == '\n' || char == '\r' || char == '\t') {
			inWord = false
		}

		if !inWord && (char != ' ' && char != '\n' && char != '\r' && char != '\t') {
			inWord = true
			wordCount++
		}
	}

	totalBytes += len(fileData)
	totalChars += charCount
	totalLines += lineCount
	totalWords += wordCount
	totalMaxLineLen += maxLineLen

	return lineData{lineCount, wordCount, charCount, byteCount, maxLineLen, fileName, ""}

}

func printLinesToConsole(lineData []lineData, lineSize lineSize) {

	for _, line := range lineData {
		if line.err != "" {
			fmt.Printf("wc: %s: %s\n", line.name, line.err)
			continue
		}
		if cmdFlags.printLines || allFlagsFalse() {
			fmt.Printf("%*d ", lineSize.lineCount, line.lineCount)
		}

		if cmdFlags.printWords || allFlagsFalse() {
			fmt.Printf("%*d ", lineSize.wordCount, line.wordCount)
		}

		if cmdFlags.printChars {
			fmt.Printf("%*d ", lineSize.charCount, line.charCount)
		}

		if cmdFlags.printBytes || allFlagsFalse() {
			fmt.Printf("%*d ", lineSize.byteCount, line.byteCount)
		}

		if cmdFlags.maxLineLength {
			fmt.Printf("%*d ", lineSize.maxLineLen, line.maxLineLen)
		}

		fmt.Printf("%*s \n", lineSize.name, line.name)
	}
}

func getMaxWidths(lines []lineData) lineSize {
	maxLineCountWidth := 0
	maxWordCountWidth := 0
	maxCharCountWidth := 0
	maxByteCountWidth := 0
	maxLineLengthWidth := 0
	maxLineNameWidth := 0

	for _, line := range lines {

		newLineWidth := len(fmt.Sprintf("%d", line.lineCount))
		lineWordWidth := len(fmt.Sprintf("%d", line.wordCount))
		lineCharWidth := len(fmt.Sprintf("%d", line.charCount))
		lineByteWidth := len(fmt.Sprintf("%d", line.byteCount))
		lineLenWidth := len(fmt.Sprintf("%d", line.maxLineLen))
		lineNameWidth := len(line.name)

		if maxLineCountWidth > newLineWidth {
			maxLineCountWidth = newLineWidth
		}
		if maxWordCountWidth > lineWordWidth {
			maxWordCountWidth = lineWordWidth
		}
		if maxCharCountWidth > lineCharWidth {
			maxCharCountWidth = lineCharWidth
		}
		if maxByteCountWidth > lineByteWidth {
			maxByteCountWidth = lineByteWidth
		}
		if maxLineLengthWidth > lineLenWidth {
			maxLineLengthWidth = lineLenWidth
		}
		if maxLineNameWidth > lineNameWidth {
			maxLineNameWidth = lineNameWidth
		}

	}
	lineSize := lineSize{maxLineCountWidth, maxWordCountWidth, maxCharCountWidth, maxByteCountWidth, maxLineLengthWidth, maxLineNameWidth}
	return lineSize
}

func printHelp() {
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
	os.Exit(0)
}

func printVersion(version string) {
	fmt.Printf("go Version %s\nCopyright 2024 The Simple Dev\nLicense MIT - No Warranty\n\nWritten By Steven Stanton\nReverse Engineered by RTFM", version)
	os.Exit(0)
}

func printAbout() {
	about := `This is a simple implementation of the wc command in Go.

This program has been reversed engineered from the GNU Coreutils wc program using only documentation and observed behavior in a clean room environment.

Author:         Steven Stanton
License:        MIT - No Warranty
Author Github:  https//github.com/StevenDStanton
Project Github: https://github.com/StevemStanton/ltfw

Part of my Linux Tools for Windows (ltfw) project.
`
	fmt.Println(about)
	os.Exit(0)
}
