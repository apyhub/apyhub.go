package data

import (
	"bytes"
	"encoding/json"

	h "github.com/apyhub/apyhub/helper"
)

func CurrencyConv(currency h.CurrencyConverter) (data float64, err error) {
	byt, err := json.Marshal(currency)
	if err != nil {
		return
	}
	byt, err = h.CallApyhub("POST", h.CurrencyConv, bytes.NewReader(byt))
	if err != nil {
		if err.Error() == "false" {
			err = h.ErrUnableToGetCurrency
			return
		}
		return
	}

	var respJsn h.CurrencyResponse
	if err = json.Unmarshal(byt, &respJsn); err != nil {
		return
	}
	return respJsn.Data, nil
}
