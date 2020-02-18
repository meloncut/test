package https

import (
	"net/http"
	"test/pay/app/payment/channels"
	"test/pay/app/payment/service"
)
func Notify(w http.ResponseWriter, r *http.Request) {
	payChannel,err := channels.Transfer(r,w)

	if err != nil {
		w.(Response).Error(400,"illegal payment channel",err)
		return
	}

	payResult := payChannel.GetPayResult()

	err = service.Recharge(payResult.OrderCode)

	if err != nil {
		println(err.Error())
		payChannel.ResponsePayFail("recharge failed")
	}
	println(payResult.OrderCode)

	payChannel.ResponsePaySuccess()
}


