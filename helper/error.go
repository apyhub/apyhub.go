package helper

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

var (
	ErrTypeNotMatch        = errors.New("type not matched")
	ErrInvalidFormat       = errors.New("format of file submitted is invalid")
	ErrNotSameType         = errors.New("image and watermark is not same type")
	ErrUnableToGetCurrency = errors.New("unable to get currency")
	ErrTokenNotSet         = errors.New("authentication required")
	ErrNoArgs              = errors.New("url or file required")
	ErrMixedType           = errors.New("mixed type can't be allowed")
)

func errMessage(body io.ReadCloser) ([]byte, error) {
	var errjsn ErrJsn
	if err := json.NewDecoder(body).Decode(&errjsn); err != nil {
		return nil, err
	}
	return nil, fmt.Errorf("%v", errjsn.Error)
}

func wrongMessage(body io.ReadCloser) ([]byte, error) {
	var wrongjsn WrongJsn
	if err := json.NewDecoder(body).Decode(&wrongjsn); err != nil {
		return nil, err
	}
	if wrongjsn.Data != nil {
		return nil, fmt.Errorf("%v", wrongjsn.Data)
	}

	return nil, fmt.Errorf("%v", wrongjsn.Error)
}
