package extract

import (
	"bytes"
	"encoding/json"
	"io"

	h "github.com/apyhub/apyhub.go/helper"
)

func ImageMetadata(urlorFile interface{}) (data map[string]interface{}, err error) {
	byt := []byte{}
	switch input := urlorFile.(type) {
	case string:
		jsn := h.Request{Url: input}
		byt, err = json.Marshal(jsn)
		if err != nil {
			return
		}
		byt, err = h.CallApyhubFromJson(h.ImageMetadataFromUrl, byt)
	case io.Reader:
		fileByt, err := io.ReadAll(input)
		if err != nil {
			return nil, err
		}
		if !h.CheckImageExt(fileByt) {
			return nil, h.ErrInvalidFormat
		}
		byt, err = h.CallApyhubFromFile("image", h.ImageMetadataFromFile, bytes.NewReader(fileByt), "sample.jpg")
	default:
		return nil, h.ErrTypeNotMatch
	}
	if err != nil {
		return
	}

	respJsn := h.MetaResponse{}
	if err = json.Unmarshal(byt, &respJsn); err != nil {
		return
	}
	return respJsn.Data, nil
}
