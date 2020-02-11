package main

import (
	"github.com/go-chi/chi"
	"meloncut/test/pay/app/payment"
	"net/http"
)

const HttpServerPort = ":8080"

func main()  {
	ServerStart()
}

func ServerStart()  {
	println("go server hello")

	err := http.ListenAndServe(HttpServerPort, RegisterRoute())

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
