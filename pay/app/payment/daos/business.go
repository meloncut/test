package daos

import "test/pay/app/payment/models"

const (
	_getOrderSQL = "SELECT * FROM `orders` WHERE `order_code` = ?"
	_orderStatusUpdateSQL = "UPDATE `orders` SET `status` = ? `WHERE` `order_code` = ? and `status` = ?"
)
func (d *Dao) GetOrderByCode(orderCode string) (*models.Order, error) {
	order := &models.Order{}
	err := d.db.Get(order,_getOrderSQL,orderCode)

	if err != nil{
		return nil,err
	}
	return order,nil
}
