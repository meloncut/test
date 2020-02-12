package service

import (
	"test/pay/app/payment/conf"
	"test/pay/app/payment/daos"
)

func Load(c *conf.Config) {
	daos.PayDB = daos.New(c)
}
