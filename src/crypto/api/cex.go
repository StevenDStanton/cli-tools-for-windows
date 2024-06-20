package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/StevenDStanton/ltfw/crypto/dataTypes"
)

const apiURL = "https://cex.io/api/ticker/%s"

func GetRate(currency string) (*dataTypes.Rate, error) {
	parsedCurrency := strings.ToUpper(currency)
	response, err := http.Get(fmt.Sprintf(apiURL, parsedCurrency))
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}

	res, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var cryptoRate CEXResponse
	err = json.Unmarshal(res, &cryptoRate)
	if err != nil {
		return nil, err
	}

	last, err := strconv.ParseFloat(cryptoRate.Last, 64)
	if err != nil {
		return nil, err
	}

	rates := dataTypes.Rate{
		Currency: cryptoRate.Pair,
		Price:    last,
	}

	return &rates, nil

}
