package https

import (
	"fmt"
	"net/http"
	"time"
)

//var TrackId string

// Response 回复代理
type Response struct {
	Begin time.Time
	http.ResponseWriter
}

// OK 向 http.ResponseWriter 写入成功返回内容
func (r Response) OK() {
	resp := `{"code": 0, "message":"success", "data":""}`

	// 记录日志
	//log.ZLogger.Info().Str("sk_id", TrackId).Str("request end：response", resp).Msg("success")

	// 设置response
	r.Header().Set("X-Cost-Time", time.Now().Sub(r.Begin).String())
	r.Header().Set("Content-Type", "application/json")
	r.WriteHeader(http.StatusOK)
	r.Write([]byte(resp))

	return
}

// Error 向 w 中写入 错误信息
func (r Response) Error(code int, message string, err error) {
	resp := fmt.Sprintf(`{"code": %d, "message":"%s", "data":""}`, code, message)

	// 记录日志
	//log.ZLogger.Info().Str("sk_id", TrackId).Err(err).Str("request end：response", resp).Msg("failed")

	// 设置response
	r.Header().Set("X-Cost-Time", time.Now().Sub(r.Begin).String())
	r.Header().Set("Content-Type", "application/json")
	r.WriteHeader(http.StatusOK)
	r.Write([]byte(resp))

	return
}

// Error 向 w 中写入 错误信息
func (r Response) Fail(code int, message string, err error) {
	resp := fmt.Sprintf(`{"code": %d, "message":"%s", "data":""}`, code, message)
	// 设置response
	r.Header().Set("X-Cost-Time", time.Now().Sub(r.Begin).String())
	r.Header().Set("Content-Type", "application/json")
	r.WriteHeader(http.StatusBadRequest)
	r.Write([]byte(resp))

	return
}

