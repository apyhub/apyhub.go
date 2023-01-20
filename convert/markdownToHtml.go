package convert

import (
	"bytes"
	"encoding/json"
	"io"

	h "github.com/apyhub/apyhub.go/helper"
)

// Convert Csv to Excel
// Output As a File
func MarkdownToHtmlAsFile(conv interface{}) (byt []byte, err error) {
	switch input := conv.(type) {
	case string:
		jsn := h.Request{Url: input}
		byt, err = json.Marshal(jsn)
		if err != nil {
			return
		}
		return h.CallApyhubFromJson(h.MdUrltoHtmlFile, byt)
	case io.Reader:
		fileByt, err := io.ReadAll(input)
		if err != nil {
			return nil, err
		}
		if h.GetFileExt(fileByt) != ".txt" {
			return nil, h.ErrInvalidFormat
		}
		return h.CallApyhubFromFile("file", h.MdFiletoHtmlFile, bytes.NewReader(fileByt), "sample.md")
	default:
		return nil, h.ErrTypeNotMatch
	}
}

// Convert Csv to Excel
// Output As a Link (AWS S3 link). It is valid upto 15 min.
func MarkdownToHtmlAsUrl(conv interface{}) (string, error) {
	var byt []byte
	var err error
	switch input := conv.(type) {
	case string:
		jsn := h.Request{Url: input}
		jsnByt, err := json.Marshal(jsn)
		if err != nil {
			return "", err
		}
		byt, err = h.CallApyhubFromJson(h.MdUrltoHtmlUrl, jsnByt)
		if err != nil {
			return "", err
		}

	case io.Reader:
		fileByt, err := io.ReadAll(input)
		if err != nil {
			return "", err
		}
		if h.GetFileExt(fileByt) != ".txt" {
			return "", h.ErrInvalidFormat
		}
		byt, err = h.CallApyhubFromFile("file", h.MdFiletoHtmlUrl, bytes.NewReader(fileByt), "sample.md")
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
