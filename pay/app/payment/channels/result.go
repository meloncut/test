package channels

import "time"

type PayResult struct {
	PayType string //支付类型
	OrderCode string //我方订单号
	TotalAmount int64 //订单总价
	PaidAmount int64  //实收金额
	NotifyTime time.Time //通知时间
	//.....
}
