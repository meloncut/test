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

type PayReq interface {
	GetPayResult() PayResult
}

//请求转换到相应的渠道
func RequestTransfer(r *http.Request) (PayReq,error) {
	content,err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil,errors.New("illegal request,request body err")
	}
	//辨别支付渠道,返回相应的request
	switch r.Header.Get("Content-Type") {
		case ContentXML:{
			req := &AliReq{Content:content}
			return req,nil
		}
		case ContentJson:{
			req := &WxReq{Content:content}
			return req,nil
		}
		default:{
			return nil,errors.New("illegal request, unknown payment channel")
		}
	}
}