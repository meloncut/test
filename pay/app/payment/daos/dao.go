package daos

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"test/pay/app/payment/conf"
)

var PayDB *Dao

type Dao struct {
	db *sqlx.DB
}

// New init mysql db
func New(c *conf.Config) (dao *Dao) {
	dsn := c.MySQLConfig["dsn"]
	println(dsn)
	sqlxDB,err := sqlx.Open("mysql",dsn)
	if err != nil {
		panic(err)
	}
	sqlxDB.SetMaxOpenConns(100)
	
	dao = &Dao{
		db:sqlxDB,
	}
	return
}

func (d *Dao) Close()  {
	err := d.db.Close()
	if err != nil {
		//log
	}
}


