package httpu

import (
	"encoding/json"
	"net"
	"net/http"
	"strings"

	"github.com/dengjiawen8955/go_utils/erru"
)

func SpeedLimitDefault(w http.ResponseWriter, r *http.Request) (err error) {
	ip := ClientIP(r)
	//ip 为空 或者太快了就报错
	if ip == "" || !ltbDefault.Allow(ip) {
		speedLimitWriteDefault(w)
		err = erru.NewError2(400, "请求频繁")
		return
	}
	return

}
func speedLimitWriteDefault(w http.ResponseWriter) (err error) {
	w.WriteHeader(400)
	resp := erru.NewError2(400, "请求频繁")
	bs, err := json.Marshal(resp)
	w.Write(bs)
	return
}

var ltbDefault = NewLRUTokenBucket(100, 1, 5)

func ClientIP(r *http.Request) string {
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	ip := strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
	if ip != "" {
		return ip
	}

	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}

	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}

	return ""
}
