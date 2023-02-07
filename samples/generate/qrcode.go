package generate

import (
	"fmt"
	"log"
	"os"

	apyhub "github.com/apyhub/apyhub.go"
)

func init() {
	apyhub.InitApyHub("Enter your apyhub token")
}

func Qrcode() {
	byt, err := apyhub.QrcodeAsFile("This is sample text.")
	if err != nil {
		log.Fatal(err)
	}

	// write bytes in specific file
	if err := os.WriteFile("./test_file/qrcode.png", byt, 0666); err != nil {
		log.Fatal(err)
	}

	url, err := apyhub.QrcodeAsURL("This is sample text.")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(url)
}
