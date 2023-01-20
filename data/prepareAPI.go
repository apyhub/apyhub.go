package data

import (
	"encoding/json"

	h "github.com/apyhub/apyhub/helper"
)

func prepareAPI(url string, respJsn interface{}) error {
	byt, err := h.CallApyhub("GET", url, nil)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(byt, respJsn); err != nil {
		return err
	}
	return nil
}
