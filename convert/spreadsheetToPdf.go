package convert

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	h "github.com/apyhub/apyhub/helper"
)

// Convert spreadsheet to Pdf
// Output As File
func SpreadsheetToPdfAsFile(conv interface{}, landscape bool) (byt []byte, err error) {
	switch input := conv.(type) {
	case string:
		jsn := h.Request{Url: input}
		byt, err = json.Marshal(jsn)
		if err != nil {
			return
		}
		url := fmt.Sprintf("%s?landscape=%t", h.SpreadsheetUrltoPdfFile, landscape)
		return h.CallApyhubFromJson(url, byt)
	case io.Reader:
		fileByt, err := io.ReadAll(input)
		if err != nil {
			return nil, err
		}

		if h.GetFileExt(fileByt) == ".xls" || h.GetFileExt(fileByt) == ".xlsx" || h.GetFileExt(fileByt) == ".ods" {
			url := fmt.Sprintf("%s?landscape=%t", h.SpreadsheetFiletoPdfFile, landscape)
			return h.CallApyhubFromFile("file", url, bytes.NewReader(fileByt), "sample.xlsx")
		} else {
			return nil, h.ErrInvalidFormat
		}

	default:
		return nil, h.ErrTypeNotMatch
	}
}

func SpreadsheetToPdfAsUrl(conv interface{}, landscape bool) (string, error) {
	var byt []byte
	var err error
	switch input := conv.(type) {
	case string:
		jsn := h.Request{Url: input}
		jsnByt, err := json.Marshal(jsn)
		if err != nil {
			return "", err
		}
		url := fmt.Sprintf("%s?landscape=%t", h.SpreadsheetUrltoPdfUrl, landscape)
		byt, err = h.CallApyhubFromJson(url, jsnByt)
		if err != nil {
			return "", err
		}

	case io.Reader:
		fileByt, err := io.ReadAll(input)
		if err != nil {
			return "", err
		}
		if h.GetFileExt(fileByt) == ".xls" || h.GetFileExt(fileByt) == ".xlsx" || h.GetFileExt(fileByt) == ".ods" {
			url := fmt.Sprintf("%s?landscape=%t", h.SpreadsheetFiletoPdfUrl, landscape)
			byt, err = h.CallApyhubFromFile("file", url, bytes.NewReader(fileByt), "sample.xlsx")
			if err != nil {
				return "", err
			}
		} else {
			return "", h.ErrInvalidFormat
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
