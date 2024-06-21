package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/StevenDStanton/cli-tools-for-windows/common"
)

type frequencykv struct {
	key   rune
	value int
}

const (
	tool    = "compress"
	version = "v0.1.0"
)

var (
	developerFlag = flag.Bool("d", false, "In Development Flag to block accidental use")
	helpFlag      = flag.Bool("help", false, "Display Help")
	versionFlag   = flag.Bool("version", false, "Display Version Information")
	encodeFlag    = flag.Bool("encode", false, "Encode File")
	decodeFlag    = flag.Bool("decode", false, "Decode file")
	inputFile     = flag.String("i", "", "Input file name")
	outputFile    = flag.String("o", "", "Output file name")
)

func init() {
	flag.Parse()

	if *helpFlag {
		printHelp()
		os.Exit(0)
	}
	if *versionFlag {
		versionInformation := common.PrintVersion(tool, version)
		fmt.Println(versionInformation)
		os.Exit(0)
	}
	if !*developerFlag {
		log.Fatalln("This program is in development and not ready for usage")
	}

}

func main() {
	menu()

}

func menu() {
	if *inputFile == "" || *outputFile == "" {
		log.Fatalln("Must have an input and output file name")
	}

	switch {
	case *encodeFlag:
		encodeFile()
	case *decodeFlag:
		decodeFile()
	default:
		log.Fatal("No option selected")
	}
}

func encodeFile() {
	frequencyMap := make(map[rune]int)
	file, err := os.Open(*inputFile)
	if err != nil {
		log.Fatalf("Unable to read file, %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(scanLinesWithNewlines)
	for scanner.Scan() {
		chunk := scanner.Text()
		buildFrequencyMap(frequencyMap, chunk)
	}

	sortedFrequencyMap := make([]frequencykv, 0, len(frequencyMap))

	for k, v := range frequencyMap {
		sortedFrequencyMap = append(sortedFrequencyMap, frequencykv{key: k, value: v})
	}

	sort.Slice(sortedFrequencyMap, func(i, j int) bool {
		return sortedFrequencyMap[i].value < sortedFrequencyMap[j].value
	})

	printMap(sortedFrequencyMap)
}

func decodeFile() {

}

func buildFrequencyMap(frequencyMap map[rune]int, chunk string) {
	for _, char := range chunk {
		frequencyMap[char]++
	}

}

func printMap(sortedFrequencyMap []frequencykv) {
	for _, kv := range sortedFrequencyMap {
		fmt.Printf("%c: %d\n", kv.key, kv.value)
	}
}

func scanLinesWithNewlines(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	if i := bytes.IndexByte(data, '\n'); i >= 0 {
		return i + 1, data[:i+1], nil
	}

	if atEOF {
		return len(data), data, nil
	}

	return 0, nil, nil
}

func printHelp() {
	help := `compress is a very basic compression tool just using Huffman
 Encoding and Decoding. 
 
 Not yet ready for use. Use at own risk

 -d developer flag allows execution - use at own risk

--help prints help

 `

	fmt.Println(help)
}
