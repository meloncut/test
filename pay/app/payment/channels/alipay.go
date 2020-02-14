package channels

import (
	"net/http"
	"time"
)

const AlipayType = "alipay"

//implement PayReq
type AliChannel struct {
	Content []byte //测试
}

func (*AliChannel) GetPayResult() PayResult {
	//TODO
	return  PayResult{
		PayType:AlipayType,
		OrderCode:   "test-123123123",
		TotalPrice: 1000,
		PaidPrice:  800,
		NotifyTime:  time.Now()}
}

func (*AliChannel) ResponsePaySuccess(w http.ResponseWriter)  {
	//TODO
	w.WriteHeader(200)
	_,_ = w.Write([]byte("success"))
}
