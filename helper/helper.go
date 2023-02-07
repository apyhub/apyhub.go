package helper

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"time"
)

var ContentType = func(contentType string) (string, string) { return "Content-Type", contentType }

var Token string

var Auth BasicAuth

// Create Form File
func CallApyhubFromFile(key string, url string, input io.Reader, filename string) ([]byte, error) {
	body, writer, err := PrepareBody(key, input, filename)
	if err != nil {
		return nil, err
	}
	writer.Close()
	return CallApyhub("POST", url, body, writer.FormDataContentType())
}

// Call Apyhub
func CallApyhubFromJson(url string, byt []byte) ([]byte, error) {
	return CallApyhub("POST", url, bytes.NewReader(byt))
}

func CallApyhub(method string, url string, body io.Reader, s ...string) ([]byte, error) {
	Client := http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	if Token == "" {
		return nil, ErrTokenNotSet
	}
	// Add the Token
	req.Header.Add("apy-token", Token)
	if len(s) == 0 {
		req.Header.Add(ContentType("application/json"))
	} else {
		req.Header.Add(ContentType(s[0]))
	}
	resp, err := Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		if resp.StatusCode == 500 {
			return errMessage(resp.Body)
		} else {
			return wrongMessage(resp.Body)
		}
	}

	byt, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return byt, nil
}

func PrepareBody(key string, input io.Reader, filename string) (*bytes.Buffer, *multipart.Writer, error) {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(key, filename)
	if err != nil {
		return nil, nil, err
	}

	if _, err := io.Copy(part, input); err != nil {
		return nil, nil, err
	}
	return body, writer, nil
}
