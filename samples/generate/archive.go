package generate

import (
	"fmt"
	"log"
	"os"

	"github.com/apyhub/apyhub.go"
	h "github.com/apyhub/apyhub.go/helper"
)

func init() {
	apyhub.InitApyHub("APT0rzEvAp76Izffo3ek7Zg7FpkK6csxzJ5puXCwWAw3lg0cFtj")
}

func Archive() {

	// file archive without protection
	zip := h.Zip{
		Urls: []string{
			"https://assets.apyhub.com/samples/sample.jpg",
			"https://assets.apyhub.com/samples/sample.docx",
			"https://assets.apyhub.com/samples/sample.pdf",
		},
	}
	byt, err := apyhub.ArchiveAsFile(zip)
	if err != nil {
		log.Fatal(err)
	}

	// write bytes in specific file
	if err := os.WriteFile("./sample.zip", byt, 0666); err != nil {
		log.Fatal(err)
	}

	url, err := apyhub.ArchiveAsURL(zip)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(url)

	// file archive with protection
	secureZip := h.ZipSecure{
		Urls: []string{
			"https://assets.apyhub.com/samples/sample.jpg",
			"https://assets.apyhub.com/samples/sample.docx",
			"https://assets.apyhub.com/samples/sample.pdf",
		},
		Password: "apyhub",
	}
	byt, err = apyhub.SecureArchiveFile(secureZip)
	if err != nil {
		log.Fatal(err)
	}

	// write bytes in specific file
	if err := os.WriteFile(".../test_file/sampleProtection.zip", byt, 0666); err != nil {
		log.Fatal(err)
	}

	url, err = apyhub.SecureArchiveAsURL(secureZip)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(url)
}
