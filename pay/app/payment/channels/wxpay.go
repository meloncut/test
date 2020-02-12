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
		OrderCode:   "test-order",
		TotalAmount: 1000,
		PaidAmount:  800,
		NotifyTime:  time.Now()}
}

