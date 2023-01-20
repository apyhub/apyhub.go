package extract

import (
	"bytes"
	"encoding/json"
	"io"

	h "github.com/apyhub/apyhub/helper"
)

func TextFromPdf(urlorFile interface{}) (data string, err error) {
	byt := []byte{}
	switch input := urlorFile.(type) {
	case string:
		jsn := h.Request{Url: input}
		byt, err = json.Marshal(jsn)
		if err != nil {
			return
		}
		byt, err = h.CallApyhubFromJson(h.TextFromPdfUrl, byt)
	case io.Reader:
		fileByt, err := io.ReadAll(input)
		if err != nil {
			return "", err
		}
		if h.GetFileExt(fileByt) != ".pdf" {
			return "", h.ErrInvalidFormat
		}
		byt, err = h.CallApyhubFromFile("file", h.TextFromPdfFile, bytes.NewReader(fileByt), "sample.pdf")
	default:
		return "", h.ErrTypeNotMatch
	}
	if err != nil {
		return
	}

	respJsn := h.Response{}
	if err = json.Unmarshal(byt, &respJsn); err != nil {
		return
	}
	return respJsn.Data, nil
}
