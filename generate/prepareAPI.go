package generate

import (
	"bytes"
	"encoding/json"

	h "github.com/apyhub/apyhub.go/helper"
)

func prepareAPI(obj interface{}, method string, url string, IsURl bool) (byt []byte, err error) {
	byt, err = callApyhub(obj, method, url)
	if err != nil {
		return nil, err
	}
	if IsURl {
		var respJsn h.Response
		if err = json.Unmarshal(byt, &respJsn); err != nil {
			return nil, err
		}
		return []byte(respJsn.Data), nil
	}
	return
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
