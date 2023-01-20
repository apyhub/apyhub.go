package convert

import (
	"encoding/json"

	h "github.com/apyhub/apyhub/helper"
)

func prepareAPI(obj interface{}, url string, IsUrl bool) (string, error) {
	jsnbyt, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}
	byt, err := h.CallApyhubFromJson(url, jsnbyt)
	if err != nil {
		return "", err
	}
	if IsUrl {

	}
	var res h.Response
	if err = json.Unmarshal(byt, &res); err != nil {
		return "", err
	}
	return res.Data, nil

}
