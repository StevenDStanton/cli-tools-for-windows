package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

const (
	VERSION = "v0.0.3"
)

var (
	developerFlag = flag.Bool("d", false, "In Development Flag to block accidental use")
	helpFlag      = flag.Bool("help", false, "Display Help")
	versionFlag   = flag.Bool("version", false, "Display Version Information")
)

func init() {
	flag.Parse()
}

func main() {

	if *helpFlag {
		printHelp()
		os.Exit(0)
	}
	if *versionFlag {
		printVersion()
		os.Exit(0)
	}
	if !*developerFlag {
		log.Fatalln("This program is in development and not ready for usage")
	}

	log.Println("Ready to use")
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

func printVersion() {
	fmt.Printf("go Version %s\nCopyright 2024 The Simple Dev\nLicense MIT - No Warranty\n\nWritten By Steven Stanton", VERSION)
	os.Exit(0)
}
