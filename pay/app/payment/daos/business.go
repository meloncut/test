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

//更新订单状态 从某状态到某状态
func (d *Dao) UpdateOrderStatus(orderCode string, from int, to int) error  {
	result,err := d.db.Exec(_orderStatusUpdateSQL,to,orderCode,from)
	if  err != nil{
		return errors.New(fmt.Sprintf("failed when update the order status from %d to %d",from,to))
	}
	effect,err := result.RowsAffected()
	if effect != 1 {
		fmt.Printf("%d effect",effect)
		return errors.New(fmt.Sprintf("failed when update the order status from %d to %d",from,to))
	}

	return nil
}
