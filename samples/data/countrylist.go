package data

import (
	"fmt"
	"log"

	apyhub "github.com/apyhub/apyhub.go"
)

func init() {
	apyhub.InitApyHub("Enter the apyhub token")
}

func Countrylist() {

	countrylist, err := apyhub.CountryList()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("country:%+v", countrylist)
}
