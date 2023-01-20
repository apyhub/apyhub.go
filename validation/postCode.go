package validation

import (
	h "github.com/apyhub/apyhub/helper"
)

const (
	IN CheckType = iota + 2
	UK
)

func PostCodeValidation(postcode string, checkType CheckType) (bool, error) {
	Postcode := h.PostRequest{Postcode: postcode}
	if checkType == IN {
		return PrepareAndCallApyhub(Postcode, h.INPostcodeValidate)
	} else if checkType == UK {
		return PrepareAndCallApyhub(Postcode, h.UKPostcodeValidate)
	}
	return false, h.ErrTypeNotMatch
}
