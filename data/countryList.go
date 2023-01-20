package data

import (
	h "github.com/apyhub/apyhub/helper"
)

func CountryList() (data []h.Country, err error) {
	var respJsn h.CountryListResponse

	if err = prepareAPI(h.CountryList, &respJsn); err != nil {
		return data, err
	}
	return respJsn.Data, nil
}
