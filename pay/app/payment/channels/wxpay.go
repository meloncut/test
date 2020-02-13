package channels

import (
"time"
)

const WxPayType = "wxpay"

//implement PayReq
type WxReq struct {
	Content []byte //测试
}

func (*WxReq) GetPayResult() PayResult {
	return  PayResult{
		PayType:WxPayType,
		OrderCode:   "test-123123123",
		TotalPrice: 1000,
		PaidPrice:  800,
		NotifyTime:  time.Now()}
}

