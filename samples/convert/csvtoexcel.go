package convert

import (
	"fmt"
	"log"
	"os"

	apyhub "github.com/apyhub/apyhub.go"
)

func init() {
	apyhub.InitApyHub("Enter the apyhub token")
}

func Csvtoexcel() {
	File, err := os.Open("../test_file/sample.csv")
	if err != nil {
		log.Fatal(err)
	}

	// csv to excel file to file
	byt, err := apyhub.CsvToExcelAsFile(File)
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile("test.xlsx", byt, 0777)
	if err != nil {
		log.Fatal(err)
	}

	// csv to excel fileurl to file
	fileurl := "https://assets.apyhub.com/samples/sample.csv"

	byt, err = apyhub.CsvToExcelAsFile(fileurl)
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile("test.xlsx", byt, 0777)
	if err != nil {
		log.Fatal(err)
	}

	// csv to excel  file to url

	url, err := apyhub.CsvToExcelAsUrl(File)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(url)

	// csv to excel  fileurl to file
	fileurl = "https://assets.apyhub.com/samples/sample.csv"

	url, err = apyhub.CsvToExcelAsUrl(fileurl)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(url)
}
