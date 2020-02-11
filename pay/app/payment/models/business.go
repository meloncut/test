package models

import "time"

//Order Status
const (
	OrderWait = 1 //等待付款
	OrderPaid = 2 //已付款
	OrderDeliver = 3 //已发货
	OrderClose = 4 //订单关闭
	OrderWaitRefund = 5 //等待退款
	OrderRefund = 6 //已退款
)

type order struct {
	ID int64
	OrderCode string //订单号
	OriginalPrice int64 //原价
	ReceiptAmount int64 //实收
	Status 	int8 //状态
	ProductId int64 //商品ID
	UnitPrice int64 //商品单价
	PQuantity int64 //商品数量
	CreatedAt time.Time
	PaidAt time.Time
	CloseAt time.Time
	DeletedAt int64
}



type product struct {
	ID int64
	Name string
	UnitPrice int64
	ProductType int8
}