package channels

import (
	"errors"
	"github.com/go-chi/chi"
	"net/http"
)

const (
	AlipayType = "alipay"
	WxpayType = "wxpay"
)

type PayChannel interface {
	GetPayResult() PayResult
	ResponsePaySuccess()
	ResponsePayFail(message string)
}

//请求转换到相应的渠道
func Transfer(r *http.Request, w http.ResponseWriter) (PayChannel, error) {
	payType := chi.URLParam(r, "payType")
	//辨别支付渠道
	switch payType {
	case WxpayType:{
		channel := &WxChannel{request:r,responseWriter:w,}
		return channel,nil
	}
	case AlipayType:{
		channel := &AliChannel{request:r,responseWriter:w,}
		return channel,nil
	}
	default:{
		return nil,errors.New("illegal request, unknown payment channel")
	}
	}
}