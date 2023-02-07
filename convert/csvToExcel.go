package convert

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/build"
	"io"
	"log"
	"os"

	h "github.com/apyhub/apyhub.go/helper"
	"github.com/princjef/gomarkdoc"
	"github.com/princjef/gomarkdoc/lang"
	"github.com/princjef/gomarkdoc/logger"
)

func Md() {
	// Create a renderer to output data
	out, err := gomarkdoc.NewRenderer()
	if err != nil {
		log.Fatal(err)
	}

	// wd, err := filepath.Abs(filepath.Dir(os.Args[0]))
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("wd:", wd)
	buildPkg, err := build.ImportDir(wd, build.ImportComment)
	if err != nil {
		log.Fatal(err)
	}

	// Create a documentation package from the build representation of our
	// package.
	logmd := logger.New(logger.DebugLevel)
	pkg, err := lang.NewPackageFromBuild(logmd, buildPkg)
	if err != nil {
		log.Fatal(err)
	}
	file, _ := os.Create("read1.md")

	mdstring, _ := out.Package(pkg)
	file.WriteString(mdstring)
	// Write the documentation out to console.
	fmt.Println()
}

// Convert Csv to Excel
// Output As a File
func CsvToExcelAsFile(conv interface{}) ([]byte, error) {
	switch input := conv.(type) {
	case string:
		jsn := h.Request{Url: input}
		byt, err := json.Marshal(jsn)
		if err != nil {
			return nil, err
		}
		return h.CallApyhubFromJson(h.CsvUrltoExcelFile, byt)
	case io.Reader:
		fileByt, err := io.ReadAll(input)
		if err != nil {
			return nil, err
		}
		if h.GetFileExt(fileByt) != ".csv" {
			return nil, h.ErrInvalidFormat
		}
		return h.CallApyhubFromFile("file", h.CsvFiletoExcelFile, bytes.NewReader(fileByt), "sample.csv")
	default:
		return nil, h.ErrTypeNotMatch
	}
}

// Convert Csv to Excel
// Output As a Link (AWS S3 link). It is valid upto 15 min.
func CsvToExcelAsURL(conv interface{}) (string, error) {
	var byt []byte
	var err error
	switch input := conv.(type) {
	case string:
		jsn := h.Request{Url: input}
		jsnByt, err := json.Marshal(jsn)
		if err != nil {
			return "", err
		}
		byt, err = h.CallApyhubFromJson(h.CsvUrltoExcelUrl, jsnByt)
		if err != nil {
			return "", err
		}

	case io.Reader:
		fileByt, err := io.ReadAll(input)
		if err != nil {
			return "", err
		}
		if h.GetFileExt(fileByt) != ".csv" {
			return "", h.ErrInvalidFormat
		}
		byt, err = h.CallApyhubFromFile("file", h.CsvFiletoExcelUrl, bytes.NewReader(fileByt), "sample.csv")
		if err != nil {
			return "", err
		}
	default:
		return "", h.ErrTypeNotMatch
	}

	var res h.Response
	if err = json.Unmarshal(byt, &res); err != nil {
		return "", err
	}
	return res.Data, nil
}
