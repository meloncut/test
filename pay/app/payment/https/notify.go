package https

import (
	"net/http"
	"test/pay/app/payment/channels"
	"test/pay/app/payment/service"
)
func Notify(w http.ResponseWriter, r *http.Request) {
	payChannel,err := channels.Transfer(r)

	if err != nil {
		w.(Response).Error(400,"illegal request params",err)
		return
	}

	payResult := payChannel.GetPayResult()

	err = service.Recharge(payResult.OrderCode)

	if err != nil {
		println(err.Error())
		w.(Response).Fail(400,"deliver failed",err)
	}
	println(payResult.OrderCode)

	payChannel.ResponsePaySuccess(w)
}


