package channels

import (
	"encoding/xml"
	"net/http"
	"time"
)

const WxPayType = "wxpay"

//->interface PayChannel
type WxChannel struct {
	request *http.Request
	responseWriter http.ResponseWriter
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
	//TODO
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
	c.responseWriter.WriteHeader(200)
	c.responseWriter.Header().Set("Content-type","application/xml")
	ctx,_ := xml.Marshal(msg)
	_,_ = c.responseWriter.Write(ctx)
}

func (c *WxChannel) ResponsePayFail(message string) {
	msg := ResultData{
		ReturnCode:CDATA{"FAIL"},
		ReturnMsg:CDATA{message},
	}

	c.responseWriter.WriteHeader(400)
	c.responseWriter.Header().Set("Content-type","application/xml")
	ctx,_ := xml.Marshal(msg)
	_,_ = c.responseWriter.Write(ctx)
}
