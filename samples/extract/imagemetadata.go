package extract

import (
	"fmt"
	"log"
	"os"

	apyhub "github.com/apyhub/apyhub.go"
)

func init() {
	apyhub.InitApyHub("Enter apyhub token")
}

func Imagemetadata() {

	File, err := os.Open("../test_file/sample.jpg")
	if err != nil {
		log.Fatal(err)
	}

	// To collect metadata from image file
	data, err := apyhub.ImageMetadata(File)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("image meta data:%+v", data)

	fileurl := "https://assets.apyhub.com/samples/sample.jpg"

	// To collect metadata from imageurl
	data, err = apyhub.ImageMetadata(fileurl)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("image meta data:%+v", data)
}
