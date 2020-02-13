package daos

const (
	_wealthGetByNameSQL = "SELECT * FROM wealth WHERE `user_id` = ? AND `name` = ?"
)
func (d *Dao) GetWealthByUserIdAndName(name string, userID int64)  {
	result,err := d.db.Exec(_wealthGetByNameSQL,name,userID)
}
