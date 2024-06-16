package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
	"unicode/utf8"
)

type TTSRequest struct {
	Model  string `json:"model"`
	Input  string `json:"input"`
	Voice  string `json:"voice"`
	Format string `json:"response_format"`
	Speed  string `json:"speed"`
}

const (
	CONFIG_FILE   = "tts.config"
	CONFIG_DIR    = ".ltfw"
	defaultVoice  = "nova"
	defaultModel  = "tts-1-hd"
	defaultFormat = "mp3"
	defaultSpeed  = "1.0"
)

var (
	configFilePath string
	OPENAI_API_KEY string
	version        string
)

var (
	inputFile     = flag.String("f", "", "Input Markdown file")
	outputFile    = flag.String("o", "", "Output audio file")
	voiceOption   = flag.String("v", defaultVoice, "Voice Selection")
	modelOption   = flag.String("m", defaultModel, "Model Selection")
	formatOption  = flag.String("fmt", defaultFormat, "Select output format")
	speedOption   = flag.String("s", defaultSpeed, "Set audio speed")
	configureMode = flag.Bool("configure", false, "Enter Configuration Mode")
	helpFlag      = flag.Bool("help", false, "Displays Help Menu")
	versionFlag   = flag.Bool("version", false, "Displays version information")
)

func init() {
	version = "1.1.0"
	configure()
	flag.Parse()
}

func main() {

	switch {
	case *helpFlag:
		printHelp()
	case *configureMode:
		writeNewConfig()
	case *versionFlag:
		printVersion()
	default:
		if *inputFile == "" || *outputFile == "" {
			fmt.Println("Usage: tts -f filename.md -o filename.mp3")
			os.Exit(0)
		}

		inputContent := readFileData(*inputFile)

		ttsRequest := TTSRequest{
			Model:  *modelOption,
			Voice:  *voiceOption,
			Format: *formatOption,
			Input:  inputContent,
			Speed:  *speedOption,
		}

		tts(ttsRequest, *outputFile)
	}
}

func readFileData(inputFile string) string {
	inputContent, err := os.ReadFile(inputFile)
	checkFatalErrorExists("Error: reading input file", err)
	if utf8.RuneCount(inputContent) > 4096 {
		log.Fatalln("Input cannot exceed 4096 characters")
	}
	return string(inputContent)
}

func tts(ttsRequest TTSRequest, outputFile string) {
	requestBody, err := json.Marshal(ttsRequest)
	checkFatalErrorExists("Error: Unable to create request payload", err)

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/audio/speech", bytes.NewBuffer(requestBody))
	checkFatalErrorExists("Error: Unable to create HTTP request", err)

	req.Header.Set("Authorization", "Bearer "+OPENAI_API_KEY)
	req.Header.Set("Content-Type", "application/json")

	makeHttpRequest(req, outputFile)

}

func makeHttpRequest(req *http.Request, outputFile string) {
	client := &http.Client{Timeout: 90 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error sending request to OpenAI API: %v", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		responseBody, _ := io.ReadAll(resp.Body)
		log.Printf("OpenAI API request failed with status code: %d, response body: %s", resp.StatusCode, responseBody)
		return
	}

	outputFileData, err := os.Create(outputFile)
	if err != nil {
		log.Printf("Error creating output file: %v", err)
		return
	}
	defer outputFileData.Close()

	_, err = io.Copy(outputFileData, resp.Body)
	if err != nil {
		log.Printf("Error saving audio file: %v", err)
		return
	}

	fmt.Printf("Audio file saved successfully: %s\n", outputFile)
}

func printHelp() {
	help := `Usage: tts [OPTION]

	--configure          enter configuration prompt for API key
	--help               displays help
	--version            displays version information

	To use the program both of the below flags are require
	-o output audio file
	-f input text file

	Optional flags
	-v voice defaults to nova. 
		Voice options are: alloy, echo, fable, onyx, nova, and shimmer

	-m model defaults to tts-1-hd
		Model options are: tts-1 and tts-1-hd


	-fmt output format defaults to mp3
		Format options are: mp3, opus, aac, flac, wav, pcm
	
	-s speed defaults to 1
		Speed options 0.25 to 4.0
	`
	fmt.Println(help)
}

func printVersion() {
	fmt.Printf("go Version %s\nCopyright 2024 The Simple Dev\nLicense MIT - No Warranty\n\nWritten By Steven Stanton", version)
	os.Exit(0)
}

//This can be improved in the future to have a single config setup
//for all ltfw. However, to avoid over engineering the solution for
//now this single setup works. I will reveiew and refactor if it becomes an issue.
//For now each file gets a config for its usage

func configure() {
	home, err := os.UserHomeDir()
	checkFatalErrorExists("Unable to read user home directory", err)

	configDir := filepath.Join(home, CONFIG_DIR)

	err = os.MkdirAll(configDir, 0755)
	checkFatalErrorExists("Unable to create config directory", err)

	configFilePath = filepath.Join(configDir, CONFIG_FILE)

	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		writeNewConfig()
		return
	}
	checkFatalErrorExists("Unknown issue accessing config", err)

	readConfig()

}

func writeNewConfig() {
	fmt.Print("Please enter your OpenAI API Key: ")
	fmt.Scanln(&OPENAI_API_KEY)
	fileData := "OPENAI_API_KEY=" + OPENAI_API_KEY
	err := os.WriteFile(configFilePath, []byte(fileData), 0600)
	checkFatalErrorExists("", err)
	if err != nil {
		log.Fatalf("Unable to save config: %v\n", err)
	}
}

func readConfig() {
	file, err := os.Open(configFilePath)
	checkFatalErrorExists("unable to read config fil", err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			if key == "OPENAI_API_KEY" {
				OPENAI_API_KEY = value
			}

		}
	}

	if err := scanner.Err(); err != nil {
		checkFatalErrorExists("unable to read config file", err)
	}

	if OPENAI_API_KEY == "" {
		writeNewConfig()
	}
}

func checkFatalErrorExists(message string, err error) {
	if err != nil {
		log.Fatalf("%s: %v", message, err)
	}
}
