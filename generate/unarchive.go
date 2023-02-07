package generate

import (
	"fmt"

	h "github.com/apyhub/apyhub.go/helper"
)

func Unarchive(conv h.Unzip) ([]string, error) {

	byt, err := prepareAPI(conv, "POST", h.UnarchiveUrl, true)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(byt))
	// return []string(byt), nil
	return nil, nil
}

func SecureUnarchive(conv h.Unzip) ([]string, error) {
	_, err := prepareAPI(conv, "POST", h.SecureUnarchiveUrl, true)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
