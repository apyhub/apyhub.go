package convert

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	h "github.com/apyhub/apyhub.go/helper"
)

// Convert Csv to Excel
// Output As a File
func WordToPdfAsFile(conv interface{}, landscape bool) (byt []byte, err error) {
	switch input := conv.(type) {
	case string:
		jsn := h.Request{Url: input}
		byt, err = json.Marshal(jsn)
		if err != nil {
			return
		}
		url := fmt.Sprintf("%s?landscape=%t", h.WordUrltoPdfFile, landscape)
		return h.CallApyhubFromJson(url, byt)
	case io.Reader:
		fileByt, err := io.ReadAll(input)
		if err != nil {
			return nil, err
		}
		if h.GetFileExt(fileByt) == ".doc" || h.GetFileExt(fileByt) == ".docx" || h.GetFileExt(fileByt) == ".odt" || h.GetFileExt(fileByt) == ".rtf" {
			url := fmt.Sprintf("%s?landscape=%t", h.WordFiletoPdfFile, landscape)
			return h.CallApyhubFromFile("file", url, bytes.NewReader(fileByt), "sample.doc")
		} else {
			return nil, h.ErrInvalidFormat
		}
	default:
		return nil, h.ErrTypeNotMatch
	}
}

// Convert Csv to Excel
// Output As a Link (AWS S3 link). It is valid upto 15 min.
func WordToPdfAsUrl(conv interface{}, landscape bool) (string, error) {
	var byt []byte
	var err error
	switch input := conv.(type) {
	case string:
		jsn := h.Request{Url: input}
		jsnByt, err := json.Marshal(jsn)
		if err != nil {
			return "", err
		}

		url := fmt.Sprintf("%s?landscape=%t", h.WordUrltoPdfUrl, landscape)
		byt, err = h.CallApyhubFromJson(url, jsnByt)
		if err != nil {
			return "", err
		}

	case io.Reader:
		fileByt, err := io.ReadAll(input)
		if err != nil {
			return "", err
		}
		if h.GetFileExt(fileByt) == ".doc" || h.GetFileExt(fileByt) == ".docx" || h.GetFileExt(fileByt) == ".odt" || h.GetFileExt(fileByt) == ".rtf" {
			url := fmt.Sprintf("%s?landscape=%t", h.WordFiletoPdfUrl, landscape)
			byt, err = h.CallApyhubFromFile("file", url, bytes.NewReader(fileByt), "sample.doc")
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
