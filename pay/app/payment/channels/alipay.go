package channels

import (
	"time"
)

const AlipayType = "alipay"

//implement PayReq
type AliReq struct {
	Content []byte //测试
}

func (*AliReq) GetPayResult() PayResult {
	return  PayResult{
		PayType:AlipayType,
		OrderCode:   "test-123123123",
		TotalPrice: 1000,
		PaidPrice:  800,
		NotifyTime:  time.Now()}
}

//func DecodeRequest(ctx []byte) gjson.Result {
//
//}
