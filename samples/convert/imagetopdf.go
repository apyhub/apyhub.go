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

func Imagetopdf() {
	File, err := os.Open("../test_file/sample.jpg")
	if err != nil {
		log.Fatal(err)
	}

	//image to pdf file to file
	byt, err := apyhub.ImageToPdfAsFile(File, true)
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile("test.pdf", byt, 0777)
	if err != nil {
		log.Fatal(err)
	}

	//image to pdf fileurl to file
	fileurl := "https://assets.apyhub.com/samples/sample.jpg"

	byt, err = apyhub.ImageToPdfAsFile(fileurl, false)
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile("test.pdf", byt, 0777)
	if err != nil {
		log.Fatal(err)
	}

	// image to pdf  file to url

	url, err := apyhub.ImageToPdfAsUrl(File, false)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(url)

	// image to pdf  fileurl to file
	fileurl = "https://assets.apyhub.com/samples/sample.html"

	url, err = apyhub.ImageToPdfAsUrl(fileurl, false)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(url)
}
