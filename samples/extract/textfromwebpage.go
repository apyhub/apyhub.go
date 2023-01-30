package extract

import (
	"fmt"
	"log"

	apyhub "github.com/apyhub/apyhub.go"
)

func init() {
	apyhub.InitApyHub("Enter Apyhub token")
}

func Textfromwebpage() {

	// To extract the data from pdf file
	data, err := apyhub.TextFromWebpage("https://apyhub.com/benefits")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("textfrompdf:", data)
}
