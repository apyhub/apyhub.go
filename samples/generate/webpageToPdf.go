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

func WebpageToPdf() {
	byt, err := apyhub.WebpageToPdfFile("https://gorm.io/docs/create.html", 0)
	if err != nil {
		log.Fatal(err)
	}

	// write bytes in specific file
	if err := os.WriteFile("./test_file/thumbnails.pdf", byt, 0666); err != nil {
		log.Fatal(err)
	}

	url, err := apyhub.WebpageToPdfURL("https://assets.apyhub.com/samples/sample.jpg", 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(url)
}
