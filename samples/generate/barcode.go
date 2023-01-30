package generate

import (
	"fmt"
	"log"
	"os"

	"github.com/apyhub/apyhub.go"
)

func init() {
	apyhub.InitApyHub("Enter your apyhub token")
}

func Barcode() {
	byt, err := apyhub.BarcodeAsFile("This is sample text.")
	if err != nil {
		log.Fatal(err)
	}

	// write bytes in specific file
	if err := os.WriteFile("./test_file/barcode.png", byt, 0666); err != nil {
		log.Fatal(err)
	}

	url, err := apyhub.BarcodeAsURL("This is sample text.")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(url)
}
