package models

import "time"

//财富类型WealthType
const (
	WealthTypeItems = "item" //道具
	WealthTypeCoin = "coin" //金币
	WealthTypeSilver = "silver" //银子
)

//财富账户
type WealthAccount struct {
	ID int64 `db:"id"`
	WealthID int64 `db:"wealth_id"`
	UserId int64 `db:"user_id"`
	Amount int64 `db:"amount"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type Wealth struct {
	ID int64 `db:"id"`
	WealthName string `db:"wealth_name"`
	WealthType string `db:"wealth_type"`
}

type WealthLog struct {
	ID int64
	Operator int64
	OperatorType string
	OrderCode string
	WealthType string //类型
	WealthName string //名称
	WealthAccountId int64
	Operate string
	BeforeChangeAmount string //修改前的值
	Variable string //变量
	ChangedAmount int64 //修改后的值
	CreatedAt time.Time
}
