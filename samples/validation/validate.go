package validate

import (
	"fmt"
	"log"

	"github.com/apyhub/apyhub.go"
	"github.com/apyhub/apyhub.go/validation"
)

func init() {
	apyhub.InitApyHub("Enter your apyhub token")
}

func Validate() {

	//Domail validation
	DNSverify, err := apyhub.EmailValidation("hello@apyhub.com", validation.DNS)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(DNSverify)

	//Acadamic validation
	Acadamicverify, err := apyhub.EmailValidation("hello@harvard.edu", validation.Acadamic)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(Acadamicverify)

	//IN postcode validation
	Inpost, err := apyhub.PostCodeValidation("520010", validation.IN)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(Inpost)

	//UK postcode validation
	Ukpost, err := apyhub.PostCodeValidation("AB11 5QN", validation.UK)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(Ukpost)

	//Vat validation
	vatValid, err := apyhub.VatValidation("NL862735944B01")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(vatValid)
}
