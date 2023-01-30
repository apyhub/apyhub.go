package extract

import (
	"fmt"
	"log"
	"os"

	apyhub "github.com/apyhub/apyhub.go"
)

func init() {
	apyhub.InitApyHub("Enter your token")
}

func Textfrompdf() {

	File, err := os.Open("../test_file/sample.pdf")
	if err != nil {
		log.Fatal(err)
	}

	// To extract the data from pdf file
	data, err := apyhub.TextFromPdf(File)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("textfrompdf:", data)

	// To extract the data from pdf url
	fileurl := "https://assets.apyhub.com/samples/sample.pdf"

	data, err = apyhub.TextFromPdf(fileurl)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("image meta data:%+v", data)
}
