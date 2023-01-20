package data

import (
	h "github.com/apyhub/apyhub/helper"
)

func CurrencyList() (data []h.Currency, err error) {
	var respJsn h.CurrencyList
	if err = prepareAPI(h.CurrenciesList, &respJsn); err != nil {
		return data, err
	}

	return respJsn.Data, nil
}
