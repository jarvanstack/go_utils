package wxu

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/jarvanstack/go_utils/logger"
)

type WxLoginResp struct {
	Openid     string `json:"openid"`
	SessionKey string `json:"session_key"`
	Unionid    string `json:"unionid"`
	Errcode    int    `json:"errcode"`
	Errmsg     string `json:"errmsg"`
}

//通过 code 调用微信登录接口
func WxLogin(appId, secret, code string) (*WxLoginResp, error) {
	uurl := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", appId, secret, code)
	r, err := http.Get(uurl)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	defer r.Body.Close()
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	resp := new(WxLoginResp)
	err = json.Unmarshal(b, resp)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	return resp, nil
}
