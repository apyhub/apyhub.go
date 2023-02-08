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
