package generate

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"

	h "github.com/apyhub/apyhub.go/helper"
)

func ArchiveAsFile(conv []interface{}) ([]byte, error) {
	urls, files, err := getFileOrUrl(conv)
	if err != nil {
		return nil, err
	}

	if urls != nil {
		return prepareAPI(h.Zip{Urls: urls}, "POST", h.ArchiveURLToFile, false)
	}
	if files != nil {
		return callApyhubZip(files, h.ArchiveFileToFile, "")
	}
	return nil, h.ErrTypeNotMatch
}

func ArchiveAsURL(conv []interface{}) (string, error) {
	urls, files, err := getFileOrUrl(conv)
	if err != nil {
		return "", err
	}

	if urls != nil {
		byt, err := prepareAPI(h.Zip{Urls: urls}, "POST", h.ArchiveURLToURL, true)
		if err != nil {
			return "", err
		}
		return string(byt), nil
	}
	if files != nil {
		byt, err := callApyhubZip(files, h.ArchiveFileToURL, "")
		if err != nil {
			return "", err
		}
		url := h.Response{}
		if err := json.Unmarshal(byt, &url); err != nil {
			return "", err
		}
		return url.Data, nil
	}

	return "", h.ErrTypeNotMatch
}

func SecureArchiveFile(password string, conv []interface{}) ([]byte, error) {
	urls, files, err := getFileOrUrl(conv)
	if err != nil {
		return nil, err
	}

	if urls != nil {
		return prepareAPI(h.ZipSecure{Urls: urls, Password: password}, "POST", h.SecureArchiveURLToFile, false)
	}
	if files != nil {
		return callApyhubZip(files, h.SecureArchiveFileToFile, password)
	}

	return nil, h.ErrTypeNotMatch
}

func SecureArchiveAsURL(password string, conv []interface{}) (string, error) {
	urls, files, err := getFileOrUrl(conv)
	if err != nil {
		return "", err
	}

	if urls != nil {
		byt, err := prepareAPI(h.ZipSecure{Urls: urls, Password: password}, "POST", h.SecureArchiveURLToURL, true)
		if err != nil {
			return "", err
		}
		return string(byt), nil
	}
	if files != nil {
		byt, err := callApyhubZip(files, h.SecureArchiveFileToURL, password)
		if err != nil {
			return "", err
		}
		url := h.Response{}
		if err := json.Unmarshal(byt, &url); err != nil {
			return "", err
		}
		return url.Data, nil
	}

	return "", h.ErrTypeNotMatch

}

func UnArchiveAsURL(conv interface{}) ([]string, error) {
	switch input := conv.(type) {
	case string:
		return prepareAPIUnZip(h.UnZip{Url: input}, "POST", h.UnArchiveAsURLToURL)

	case io.Reader:
		var respUnZipJsn h.UnZipResponse
		fileByt, err := io.ReadAll(input)
		if err != nil {
			return nil, err
		}
		byt, err := h.CallApyhubFromFile("file", h.UnArchiveFileToURL, bytes.NewReader(fileByt), "sample"+h.GetFileExt(fileByt))
		if err != nil {
			return nil, err
		}
		if err = json.Unmarshal(byt, &respUnZipJsn); err != nil {
			return nil, err
		}
		return respUnZipJsn.Data, nil
	default:
		return nil, h.ErrTypeNotMatch

	}
}

func SecureUnArchiveAsURL(password string, conv interface{}) ([]string, error) {
	switch input := conv.(type) {
	case string:
		SecureZip := h.SecureUnZip{Url: input, Password: password}
		return prepareAPIUnZip(SecureZip, "POST", h.SecureUnArchiveURLToURL)
	case io.Reader:
		var respUnZipJsn h.UnZipResponse

		fileByt, err := io.ReadAll(input)
		if err != nil {
			return nil, err
		}

		byt, err := callApyhubSecure(h.SecureUnArchiveFileToURL, bytes.NewReader(fileByt), "sample"+h.GetFileExt(fileByt), password)
		if err != nil {
			return nil, err
		}
		if err := json.Unmarshal(byt, &respUnZipJsn); err != nil {
			return nil, err
		}
		return respUnZipJsn.Data, nil
	default:
		return nil, h.ErrTypeNotMatch
	}
}

func callApyhubSecure(url string, input *bytes.Reader, filename string, password string) ([]byte, error) {
	body, writer, err := h.PrepareBody("file", input, filename)
	if err != nil {
		return nil, err
	}

	if err = writer.WriteField("password", password); err != nil {
		return nil, err
	}
	writer.Close()
	return h.CallApyhub("POST", url, body, writer.FormDataContentType())
}

func callApyhubZip(input []io.Reader, url string, password string) ([]byte, error) {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	for _, file := range input {

		fileByt, err := io.ReadAll(file)
		if err != nil {
			return nil, err
		}
		part, err := writer.CreateFormFile("files", "sample"+h.GetFileExt(fileByt))
		if err != nil {
			return nil, err
		}
		if _, err := io.Copy(part, bytes.NewReader(fileByt)); err != nil {
			return nil, err
		}
		if password != "" {
			if err = writer.WriteField("password", password); err != nil {
				return nil, err
			}
		}
	}
	writer.Close()

	return h.CallApyhub("POST", url, body, writer.FormDataContentType())
}

func getFileOrUrl(conv []interface{}) ([]string, []io.Reader, error) {
	urls := []string(nil)
	files := []io.Reader(nil)
	if len(conv) == 0 {
		return nil, nil, h.ErrNoArgs
	}
	for _, v := range conv {
		switch input := v.(type) {
		case string:
			if files != nil {
				break
			}
			urls = append(urls, input)
		case []string:
			if files != nil {
				break
			}
			urls = append(urls, input...)
		case io.Reader:
			if urls != nil {
				break
			}
			files = append(files, input)
		case []io.Reader:
			if urls != nil {
				break
			}
			files = append(files, input...)
		default:
			return nil, nil, h.ErrTypeNotMatch
		}
	}
	return urls, files, nil
}
