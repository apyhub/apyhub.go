package generate

import (
	"bytes"
	"encoding/json"

	h "github.com/apyhub/apyhub.go/helper"
)

func prepareAPI(obj interface{}, method string, url string, IsURl bool) (byt []byte, err error) {
	jsnbyt, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	byt, err = h.CallApyhub(method, url, bytes.NewReader(jsnbyt))
	if err != nil {
		return nil, err
	}
	if IsURl {
		var respJsn h.Response
		err = json.Unmarshal(byt, &respJsn)
		if err != nil {
			return nil, err
		}
		byt = []byte(respJsn.Data)
	}
	return
}
