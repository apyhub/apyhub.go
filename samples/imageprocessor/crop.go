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

func Crop() {

	File, err := os.Open("../test_file/flower.jpeg")
	if err != nil {
		log.Fatal(err)
	}

	//compress the image as file to file
	byt, err := apyhub.CropImageAsFile(File, 250, 350)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile("test.png", byt, 0777)
	if err != nil {
		log.Fatal(err)
	}

	fileurl := "https://assets.apyhub.com/samples/sample.jpg"

	//compress the image as fileurl to file
	byt, err = apyhub.CropImageAsFile(fileurl, 350, 350)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile("test.png", byt, 0777)
	if err != nil {
		log.Fatal(err)
	}

	//compress the image as file to url
	url, err := apyhub.CropImageAsURL(File, 350, 250)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(url)

	//compress the image as fileurl to url
	url, err = apyhub.CropImageAsURL(fileurl, 350, 250)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(url)
}
