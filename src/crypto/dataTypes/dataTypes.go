package dataTypes

import (
	"fmt"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type Rate struct {
	Currency string
	Price    float64
}

func (r *Rate) String() string {
	p := message.NewPrinter(language.English)
	return fmt.Sprintf("%s: $%s", r.Currency, p.Sprintf("%f", r.Price))
}
