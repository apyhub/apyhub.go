package validation

import (
	h "github.com/apyhub/apyhub.go/helper"
)

type CheckType int

const (
	DNS CheckType = iota
	Acadamic
)

func EmailValidation(email string, checkType CheckType) (bool, error) {
	Email := h.EmailRequest{Email: email}
	if checkType == DNS {
		return PrepareAndCallApyhub(Email, h.EmailDNS)
	} else if checkType == Acadamic {
		return PrepareAndCallApyhub(Email, h.EmailAcadamic)
	}
	return false, h.ErrTypeNotMatch
}
