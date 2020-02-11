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
	ID int64
	WealthName string
	WealthType string
	UserId int64
	Amount int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type WealthLog struct {
	ID int64
	Operator int64
	OperatorType string
	OrderCode string
	WealthType string //类型
	WealthName string //名称
	WealthId int64
	Operate string
	BeforeChangeAmount string //修改前的值
	Variable string //变量
	ChangedAmount int64 //修改后的值
	CreatedAt time.Time
}
