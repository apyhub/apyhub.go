package imageprocessor

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	apyhub "github.com/apyhub/apyhub.go"
)

func init() {
	apyhub.InitApyHub("APT0rzEvAp76Izffo3ek7Zg7FpkK6csxzJ5puXCwWAw3lg0cFtj")
}

func Resize() {

	File, err := os.Open("../test_file/sample.jpg")
	if err != nil {
		log.Fatal(err)
	}

	//compress the image as file to file
	byt, err := apyhub.ResizeImageAsFile(File, 10, 250)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile("test.png", byt, 0777)
	if err != nil {
		log.Fatal(err)
	}

	fileurl := "https://assets.apyhub.com/samples/sample.jpg"

	//compress the image as fileurl to file
	byt, err = apyhub.ResizeImageAsFile(fileurl, 350, 250)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile("test.png", byt, 0777)
	if err != nil {
		log.Fatal(err)
	}

	//compress the image as file to url
	url, err := apyhub.ResizeImageAsURL(File, 350, 250)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(url)

	//compress the image as fileurl to url
	url, err = apyhub.ResizeImageAsURL(fileurl, 350, 250)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(url)
}
