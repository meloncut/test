package channels

import (
	"encoding/xml"
	"net/http"
	"time"
)

const WxPayType = "wxpay"

//implement PayReq
type WxChannel struct {
	Content []byte //测试
	http.ResponseWriter
}

type CDATA struct {
	Text string `xml:",cdata"`
}

type ResultData struct {
	ReturnCode CDATA `xml:"return_code"`
	ReturnMsg CDATA  `xml:"return_msg"`
}
func (*WxChannel) GetPayResult() PayResult {
	return  PayResult{
		PayType:WxPayType,
		OrderCode:   "test-123123123",
		TotalPrice: 1000,
		PaidPrice:  800,
		NotifyTime:  time.Now()}
}

func (c *WxChannel) ResponsePaySuccess()  {
	msg := ResultData{
		ReturnCode:CDATA{"SUCCESS"},
		ReturnMsg:CDATA{"OK"},
	}
	c.WriteHeader(200)
	c.Header().Set("Content-type","application/xml")
	ctx,_ := xml.Marshal(msg)
	_,_ = c.Write(ctx)
}

