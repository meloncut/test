package models

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
	CreatedAt []uint8 `db:"created_at"`
	UpdatedAt []uint8 `db:"updated_at"`
}

type Wealth struct {
	ID int64 `db:"id"`
	WealthName string `db:"wealth_name"`
	WealthType string `db:"wealth_type"`
	CreatedAt []uint8 `db:"created_at"`
}
