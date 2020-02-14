package channels

import (
	"net/http"
	"time"
)

const AlipayType = "alipay"

//implement PayReq
type AliChannel struct {
	Content []byte //测试
	http.ResponseWriter
}

func (*AliChannel) GetPayResult() PayResult {
	return  PayResult{
		PayType:AlipayType,
		OrderCode:   "test-123123123",
		TotalPrice: 1000,
		PaidPrice:  800,
		NotifyTime:  time.Now()}
}

func (c *AliChannel) ResponsePaySuccess()  {
	c.WriteHeader(200)
	_,_ = c.Write([]byte("success"))
}
