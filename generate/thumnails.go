package generate

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	h "github.com/apyhub/apyhub.go/helper"
)

// Image Crop
// Output As a File
func ThumnailsImageAsFile(conv interface{}, width int, height int) (byt []byte, err error) {
	switch input := conv.(type) {
	case string:
		url := fmt.Sprintf("%s?width=%d&height=%d", h.ThumnailsImageUrlToFile, width, height)
		return prepareAPI(h.Request{Url: input}, "GET", url, false)

	case io.Reader:
		filebyt, err := io.ReadAll(input)
		if err != nil {
			return nil, err
		}
		url := fmt.Sprintf("%s?width=%d&height=%d", h.ThumnailsImageFileToFile, width, height)
		if !h.CheckImageExt(filebyt) {
			return nil, h.ErrInvalidFormat
		}
		return h.CallApyhubFromFile("image", url, bytes.NewReader(filebyt), "sample.png")
	default:
		return nil, h.ErrTypeNotMatch
	}
}

// Image Crop
// Output As a Link (AWS S3 link). It is valid upto 15 min.
func ThumnailsImageAsURL(conv interface{}, width int, height int) (string, error) {
	var byt []byte
	switch input := conv.(type) {
	case string:
		url := fmt.Sprintf("%s?width=%d&height=%d", h.ThumnailsImageUrlToUrl, width, height)
		byt, err := prepareAPI(h.Request{Url: input}, "POST", url, true)
		if err != nil {
			return "", err
		}
		return string(byt), nil

	case io.Reader:
		filebyt, err := io.ReadAll(input)
		if err != nil {
			return "", err
		}
		url := fmt.Sprintf("%s?width=%d&height=%d", h.ThumnailsImageFileToUrl, width, height)
		if !h.CheckImageExt(filebyt) {
			return "", h.ErrInvalidFormat
		}

		byt, err = h.CallApyhubFromFile("image", url, bytes.NewReader(filebyt), "sample.png")
		if err != nil {
			return "", err
		}

		var res h.Response
		if err = json.Unmarshal(byt, &res); err != nil {
			return "", err
		}
		return res.Data, nil
	default:
		return "", h.ErrTypeNotMatch
	}

}
