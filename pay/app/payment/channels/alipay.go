package channels

import (
	"net/http"
	"time"
)


//->interface PayChannel
type AliChannel struct {
	request *http.Request
	responseWriter http.ResponseWriter
}

func (c *AliChannel) GetPayResult() PayResult {
	p := c.paramsAnalysis()
	//TODO
	return  PayResult{
		PayType:AlipayType,
		OrderCode:   "test-123123123",
		TotalPrice: 1000,
		PaidPrice:  800,
		NotifyTime:  time.Now()}
}

func (c *AliChannel) paramsAnalysis() map[string]string {
	var params map[string]string
	//..解析参数 验签
	return params
}

func (c *AliChannel) ResponsePaySuccess()  {
	//TODO
	c.responseWriter.WriteHeader(200)
	_,_ = c.responseWriter.Write([]byte("success"))
}

func (c *AliChannel) ResponsePayFail(message string)  {
	//TODO
	c.responseWriter.WriteHeader(400)
	_,_ = c.responseWriter.Write([]byte("fail"))
}

