package generate

import (
	"fmt"
	"log"
	"os"

	apyhub "github.com/apyhub/apyhub.go"
)

func init() {
	apyhub.InitApyHub("Enter apyhub token")
}

func Archive() {

	// file archive without protection As a files or file urls output as file.
	Urls := []string{
		"https://assets.apyhub.com/samples/sample.jpg",
		"https://assets.apyhub.com/samples/sample.docx",
		"https://assets.apyhub.com/samples/sample.pdf",
	}

	byt, err := apyhub.ArchiveAsFile(Urls)
	if err != nil {
		log.Fatal(err)
	}

	// write bytes in specific file
	if err := os.WriteFile("./sample.zip", byt, 0666); err != nil {
		log.Fatal(err)
	}

	// file archive without protection As a files or file urls output return as a url.
	url, err := apyhub.ArchiveAsURL(Urls)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(url)

	// file archive with protection As a files or file urls output return as a file.
	Urls = []string{
		"https://assets.apyhub.com/samples/sample.jpg",
		"https://assets.apyhub.com/samples/sample.docx",
		"https://assets.apyhub.com/samples/sample.pdf",
	}

	byt, err = apyhub.SecureArchiveFile("apyhub", Urls)
	if err != nil {
		log.Fatal(err)
	}

	// write bytes in specific file
	if err := os.WriteFile(".../test_file/sampleProtection.zip", byt, 0666); err != nil {
		log.Fatal(err)
	}

	// file archive with protection As a files or file urls output return as a url.
	url, err = apyhub.SecureArchiveAsURL("password", Urls)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(url)
}

func Unarchive() {

	//file unarchive without protection As file or url and output return as slice of urls
	File, err := os.Open("./test_file/testziplink.zip.")
	if err != nil {
		log.Fatal(err)
	}
	output, err := apyhub.UnArchiveAsURL(File)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(output)

	//file unarchive with protection As file or url and output return as slice of urls
	output, err = apyhub.SecureUnArchiveAsURL("some-secure-password", "https://assets.apyhub.com/samples/secure-archive.zip")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(output)
}
