package extract

import (
	"bytes"
	"encoding/json"
	"io"

	h "github.com/apyhub/apyhub.go/helper"
)

func UnArchiveAsURL(conv interface{}) ([]string, error) {
	switch input := conv.(type) {
	case string:
		return prepareAPIUnZip(h.UnZip{Url: input}, "POST", h.UnArchiveAsURLToURL)

	case io.Reader:
		var respUnZipJsn h.UnZipResponse
		fileByt, err := io.ReadAll(input)
		if err != nil {
			return nil, err
		}
		byt, err := h.CallApyhubFromFile("file", h.UnArchiveFileToURL, bytes.NewReader(fileByt), "sample"+h.GetFileExt(fileByt))
		if err != nil {
			return nil, err
		}
		if err = json.Unmarshal(byt, &respUnZipJsn); err != nil {
			return nil, err
		}
		return respUnZipJsn.Data, nil
	default:
		return nil, h.ErrTypeNotMatch

	}
}

func SecureUnArchiveAsURL(password string, conv interface{}) ([]string, error) {
	switch input := conv.(type) {
	case string:
		SecureZip := h.SecureUnZip{Url: input, Password: password}
		return prepareAPIUnZip(SecureZip, "POST", h.SecureUnArchiveURLToURL)
	case io.Reader:
		var respUnZipJsn h.UnZipResponse

		fileByt, err := io.ReadAll(input)
		if err != nil {
			return nil, err
		}

		byt, err := callApyhubSecure(h.SecureUnArchiveFileToURL, bytes.NewReader(fileByt), "sample"+h.GetFileExt(fileByt), password)
		if err != nil {
			return nil, err
		}
		if err := json.Unmarshal(byt, &respUnZipJsn); err != nil {
			return nil, err
		}
		return respUnZipJsn.Data, nil
	default:
		return nil, h.ErrTypeNotMatch
	}
}

func callApyhubSecure(url string, input *bytes.Reader, filename string, password string) ([]byte, error) {
	body, writer, err := h.PrepareBody("file", input, filename)
	if err != nil {
		return nil, err
	}

	if err = writer.WriteField("password", password); err != nil {
		return nil, err
	}
	writer.Close()
	return h.CallApyhub("POST", url, body, writer.FormDataContentType())
}

func prepareAPIUnZip(obj interface{}, method string, url string) (byt []string, err error) {
	jsnByt, err := callApyhub(obj, method, url)
	if err != nil {
		return nil, err
	}
	var respUnZipJsn h.UnZipResponse
	if err = json.Unmarshal(jsnByt, &respUnZipJsn); err != nil {
		return nil, err
	}
	return respUnZipJsn.Data, nil
}

func callApyhub(obj interface{}, method string, url string) (byt []byte, err error) {
	jsnbyt, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	byt, err = h.CallApyhub(method, url, bytes.NewReader(jsnbyt))
	if err != nil {
		return nil, err
	}

	return
}
