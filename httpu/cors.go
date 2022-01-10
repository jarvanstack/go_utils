package httpu

import "net/http"

func CORSAll(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*") //允许访问所有域
}
