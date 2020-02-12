package main

import (
	"flag"
	"github.com/go-chi/chi"
	"net/http"
	"test/pay/app/payment"
	"test/pay/app/payment/conf"
	"test/pay/app/payment/service"
)

const HttpServerPort = ":8080"

func main()  {
	ServerStart()
}

func ServerStart()  {
	flag.StringVar(&conf.ConfigPath, "conf", "config.json", "Set the config file path.")
	println(conf.ConfigPath)
	err := conf.Load()
	service.Load(conf.Conf)

	if err != nil {
		println("service config load failed")
		return
	}
	println("go server hello")
	err = http.ListenAndServe(HttpServerPort, RegisterRoute())

	if err != nil {
		println("http server start error")
	}
}

func RegisterRoute() *chi.Mux{
	r := chi.NewRouter()

	//heath check
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})

	r.Route("/api", func(r chi.Router) {
		//支付模块
		r.Mount("/payment", payment.Routes())
	})

	return r
}
