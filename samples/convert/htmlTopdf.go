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

func HtmlTopdf() {
	File, err := os.Open("../test_file/sample.html")
	if err != nil {
		log.Fatal(err)
	}

	// html to pdf file to file
	byt, err := apyhub.HtmlToPdfAsFile(File, true)
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile("test.pdf", byt, 0777)
	if err != nil {
		log.Fatal(err)
	}

	// html to pdf fileurl to file
	fileurl := "https://assets.apyhub.com/samples/sample.html"

	byt, err = apyhub.HtmlToPdfAsFile(fileurl, false)
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile("test.pdf", byt, 0777)
	if err != nil {
		log.Fatal(err)
	}

	// html to pdf  file to url

	url, err := apyhub.HtmlToPdfAsUrl(File, false)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(url)

	// html to pdf  fileurl to file
	fileurl = "https://assets.apyhub.com/samples/sample.html"

	url, err = apyhub.HtmlToPdfAsUrl(fileurl, false)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(url)
}
