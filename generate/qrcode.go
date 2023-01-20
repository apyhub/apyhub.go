package generate

import (
	h "github.com/apyhub/apyhub.go/helper"
)

func QrcodeAsFile(content string) (byt []byte, err error) {
	return prepareAPI(h.ContentRequest{Content: content}, "POST", h.QrcodeAsFile, false)
}

func QrcodeAsURL(content string) (string, error) {
	byt, err := prepareAPI(h.ContentRequest{Content: content}, "POST", h.QrcodeAsURL, true)
	if err != nil {
		return "", err
	}
	return string(byt), nil
}
