package service

import (
	"test/pay/app/payment/conf"
	"test/pay/app/payment/daos"
	"test/pay/app/payment/log"
)
func Load(c *conf.Config) {
	//Database init
	daos.PayDB = daos.New(c)
	//log init
	log.New()
}
