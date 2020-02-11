package models

import "time"

const AliPay  = "alipay"
const WxPay = "wxpay"

type PayLog struct {
	ID int64
	OrderCode string //业务订单号
	PayType string //支付类型
	ProductID int64
	ActualPaid int64
	UserID int64
	Status int8
	CreatedAt time.Time
	UpdatedAt time.Time
}
