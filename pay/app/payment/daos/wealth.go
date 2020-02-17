package daos

import (
	"errors"
	"fmt"
	"test/pay/app/payment/models"
)

const (
	_wealthAccountGetSQL = "SELECT * FROM `wealth_accounts` WHERE `user_id` = ? AND `wealth_id` = ?"
	_wealthIncreaseSQL = "UPDATE `wealth_accounts` SET amount = amount+? WHERE id = ? AND amount = ?"
)

//获取对应的财富账户
func (d *Dao) GetWealthAccountByUserIDAndWealthID(userID int64, wealthID int64) (*models.WealthAccount, error){
	wealthAccount := &models.WealthAccount{}
	err := d.db.Get(wealthAccount,_wealthAccountGetSQL,userID,wealthID)
	if err != nil {
		return nil, err
	}
	return wealthAccount, nil
}

//财富账户增加财富
func (d *Dao) UpdateWealthAccountIncrease(wealthAccount *models.WealthAccount, increaseAmount int64) error {
	result,err := d.db.Exec(_wealthIncreaseSQL,increaseAmount,wealthAccount.ID,wealthAccount.Amount)
	if err != nil{
		return err
	}
	effect,err := result.RowsAffected()
	if effect != 1 || err != nil {
		return  errors.New(fmt.Sprintf(
			"failed when update the wealth account (id %d) from %d to %d",
			wealthAccount.ID,
			wealthAccount.Amount,
			wealthAccount.Amount+increaseAmount,
			))
	}
	return nil
}

