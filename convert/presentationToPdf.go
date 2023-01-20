package convert

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	h "github.com/apyhub/apyhub/helper"
)

// Convert PPT to Pdf
// Output As File
func PresentationToPdfAsFile(conv interface{}, landscape bool) (byt []byte, err error) {
	switch input := conv.(type) {
	case string:
		jsn := h.Request{Url: input}
		byt, err = json.Marshal(jsn)
		if err != nil {
			return
		}
		url := fmt.Sprintf("%s?landscape=%t", h.PresentationUrltoPdfFile, landscape)
		return h.CallApyhubFromJson(url, byt)
	case io.Reader:
		fileByt, err := io.ReadAll(input)
		if err != nil {
			return nil, err
		}
		if h.GetFileExt(fileByt) == ".pptx" || h.GetFileExt(fileByt) == ".odp" {
			url := fmt.Sprintf("%s?landscape=%t", h.PresentationFiletoPdfFile, landscape)
			return h.CallApyhubFromFile("file", url, bytes.NewReader(fileByt), "sample.pptx")
		} else {
			return nil, h.ErrInvalidFormat
		}
	default:
		return nil, h.ErrTypeNotMatch
	}
}

func PresentationToPdfAsUrl(conv interface{}, landscape bool) (string, error) {
	var byt []byte
	var err error
	switch input := conv.(type) {
	case string:
		jsn := h.Request{Url: input}
		jsnByt, err := json.Marshal(jsn)
		if err != nil {
			return "", err
		}
		url := fmt.Sprintf("%s?landscape=%t", h.PresentationUrltoPdfUrl, landscape)
		byt, err = h.CallApyhubFromJson(url, jsnByt)
		if err != nil {
			return "", err
		}

	case io.Reader:
		fileByt, err := io.ReadAll(input)
		if err != nil {
			return "", err
		}
		if h.GetFileExt(fileByt) == ".pptx" || h.GetFileExt(fileByt) == ".odp" {
			url := fmt.Sprintf("%s?landscape=%t", h.PresentationFiletoPdfUrl, landscape)
			byt, err = h.CallApyhubFromFile("file", url, bytes.NewReader(fileByt), "sample.ppt")
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
