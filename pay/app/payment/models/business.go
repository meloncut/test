package models

//Order Status
const (
	OrderWait = 1 //等待付款
	OrderPaid = 2 //已付款
	OrderDeliver = 3 //已发货
	OrderClose = 4 //订单关闭
	OrderWaitRefund = 5 //等待退款
	OrderRefund = 6 //已退款
	OrderDel = 7 //已删除
)

type Order struct {
	ID int64 `"db":id`
	OrderCode string `db:"order_code"` //订单号
	UserID int64 `db:"user_id"`
	OriginalPrice int64 `db:"original_price"` //原价
	ReceiptAmount int64 `db:"receipt_amount"` //实收
	Status 	int8 `db:"status"` //状态
	WealthID int64 `db:"wealth_id"` //财富ID
	Amount int64 `db:"amount"` //财富数量
	UnitPrice int64 `db:"unit_price"` //商品单价
	Quantity int64 `db:"quantity"` //商品数量
	CreatedAt []uint8 `db:"created_at"`
	UpdatedAt []uint8 `db:"updated_at"`
	PaidAt []uint8 `db:"paid_at"`
	CloseAt []uint8 `db:"closed_at"`
}

//如果存在财富组合如(金币+银子)，则需要抽象出商品
//type Product struct {
//	ID int64
//	Name string
//	UnitPrice int64
//	ProductType int8
//	WealthData string
//}