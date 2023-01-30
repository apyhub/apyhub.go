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

func Thumbnails() {

	File, err := os.Open("./test_file/flower.jpeg")
	if err != nil {
		log.Fatal(err)
	}

	//thumbnails the image as file to file
	byt, err := apyhub.ThumnailsImageAsFile(File, 200, 250)
	if err != nil {
		log.Fatal(err)
	}

	// write bytes in specific file
	if err := os.WriteFile("./test_file/thumbnails.png", byt, 0666); err != nil {
		log.Fatal(err)
	}

	//thumbnails the image as fileurl to file
	byt, err = apyhub.ThumnailsImageAsFile("https://assets.apyhub.com/samples/sample.jpg", 200, 250)
	if err != nil {
		log.Fatal(err)
	}

	// write bytes in specific file
	if err := os.WriteFile("./test_file/thumbnails.png", byt, 0666); err != nil {
		log.Fatal(err)
	}

	url, err := apyhub.ThumnailsImageAsURL(File, 200, 250)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(url)

	//thumbnails the image as fileurl to file
	url, err = apyhub.ThumnailsImageAsURL("https://assets.apyhub.com/samples/sample.jpg", 200, 250)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(url)
}
