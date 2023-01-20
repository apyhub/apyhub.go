package search

import (
	"bytes"
	"encoding/json"
	"fmt"

	h "github.com/apyhub/apyhub/helper"
)

func FuzzySearch(source string, target string, unicodeNormalize bool) (data []string, err error) {
	reqJsn := h.FuzzyRequest{Source: source, Target: target}
	byt, err := json.Marshal(reqJsn)
	if err != nil {
		return
	}
	url := fmt.Sprintf("%s?unicode-normalized=%t", h.FuzzySearch, unicodeNormalize)
	byt, err = h.CallApyhub("POST", url, bytes.NewReader(byt))
	if err != nil {
		return
	}

	respJsn := h.FuzzyResponse{}
	if err = json.Unmarshal(byt, &respJsn); err != nil {
		return
	}

	return respJsn.Data, nil
}
