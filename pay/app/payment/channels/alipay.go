package channels

import "time"

const AlipayType = "alipay"

//implement PayReq
type AliReq struct {
	Content []byte //测试
}

func (*AliReq) GetPayResult() PayResult {
	return  PayResult{
		PayType:AlipayType,
		OrderCode:   "test-order",
		TotalAmount: 1000,
		PaidAmount:  800,
		NotifyTime:  time.Now()}
}

//func DecodeRequest(ctx []byte)  {
//
//}
