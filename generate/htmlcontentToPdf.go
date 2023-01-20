package generate

import (
	h "github.com/apyhub/apyhub.go/helper"
)

func HtmlContentToPdfFile(content string) ([]byte, error) {
	return prepareAPI(h.ContentRequest{Content: content}, "POST", h.HtmlContentToPdfFile, false)
}

func HtmlContentToPdfAsUrl(content string) (string, error) {
	byt, err := prepareAPI(h.ContentRequest{Content: content}, "POST", h.HtmlContentToPdfURL, true)
	if err != nil {
		return "", err
	}
	return string(byt), nil
}
