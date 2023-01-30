package imageprocessor

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	apyhub "github.com/apyhub/apyhub.go"
	"github.com/apyhub/apyhub.go/helper"
)

func init() {
	apyhub.InitApyHub("APT0rzEvAp76Izffo3ek7Zg7FpkK6csxzJ5puXCwWAw3lg0cFtj")
}

func Watermark() {

	image, err := os.Open("../test_file/sample.jpg")
	if err != nil {
		log.Fatal(err)
	}

	Watermarkimage, err := os.Open("../test_file/flower.jpeg")
	if err != nil {
		log.Fatal(err)
	}
	conv := helper.Watermark{
		Image:     image,
		Watermark: Watermarkimage,
	}
	//compress the image as file to file
	byt, err := apyhub.WatermarkImageAsFile(conv, 350, 250)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile("test.png", byt, 0777)
	if err != nil {
		log.Fatal(err)
	}

	imageurl := "https://assets.apyhub.com/samples/sample.jpg"
	waterimageurl := "https://assets.apyhub.com/samples/sample.jpg"

	urlfiles := helper.Watermark{
		Image:     imageurl,
		Watermark: waterimageurl,
	}

	//compress the image as fileurl to file
	byt, err = apyhub.WatermarkImageAsFile(urlfiles, 350, 250)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile("test.png", byt, 0777)
	if err != nil {
		log.Fatal(err)
	}

	//compress the image as file to url
	url, err := apyhub.WatermarkImageAsURL(conv, 250, 350)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(url)

	//compress the image as fileurl to url
	url, err = apyhub.WatermarkImageAsURL(urlfiles, 350, 250)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(url)
}
