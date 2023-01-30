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

func MarkdownToHtml() {
	File, err := os.Open("../test_file/sample.md")
	if err != nil {
		log.Fatal(err)
	}

	// markdown to html file to file
	byt, err := apyhub.MarkdownToHtmlAsFile(File)
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile("test.html", byt, 0777)
	if err != nil {
		log.Fatal(err)
	}

	// csv to html fileurl to file
	fileurl := "https://assets.apyhub.com/samples/sample.md"

	byt, err = apyhub.MarkdownToHtmlAsFile(fileurl)
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile("test.html", byt, 0777)
	if err != nil {
		log.Fatal(err)
	}

	// csv to html  file to url

	url, err := apyhub.MarkdownToHtmlAsUrl(File)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(url)

	// csv to html  fileurl to file
	fileurl = "https://assets.apyhub.com/samples/sample.md"

	url, err = apyhub.MarkdownToHtmlAsUrl(fileurl)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(url)
}
