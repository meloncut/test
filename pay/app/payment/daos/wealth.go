package daos

import (
	"errors"
	"test/pay/app/payment/log"
	"test/pay/app/payment/models"
)

const (
	_wealthAccountGetSQL = "SELECT * FROM `wealth_accounts` WHERE `user_id` = ? AND `wealth_id` = ?"
	_wealthIncreaseSQL = "UPDATE `wealth_accounts` SET amount = amount+? WHERE id = ? AND amount = ?"
)

//充值操作
func (d *Dao) Recharge(orderCode string) error {

	order,err := d.GetOrderByCode(orderCode)
	if err != nil {
		return err
	}
	//如果已经发货,则什么都不做
	if order.Status != models.OrderPaid {
		switch order.Status {
			case models.OrderWait:
				{
					err = d.UpdateOrderPaid(orderCode)
					if err != nil {
						return err
					}
					break
				}
			case models.OrderDeliver:
				{
					log.ZLogger.Printf("order %s has delivered", orderCode)
					return nil
				}
			default:
				{
					log.ZLogger.Printf("order %s status is illegal", orderCode)
					return nil
				}
		}
	}

	//获取到用户的财富账户
	wealthAccount,err := d.GetWealthAccountByUserIDAndWealthID(order.UserID,order.WealthID)
	if err != nil {
		return err
	}

	//开始发货
	err = d.deliver(wealthAccount, order)

	if err != nil {
		return err
	}

	return nil
}

//获取对应的财富账户
func (d *Dao) GetWealthAccountByUserIDAndWealthID(userID int64, wealthID int64) (*models.WealthAccount, error){
	wealthAccount := &models.WealthAccount{}
	err := d.db.Get(wealthAccount,_wealthAccountGetSQL,userID,wealthID)
	if err != nil {
		return nil, err
	}
	return wealthAccount, nil
}

//发货数据库操作
func (d *Dao) deliver(wealthAccount *models.WealthAccount, order *models.Order) error {
	trans,err := d.db.Begin()
	if err != nil {
		return err
	}
	//更新订单状态为已发货
	result,err := trans.Exec(_orderStatusUpdateSQL,models.OrderDeliver,order.OrderCode,models.OrderPaid)
	if err != nil{
		_ = trans.Rollback()
		return err
	}
	effect,err := result.RowsAffected()
	if effect != 1 || err != nil {
		_ = trans.Rollback()
		return  errors.New("deliver failed when update the order")
	}

	//更新财富账户数据
	result,err = trans.Exec(_wealthIncreaseSQL,order.Amount,wealthAccount.ID,wealthAccount.Amount)
	if err != nil{
		_ = trans.Rollback()
		return err
	}
	effect,err = result.RowsAffected()
	if effect != 1 || err != nil {
		_ = trans.Rollback()
		return  errors.New("deliver failed when update the wealth account")
	}

	err = trans.Commit()
	if err != nil {
		return err
	}

	return nil
}

