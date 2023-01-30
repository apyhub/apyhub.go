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

func PresentationToPdf() {
	File, err := os.Open("../test_file/sample.pptx")
	if err != nil {
		log.Fatal("error", err)
	}

	// ppt to pdf file to file
	byt, err := apyhub.PresentationToPdfAsFile(File, true)
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile("test.pdf", byt, 0777)
	if err != nil {
		log.Fatal(err)
	}

	// ppt to pdf fileurl to file
	fileurl := "https://assets.apyhub.com/samples/sample.md"

	byt, err = apyhub.PresentationToPdfAsFile(fileurl, true)
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile("test.html", byt, 0777)
	if err != nil {
		log.Fatal(err)
	}

	// ppt to pdf  file to url

	url, err := apyhub.PresentationToPdfAsUrl(File, true)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(url)

	// ppt to pdf  fileurl to file
	fileurl = "https://assets.apyhub.com/samples/sample.md"

	url, err = apyhub.PresentationToPdfAsUrl(fileurl, true)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(url)
}
