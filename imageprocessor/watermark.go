package imageprocessor

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"

	h "github.com/apyhub/apyhub/helper"
)

// Image Crop
// Output As a File
func WatermarkImageAsFile(watermark h.Watermark, width int, height int) (byt []byte, err error) {
	switch image := watermark.Image.(type) {
	case string:
		WatermarkString, ok := watermark.Watermark.(string)
		if !ok {
			err = h.ErrNotSameType
			return
		}

		// prepare json
		jsn := h.WatermarkRequest{ImageUrl: image, WatermarkUrl: WatermarkString}
		byt, err = json.Marshal(jsn)
		if err != nil {
			return
		}
		url := fmt.Sprintf("%s?width=%d&height=%d", h.WaterMarkImageUrlToFile, width, height)
		return h.CallApyhubFromJson(url, byt)
	case io.Reader:
		WatermarkReader, ok := watermark.Watermark.(io.Reader)
		if !ok {
			err = h.ErrNotSameType
			return
		}
		url := fmt.Sprintf("%s?width=%d&height=%d", h.WaterMarkImageFileToFile, width, height)

		// prepare watermark body
		body, writer, err := prepareBody(image, WatermarkReader)
		if err != nil {
			return nil, err
		}
		writer.Close()
		return h.CallApyhub("POST", url, body, writer.FormDataContentType())
	default:
		return nil, h.ErrTypeNotMatch
	}
}

// Image Crop
// Output As a Link (AWS S3 link). It is valid upto 15 min.
func WatermarkImageAsURL(watermark h.Watermark, width int, height int) (string, error) {
	var byt []byte
	var err error
	switch image := watermark.Image.(type) {
	case string:
		WatermarkString, ok := watermark.Watermark.(string)
		if !ok {
			return "", h.ErrNotSameType
		}

		// prepare json
		jsn := h.WatermarkRequest{ImageUrl: image, WatermarkUrl: WatermarkString}
		jsnByt, err := json.Marshal(jsn)
		if err != nil {
			return "", err
		}
		url := fmt.Sprintf("%s?width=%d&height=%d", h.WaterMarkImageUrlToUrl, width, height)
		byt, err = h.CallApyhubFromJson(url, jsnByt)
		if err != nil {
			return "", err
		}
	case io.Reader:
		WatermarkReader, ok := watermark.Watermark.(io.Reader)
		if !ok {
			return "", h.ErrNotSameType
		}
		url := fmt.Sprintf("%s?width=%d&height=%d", h.WaterMarkImageFileToUrl, width, height)

		// prepare watermark body
		body, writer, err := prepareBody(image, WatermarkReader)
		if err != nil {
			return "", err
		}
		writer.Close()
		byt, err = h.CallApyhub("POST", url, body, writer.FormDataContentType())
		if err != nil {
			return "", err
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

func prepareBody(imageReader io.Reader, watermarkReader io.Reader) (*bytes.Buffer, *multipart.Writer, error) {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	imagebyt, err := io.ReadAll(imageReader)
	if err != nil {
		return nil, nil, err
	}
	if !h.CheckImageExt(imagebyt) {
		return nil, nil, h.ErrInvalidFormat
	}
	watermarkbyt, err := io.ReadAll(watermarkReader)
	if err != nil {
		return nil, nil, err
	}
	if !h.CheckImageExt(watermarkbyt) {
		return nil, nil, h.ErrInvalidFormat
	}
	image, err := writer.CreateFormFile("image", "image.png")
	if err != nil {
		return nil, nil, err
	}

	if _, err := io.Copy(image, bytes.NewReader(imagebyt)); err != nil {
		return nil, nil, err
	}

	watermark, err := writer.CreateFormFile("watermark", "watermark.png")
	if err != nil {
		return nil, nil, err
	}

	if _, err := io.Copy(watermark, bytes.NewReader(watermarkbyt)); err != nil {
		return nil, nil, err
	}
	return body, writer, nil
}
