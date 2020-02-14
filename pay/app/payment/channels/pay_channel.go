package channels

import (
	"errors"
	"io/ioutil"
	"net/http"
)

const (
	ContentJson = "application/json"
	ContentXML = "application/xml"
)

type PayChannel interface {
	GetPayResult() PayResult
	ResponsePaySuccess(w http.ResponseWriter)
}

//请求转换到相应的渠道
func Transfer(r *http.Request) (PayChannel, error) {
	content,err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil,errors.New("illegal request,request body err")
	}
	//辨别支付渠道,返回相应的request
	switch r.Header.Get("Content-Type") {
		case ContentXML:{
			req := &WxChannel{Content:content}
			return req,nil
		}
		//TODO 支付宝为Post请求
		//case ContentJson:{
		//	req := &AliChannel{Content:content}
		//	return req,nil
		//}
		default:{
			return nil,errors.New("illegal request, unknown payment channel")
		}
	}
}