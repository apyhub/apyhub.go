package generate

import (
	"fmt"

	h "github.com/apyhub/apyhub/helper"
)

func ScreenshotWebpageAsFile(url string, quality int, delay int) ([]byte, error) {
	url = fmt.Sprintf("%s?url=%s&quality=%d&delay=%d", h.ScreenshotWebpageAsFile, url, quality, delay)
	return prepareAPI(h.Request{Url: url}, "GET", url, false)
}

func ScreenshotWebpageAsURL(url string, quality int, delay int) (string, error) {
	url = fmt.Sprintf("%s?url=%s&quality=%d&delay=%d", h.ScreenshotWebpageAsURL, url, quality, delay)
	byt, err := prepareAPI(h.Request{Url: url}, "GET", url, true)
	if err != nil {
		return "", err
	}
	return string(byt), nil
}
