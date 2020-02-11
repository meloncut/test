package https

import (
	//"meloncut/test/pay/app/payment/channels"
	"meloncut/test/pay/app/payment/channels"
	"net/http"
)

type NotifyReq interface {

}

func Notify(w http.ResponseWriter, r *http.Request) {
	payReq,err := channels.RequestTransfer(r)

	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
	}

	payResult := payReq.GetPayResult()

	println(payResult.OrderCode)
	w.Write([]byte("notify"))
}


