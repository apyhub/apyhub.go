package extract

import (
	"bytes"
	"encoding/json"
	"io"

	h "github.com/apyhub/apyhub/helper"
)

func TextFromWord(urlorFile interface{}) (data string, err error) {
	byt := []byte{}
	switch input := urlorFile.(type) {
	case string:
		jsn := h.Request{Url: input}
		byt, err = json.Marshal(jsn)
		if err != nil {
			return
		}
		byt, err = h.CallApyhubFromJson(h.TextFromWordUrl, byt)
	case io.Reader:
		fileByt, err := io.ReadAll(input)
		if err != nil {
			return "", err
		}
		if !h.CheckDocumentExt(fileByt) {
			return "", h.ErrInvalidFormat
		}
		byt, err = h.CallApyhubFromFile("file", h.TextFromWordFile, bytes.NewReader(fileByt), "sample.docx")
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
