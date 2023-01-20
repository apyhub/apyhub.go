package data

import (
	h "github.com/apyhub/apyhub.go/helper"
)

func Timezones() (data []h.Timezone, err error) {
	var respJsn h.TimezoneResponse
	if err = prepareAPI(h.Timezones, &respJsn); err != nil {
		return data, err
	}

	return respJsn.Data, nil
}
