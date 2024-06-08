package main

import (
	"bufio"
	"fmt"
	"io"
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
	debug         bool
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
	const version = "1.0.4"
	parseArgs()

	if cmdFlags.debug {

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
	if len(fileNames) == 0 {
		lines = append(lines, parseUserInput())
	} else {
		for _, filename := range fileNames {
			if filename == "-" {
				lines = append(lines, parseUserInput())
				continue
			}
			lines = append(lines, parseFile(filename))
		}
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
	processingFlags := true
	nullSeparatesFileNames := false

	for _, arg := range args {

		if strings.HasPrefix(arg, "--files0-from=") {
			fmt.Println("Not Implemented due to Windows limitations")
			os.Exit(1)
		}

		if !processingFlags && !nullSeparatesFileNames {
			fileNames = append(fileNames, arg)
			continue
		}
		if arg == "--" {
			processingFlags = false
			continue
		}

		if strings.HasPrefix(arg, "--") {
			// Handle long options
			switch arg {
			case "--byte":
				cmdFlags.printBytes = true
			case "--chars":
				cmdFlags.printChars = true
			case "--lines":
				cmdFlags.printLines = true
			case "--max-line-length":
				cmdFlags.maxLineLength = true
			case "--words":
				cmdFlags.printWords = true
			case "--help":
				cmdFlags.helpFlag = true
			case "--version":
				cmdFlags.versionFlag = true
			case "--about":
				cmdFlags.aboutFlag = true
			default:
				fmt.Printf("wc: unrecognized option %s\n", arg)
				fmt.Println("Try 'wc --help' for more information.")
				os.Exit(1)
			}

			continue
		}

		if strings.HasPrefix(arg, "-") && arg != "-" {
			for _, flag := range arg {
				switch flag {
				case '-':
					continue
				case 'c':
					cmdFlags.printBytes = true
				case 'm':
					cmdFlags.printChars = true
				case 'l':
					cmdFlags.printLines = true
				case 'L':
					cmdFlags.maxLineLength = true
				case 'w':
					cmdFlags.printWords = true
				case 'd':
					cmdFlags.debug = true
				default:
					fmt.Printf("wc: invalid option -- %s\n", string(flag))
					fmt.Println("Try 'wc --help' for more information.")
					os.Exit(1)
				}
			}
			continue
		}

		if !nullSeparatesFileNames {
			fileNames = append(fileNames, arg)
		}

	}

}

func parseFile(fileName string) lineData {
	fileData, err := os.ReadFile(fileName)
	if err != nil {
		return lineData{0, 0, 0, 0, 0, fileName, "No such file or directory"}
	}

	return parseText(fileData, fileName)
}

func parseUserInput() lineData {
	reader := bufio.NewReader(os.Stdin)
	userInput, err := reader.ReadBytes('\x00')
	if err != nil {
		if err == io.EOF {
			return parseText(userInput, "-")
		}
		return lineData{0, 0, 0, 0, 0, "-", "Error Reading Standard Input"}
	}
	return parseText(userInput, "-")
}

func parseText(text []byte, lineName string) lineData {
	charCount := 0
	lineCount := 0
	maxLineLen := 0
	wordCount := 0
	byteCount := len(text)

	stringData := string(text)
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

	totalBytes += byteCount
	totalChars += charCount
	totalLines += lineCount
	totalWords += wordCount
	totalMaxLineLen += maxLineLen

	return lineData{lineCount, wordCount, charCount, byteCount, maxLineLen, lineName, ""}
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

		fmt.Printf("%s \n", line.name)
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

		if maxLineCountWidth < newLineWidth {
			maxLineCountWidth = newLineWidth
		}
		if maxWordCountWidth < lineWordWidth {
			maxWordCountWidth = lineWordWidth
		}
		if maxCharCountWidth < lineCharWidth {
			maxCharCountWidth = lineCharWidth
		}
		if maxByteCountWidth < lineByteWidth {
			maxByteCountWidth = lineByteWidth
		}
		if maxLineLengthWidth < lineLenWidth {
			maxLineLengthWidth = lineLenWidth
		}
		if maxLineNameWidth < lineNameWidth {
			maxLineNameWidth = lineNameWidth
		}

	}
	lineSize := lineSize{maxLineCountWidth, maxWordCountWidth, maxCharCountWidth, maxByteCountWidth, maxLineLengthWidth}
	return lineSize
}

func printHelp() {
	help := `Usage: wc [OPTION]... [FILE]...

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
-d                     Enables Debugging

--files0-from=F        Has not been included in the Windows version due to issues with null terminators in Windows. 

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
