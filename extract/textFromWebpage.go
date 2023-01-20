package extract

import (
	"encoding/json"

	h "github.com/apyhub/apyhub/helper"
)

func TextFromWebpage(url string) (data string, err error) {
	url = h.TextFromWebpage + "?url=" + url
	byt, err := h.CallApyhub("GET", url, nil)
	if err != nil {
		return
	}

	var respJsn h.Response
	if err = json.Unmarshal(byt, &respJsn); err != nil {
		return
	}

	return respJsn.Data, nil
}
