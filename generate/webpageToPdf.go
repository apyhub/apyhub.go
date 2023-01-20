package generate

import (
	"fmt"

	h "github.com/apyhub/apyhub.go/helper"
)

func WebpageToPdfFile(url string, margin float64) ([]byte, error) {
	url = fmt.Sprintf("%s?margin=%f", url, margin)
	return prepareAPI(h.Request{Url: url}, "POST", h.WebpageToPdfFile, false)
}

func WebpageToPdfURL(url string, margin float64) (string, error) {
	url = fmt.Sprintf("%s?margin=%f", url, margin)
	byt, err := prepareAPI(h.Request{Url: url}, "POST", h.WebpageToPdfURL, true)
	if err != nil {
		return "", err
	}
	return string(byt), nil
}
