package convert

import (
	"fmt"
	"log"
	"os"

	apyhub "github.com/apyhub/apyhub.go"
)

func init() {
	apyhub.InitApyHub("Enter the apyhub token")
}

func SpreadsheetToPdf() {
	File, err := os.Open("../test_file/sample.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	// csv to pdf file to file
	byt, err := apyhub.SpreadsheetToPdfAsFile(File, true)
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile("test.pdf", byt, 0777)
	if err != nil {
		log.Fatal(err)
	}

	// csv to pdf fileurl to file
	fileurl := "https://assets.apyhub.com/samples/sample.xlsx"

	byt, err = apyhub.SpreadsheetToPdfAsFile(fileurl, true)
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile("test.pdf", byt, 0777)
	if err != nil {
		log.Fatal(err)
	}

	// csv to pdf  file to url

	url, err := apyhub.SpreadsheetToPdfAsUrl(File, true)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(url)

	// csv to pdf  fileurl to file
	fileurl = "https://assets.apyhub.com/samples/sample.xlsx"

	url, err = apyhub.SpreadsheetToPdfAsUrl(fileurl, true)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(url)
}
