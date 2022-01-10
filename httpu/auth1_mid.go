package httpu

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/dengjiawen8955/go_utils/constsu"
	"github.com/dengjiawen8955/go_utils/erru"
	"github.com/dengjiawen8955/go_utils/stringu"
)

func GetUserIdByContext(ctx context.Context) (userId int64, err error) {
	uif, ok := ctx.Value(constsu.TokenKeyUserId).(float64)
	if !ok {
		err = erru.NewError2(401, "token解析userId失败")
		return
	}
	userId = int64(uif)
	return
}
func GetUserIdStrByContext(ctx context.Context) (userId string, err error) {
	userId, ok := ctx.Value(constsu.TokenKeyUserId).(string)
	if !ok {
		err = erru.NewError2(401, "token解析userId失败")
		return
	}
	return
}

func Auth1Default(w http.ResponseWriter, r *http.Request, privateKey string) (req *http.Request, err error) {
	req, err = auth1(w, r, privateKey)
	if err != nil {
		//自动回写错误
		errResp := erru.NewError2(401, "401")
		switch e := err.(type) {
		case *erru.Err2:
			errResp.Msg = e.Msg
			errResp.Notice = e.Notice
		default:
		}
		w.WriteHeader(401)
		bs, err := json.Marshal(errResp)
		if err != nil {
			return nil, err
		}
		_, err = w.Write(bs)
		return nil, err
	}
	return
}
func auth1(w http.ResponseWriter, r *http.Request, privateKey string) (req *http.Request, err error) {
	req = r
	//1. 拿到 token
	t := r.Header.Get("authorization")
	if t == "" {
		err = erru.NewError2(401, "401拿不到authorization")
		return
	}
	//2.解析 token
	mc, err := stringu.TokenUnmarshal(t, privateKey)
	if err != nil {
		err = erru.NewError2(401, "401解析token失败")
		return
	}
	//3.拿到userId
	userId, ok := mc[constsu.TokenKeyUserId]
	if !ok {
		err = erru.NewError2(401, "401拿不到userId")
		return
	}
	//赋值 ctx
	ctx := context.WithValue(r.Context(), "userId", userId)
	req = r.WithContext(ctx)
	//4.拿到时间戳
	ts, ok := mc[constsu.TokenKeyTs].(float64)
	if !ok {
		err = erru.NewError2(401, "401时间戳不存在")
		return
	}
	now := time.Now().Unix()
	//过期
	if float64(now)-ts > float64(constsu.ExpSeconds) {
		err = erru.NewError2(401, "401过期")
		return
	}
	return
}
