package data

import (
	"fmt"
	"log"

	apyhub "github.com/apyhub/apyhub.go"
)

func init() {
	apyhub.InitApyHub("Enter the apyhub token")
}

func Currencylist() {

	currency, err := apyhub.CurrencyList()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("currencylist:%+v", currency)
}
