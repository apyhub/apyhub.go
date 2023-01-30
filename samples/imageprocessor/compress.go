package imageprocessor

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	apyhub "github.com/apyhub/apyhub.go"
)

func init() {
	apyhub.InitApyHub("Enter Apyhub token")
}

func Compress() {

	File, err := os.Open("../test_file/sample.jpg")
	if err != nil {
		log.Fatal(err)
	}

	//compress the image as file to file
	byt, err := apyhub.CompressImageAsFile(File, 20)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile("test.png", byt, 0777)
	if err != nil {
		log.Fatal(err)
	}

	fileurl := "https://assets.apyhub.com/samples/sample.jpg"

	//compress the image as fileurl to file
	byt, err = apyhub.CompressImageAsFile(fileurl, 100)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile("test.png", byt, 0777)
	if err != nil {
		log.Fatal(err)
	}

	// compress the image as file to url
	url, err := apyhub.CompressImageAsURL(File, 100)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(url)

	//compress the image as fileurl to url
	url, err = apyhub.CompressImageAsURL(fileurl, 100)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(url)
}
