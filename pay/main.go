package main

import (
	"flag"
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"test/pay/app/payment"
	"test/pay/app/payment/conf"
	"test/pay/app/payment/https"
	"test/pay/app/payment/service"
	"time"
)

const HttpServerPort = ":8080"

// -> interface http.Handler
type CustomHandler struct {
	r *chi.Mux
}
func NewCustomHandler() *CustomHandler {
	a := &CustomHandler{
		r:RegisterRoute(),
	}
	return a
}

func (agent *CustomHandler) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	agent.r.ServeHTTP(https.Response{Begin: time.Now(), ResponseWriter: w}, r)
}

func main()  {
	ServerStart()
}

func ServerStart()  {
	flag.StringVar(&conf.ConfigPath, "conf", "config.json", "Set the config file path.")
	println(conf.ConfigPath)
	err := conf.Load()
	service.Load(conf.Conf)

	println("go server hello")
	go func() {
		if err != nil {
			log.Fatalln("http listen failed", err)
			return
		}
		err = http.ListenAndServe(HttpServerPort, NewCustomHandler())
	}()

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
