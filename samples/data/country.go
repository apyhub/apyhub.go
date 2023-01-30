package data

import (
	"fmt"
	"log"

	apyhub "github.com/apyhub/apyhub.go"
)

func init() {
	apyhub.InitApyHub("Enter the apyhub token")
}

func Country() {

	country, err := apyhub.Country("IN")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("country:%+v", country)
}
