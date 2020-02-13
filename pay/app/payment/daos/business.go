package daos

import (
	"errors"
	"fmt"
	"test/pay/app/payment/models"
)

const (
	_getOrderSQL = "SELECT * FROM `orders` WHERE `order_code` = ?"
	_orderStatusUpdateSQL = "UPDATE `orders` SET `status` = ? WHERE `order_code` = ? and `status` = ?"
)
func (d *Dao) GetOrderByCode(orderCode string) (*models.Order, error) {
	order := &models.Order{}
	err := d.db.Get(order,_getOrderSQL,orderCode)

	if err != nil{
		return nil,err
	}
	return order,nil
}

//更新订单状态为已付款
func (d *Dao) UpdateOrderPaid(orderCode string) error {
	result,err := d.db.Exec(_orderStatusUpdateSQL,models.OrderPaid,orderCode,models.OrderWait)
	if  err != nil{
		return errors.New("deliver failed when update the order status to paid")
	}
	effect,err := result.RowsAffected()
	if effect != 1 {
		fmt.Printf("%d effect",effect)
		return errors.New("deliver failed when update the order status to paid")
	}


	return nil
}
