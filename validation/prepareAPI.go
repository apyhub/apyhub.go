package validation

import (
	"encoding/json"

	h "github.com/apyhub/apyhub.go/helper"
)

type Response struct {
	Data bool
}

func PrepareAndCallApyhub(ReqStruct interface{}, APILink string) (bool, error) {
	byt, err := json.Marshal(ReqStruct)
	if err != nil {
		return false, err
	}
	byt, err = h.CallApyhubFromJson(APILink, byt)
	if err != nil {
		return false, err
	}
	RespJsn := Response{}
	if err := json.Unmarshal(byt, &RespJsn); err != nil {
		return false, err
	}

	return RespJsn.Data, nil
}
