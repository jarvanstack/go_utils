package restful_util

import (
	"encoding/json"
	"log"
)

type RestfulResp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func NotFound(in interface{}) []byte {
	restfulResp := RestfulResp{
		Code: 404,
		Msg:  "NOT_FOUND",
		Data: in,
	}
	bytes, err := json.Marshal(restfulResp)
	if err != nil {
		log.Printf("string=%s\n", "解析json绑定错误")
	}
	return bytes

}
func Ok(in interface{}) []byte {
	restfulResp := RestfulResp{
		Code: 200,
		Msg:  "OK",
		Data: in,
	}
	bytes, err := json.Marshal(restfulResp)
	if err != nil {
		log.Printf("string=%s\n", "解析json绑定错误")
	}
	return bytes

}
func Fail(in interface{}) []byte {
	restfulResp := RestfulResp{
		Code: 400,
		Msg:  "BAD_REQUEST",
		Data: in,
	}
	bytes, err := json.Marshal(restfulResp)
	if err != nil {
		log.Printf("string=%s\n", "解析json绑定错误")
	}
	return bytes

}
