package convert

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	h "github.com/apyhub/apyhub/helper"
)

// Convert Html to Pdf
// Output As File
func HtmlToPdfAsFile(conv interface{}, landscape bool) (byt []byte, err error) {
	switch input := conv.(type) {
	case string:
		jsn := h.Request{Url: input}
		byt, err = json.Marshal(jsn)
		if err != nil {
			return
		}
		url := fmt.Sprintf("%s?landscape=%t", h.HtmlUrltoPdfFile, landscape)
		return h.CallApyhubFromJson(url, byt)
	case io.Reader:
		fileByt, err := io.ReadAll(input)
		if err != nil {
			return nil, err
		}
		if h.GetFileExt(fileByt) != ".html" {
			return nil, h.ErrInvalidFormat
		}
		url := fmt.Sprintf("%s?landscape=%t", h.HtmlFiletoPdfFile, landscape)
		return h.CallApyhubFromFile("file", url, bytes.NewReader(fileByt), "sample.html")
	default:
		return nil, h.ErrTypeNotMatch
	}
}

func HtmlToPdfAsUrl(conv interface{}, landscape bool) (string, error) {
	var byt []byte
	var err error
	switch input := conv.(type) {
	case string:
		jsn := h.Request{Url: input}
		jsnByt, err := json.Marshal(jsn)
		if err != nil {
			return "", err
		}
		url := fmt.Sprintf("%s?landscape=%t", h.HtmlUrltoPdfUrl, landscape)
		byt, err = h.CallApyhubFromJson(url, jsnByt)
		if err != nil {
			return "", err
		}

	case io.Reader:
		fileByt, err := io.ReadAll(input)
		if err != nil {
			return "", err
		}
		if h.GetFileExt(fileByt) != ".html" {
			return "", h.ErrInvalidFormat
		}
		url := fmt.Sprintf("%s?landscape=%t", h.HtmlFiletoPdfUrl, landscape)
		byt, err = h.CallApyhubFromFile("file", url, bytes.NewReader(fileByt), "index.html")
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
