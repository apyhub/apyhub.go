package generate

import (
	h "github.com/apyhub/apyhub.go/helper"
)

func IcalGeneratorAsFile(calendar h.CalenderEvent) ([]byte, error) {
	return prepareAPI(calendar, "POST", h.IcalAsFile, false)
}

func IcalGeneratorAsURL(calendar h.CalenderEvent) (string, error) {
	byt, err := prepareAPI(calendar, "POST", h.IcalAsURL, true)
	if err != nil {
		return "", err
	}
	return string(byt), nil
}
