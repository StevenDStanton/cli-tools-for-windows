package main

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/StevenDStanton/ltfw/crypto/api"
)

func fetchRate(pair string, wg *sync.WaitGroup, ch chan<- string) {
	defer wg.Done()
	response, err := api.GetRate(pair)
	if err != nil {
		ch <- fmt.Sprintf("Error fetching rate for %s: %v", pair, err)
		return
	}
	ch <- response.String() // Convert Rate to string using the String method
}

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		log.Fatalln("Must specify at least one pair such as BTC/USD")
	}

	var wg sync.WaitGroup
	ch := make(chan string, len(args))

	for _, pair := range args {
		wg.Add(1)
		go fetchRate(pair, &wg, ch)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	// Read from the channel and print responses
	for response := range ch {
		fmt.Println(response)
	}
}
