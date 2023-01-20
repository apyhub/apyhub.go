package data

import (
	"fmt"

	h "github.com/apyhub/apyhub/helper"
)

func Country(countrycode string) (data h.Country, err error) {
	url := fmt.Sprintf("%s?country=%s", h.CountryInfo, countrycode)
	respJsn := new(h.CountryResponse)
	if err = prepareAPI(url, respJsn); err != nil {
		return data, err
	}

	return respJsn.Data, nil
}
