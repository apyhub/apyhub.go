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

func Screenshot() {
	byt, err := apyhub.ScreenshotWebpageAsFile("https://www.google.com/", 1, 0)
	if err != nil {
		log.Fatal(err)
	}

	// write bytes in specific file
	if err := os.WriteFile("./test_file/screenshot.png", byt, 0666); err != nil {
		log.Fatal(err)
	}

	url, err := apyhub.ScreenshotWebpageAsURL("https://www.google.com/", 1, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(url)
}
