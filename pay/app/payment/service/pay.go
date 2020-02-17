package service

import (
	"test/pay/app/payment/daos"
	"test/pay/app/payment/log"
	"test/pay/app/payment/models"
)

//充值操作
func Recharge(orderCode string) error {
	order,err := daos.PayDB.GetOrderByCode(orderCode)
	if err != nil {
		return err
	}
	//如果已经发货,则什么都不做
	if order.Status != models.OrderPaid {
		switch order.Status {
		case models.OrderWait:
			{
				err = daos.PayDB.UpdateOrderStatus(orderCode, models.OrderWait, models.OrderPaid)
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
	wealthAccount,err := daos.PayDB.GetWealthAccountByUserIDAndWealthID(order.UserID,order.WealthID)
	if err != nil {
		return err
	}

	//开始发货

	//修改订单状态到已发货
	err = daos.PayDB.UpdateOrderStatus(orderCode, models.OrderPaid, models.OrderDeliver)
	if err != nil{
		return err
	}

	//增加财富数量
	err = daos.PayDB.UpdateWealthAccountIncrease(wealthAccount,order.Amount)
	if err != nil {
		log.ZLogger.Print("the wealth account increase failed, start to rollback the order status")
		rollBackErr := daos.PayDB.UpdateOrderStatus(orderCode, models.OrderDeliver, models.OrderPaid)
		if rollBackErr != nil {
			return  rollBackErr
		}
		return err
	}

	return nil
}
