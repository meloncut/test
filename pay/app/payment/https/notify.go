package https

import (
	"net/http"
	"test/pay/app/payment/channels"
	"test/pay/app/payment/daos"
)

type NotifyReq interface {

}

func Notify(w http.ResponseWriter, r *http.Request) {
	payReq,err := channels.RequestTransfer(r)

	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}

	err = daos.PayDB.Recharge(payReq.GetPayResult().OrderCode)

	if err != nil {
		w.(Response).Fail(400,"deliver failed",err)
	}

	payResult := payReq.GetPayResult()

	println(payResult.OrderCode)

	w.(Response).OK()
}


