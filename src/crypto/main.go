package main

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/StevenDStanton/ltfw/common"
	"github.com/StevenDStanton/ltfw/crypto/api"
)

var (
	args []string
)

const (
	tool    = "Crypto"
	version = "v1.0.6"
)

func init() {
	args = os.Args[1:]
	if len(args) < 1 {
		log.Fatalln("Must specify at least one pair such as BTC/USD")
	}
	common.PrintVersion(tool, version)
}

func fetchRate(pair string) {
	response, err := api.GetRate(pair)
	if err != nil {
		fmt.Printf("Error fetching rate for %s: %v", pair, err)
		return
	}
	fmt.Println(response)
}

func main() {
	var wg sync.WaitGroup
	for _, pair := range args {
		wg.Add(1)
		go func() {
			fetchRate(pair)
			wg.Done()
		}()
	}
	wg.Wait()
}
