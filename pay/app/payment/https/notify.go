package https

import (
	"net/http"
	"test/pay/app/payment/channels"
	"test/pay/app/payment/daos"
)
func Notify(w http.ResponseWriter, r *http.Request) {
	payReq,err := channels.RequestTransfer(r)

	if err != nil {
		w.(Response).Error(400,"illegal request params",err)
		return
	}

	payResult := payReq.GetPayResult()

	err = daos.PayDB.Recharge(payResult.OrderCode)

	if err != nil {
		println(err.Error())
		w.(Response).Fail(400,"deliver failed",err)
	}


	println(payResult.OrderCode)

	_,_ = w.Write([]byte("success"))
}


