package generate

import (
	h "github.com/apyhub/apyhub/helper"
)

func ArchiveAsFile(conv h.Zip) ([]byte, error) {
	return prepareAPI(conv, "POST", h.ArchiveAsFile, false)
}

func ArchiveAsURL(conv h.Zip) (string, error) {
	byt, err := prepareAPI(conv, "POST", h.ArchiveAsURL, true)
	if err != nil {
		return "", err
	}
	return string(byt), nil
}

func SecureArchiveFile(conv h.ZipSecure) ([]byte, error) {
	return prepareAPI(conv, "POST", h.SecureArchiveFile, false)
}

func SecureArchiveAsURL(conv h.ZipSecure) (string, error) {
	byt, err := prepareAPI(conv, "POST", h.SecureArchiveURL, true)
	if err != nil {
		return "", err
	}
	return string(byt), nil
}
