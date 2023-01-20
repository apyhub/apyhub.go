package validation

import (
	h "github.com/apyhub/apyhub/helper"
)

func VatValidation(vat string) (bool, error) {
	vatNumber := h.VatRequest{Vat: vat}
	return PrepareAndCallApyhub(vatNumber, h.VatValidate)
}
