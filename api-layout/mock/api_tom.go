package mock

import (
	"encoding/json"
)

type TomService struct {

}

func NewTomService() *TomService{
	return &TomService{}
}

func (_ *TomService) post(params interface{}) (interface{}, error) {
	var res int
	switch v := params.(type) {
	case int:
		res = v + 3
	default:
		res = 3
	}
	resp := make(map[string]interface{})
	resp["id"] = "tom"
	resp["res"] = res
	jsonRes, _ := json.Marshal(resp)
	return string(jsonRes), nil
}