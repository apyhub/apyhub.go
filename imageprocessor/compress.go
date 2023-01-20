package imageprocessor

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	h "github.com/apyhub/apyhub.go/helper"
)

// Image Compression
// Output As a File
func CompressImageAsFile(conv interface{}, quality int) (byt []byte, err error) {
	switch input := conv.(type) {
	case string:
		jsn := h.Request{Url: input}
		byt, err = json.Marshal(jsn)
		if err != nil {
			return
		}
		url := fmt.Sprintf("%s?quality=%d", h.CompressImageUrlToFile, quality)
		return h.CallApyhubFromJson(url, byt)
	case io.Reader:
		fileByt, err := io.ReadAll(input)
		if err != nil {
			return nil, err
		}
		url := fmt.Sprintf("%s?quality=%d", h.CompressImageFileToFile, quality)
		if h.CheckImageExt(fileByt) {
			return h.CallApyhubFromFile("image", url, bytes.NewReader(fileByt), "sample.png")
		} else {
			return nil, h.ErrInvalidFormat
		}
	default:
		return nil, h.ErrTypeNotMatch
	}
}

// Image Compression
// Output As a Link (AWS S3 link). It is valid upto 15 min.
func CompressImageAsURL(conv interface{}, quality int) (string, error) {
	var byt []byte
	var err error
	switch input := conv.(type) {
	case string:
		jsn := h.Request{Url: input}
		jsnByt, err := json.Marshal(jsn)
		if err != nil {
			return "", err
		}
		url := fmt.Sprintf("%s?quality=%d", h.CompressImageUrlToUrl, quality)
		byt, err = h.CallApyhubFromJson(url, jsnByt)
		if err != nil {
			return "", err
		}
	case io.Reader:
		fileByt, err := io.ReadAll(input)
		if err != nil {
			return "", err
		}
		url := fmt.Sprintf("%s?quality=%d", h.CompressImageFileToUrl, quality)
		if h.CheckImageExt(fileByt) {
			byt, err = h.CallApyhubFromFile("image", url, bytes.NewReader(fileByt), "sample.png")
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
	err = json.Unmarshal(byt, &res)
	if err != nil {
		return "", err
	}
	return res.Data, nil
}
