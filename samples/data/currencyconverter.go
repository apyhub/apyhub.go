package data

import (
	"fmt"
	"log"

	apyhub "github.com/apyhub/apyhub.go"
	"github.com/apyhub/apyhub.go/helper"
)

func init() {
	apyhub.InitApyHub("Enter the apyhub token")
}

func Currencyconverter() {

	conv := helper.CurrencyConverter{
		Source: "usd",
		Target: "inr",
		Date:   "2023-01-27",
	}

	currency, err := apyhub.CurrencyConv(conv)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("currency:", currency)
}
