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
	XMLName    xml.Name `xml:"xml"`
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

func (*WxChannel) ResponsePaySuccess(w http.ResponseWriter)  {
	msg := ResultData{
		ReturnCode:CDATA{"SUCCESS"},
		ReturnMsg:CDATA{"OK"},
	}
	w.WriteHeader(200)
	w.Header().Set("Content-type","application/xml")
	ctx,_ := xml.Marshal(msg)
	_,_ = w.Write(ctx)
}

