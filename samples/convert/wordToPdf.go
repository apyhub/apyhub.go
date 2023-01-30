package convert

import (
	"fmt"
	"log"
	"os"

	apyhub "github.com/apyhub/apyhub.go"
)

func init() {
	apyhub.InitApyHub("APT0rzEvAp76Izffo3ek7Zg7FpkK6csxzJ5puXCwWAw3lg0cFtj")
}

func WordToPdf() {
	File, err := os.Open("../test_file/sample.docx")
	if err != nil {
		log.Fatal(err)
	}

	// word to excel file to file
	byt, err := apyhub.WordToPdfAsFile(File, true)
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile("test.pdf", byt, 0777)
	if err != nil {
		log.Fatal(err)
	}

	// csv to excel fileurl to file
	fileurl := "https://assets.apyhub.com/samples/sample.docx"

	byt, err = apyhub.WordToPdfAsFile(fileurl, true)
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile("test.pdf", byt, 0777)
	if err != nil {
		log.Fatal(err)
	}

	// csv to excel  file to url

	url, err := apyhub.WordToPdfAsUrl(File, true)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(url)

	// csv to excel  fileurl to file
	fileurl = "https://assets.apyhub.com/samples/sample.docx"

	url, err = apyhub.WordToPdfAsUrl(fileurl, true)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(url)
}
