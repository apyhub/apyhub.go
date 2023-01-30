package extract

import (
	"fmt"
	"log"
	"os"

	apyhub "github.com/apyhub/apyhub.go"
)

func init() {
	apyhub.InitApyHub("Enter Apyhub token")
}

func Textfromword() {

	File, err := os.Open("../test_file/word.docx")
	if err != nil {
		log.Fatal(err)
	}

	// To extract the data from word file
	data, err := apyhub.TextFromWord(File)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("textfrompdf:", data)

	// To extract the data from word url
	fileurl := "https://assets.apyhub.com/samples/sample.docx"

	data, err = apyhub.TextFromWord(fileurl)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("textfrompdf", data)
}
