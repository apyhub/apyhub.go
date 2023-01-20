package generate

import (
	h "github.com/apyhub/apyhub.go/helper"
)

func BarcodeAsFile(content string) ([]byte, error) {
	return prepareAPI(h.ContentRequest{Content: content}, "POST", h.BarcodeAsFile, false)
}

func BarcodeAsURL(content string) (string, error) {
	byt, err := prepareAPI(h.ContentRequest{Content: content}, "POST", h.BarcodeAsURL, true)
	if err != nil {
		return "", err
	}
	return string(byt), nil
}
