package main

import (
	"flag"
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"test/pay/app/payment"
	"test/pay/app/payment/conf"
	"test/pay/app/payment/daos"
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

	server := &http.Server{
		Addr:    HttpServerPort,
		Handler: NewCustomHandler(),
	}

	go func() {
		err = server.ListenAndServe()
		if err != nil {
			log.Fatalln("http listen failed", err)
			return
		}
	}()

	// 信号量监听
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGHUP)
	go func() {
		log.Println("Signaler running")
		sig := <-sigs

		log.Println("Receive a signal：" + sig.String())

		// 退出http
		_ = server.Close()

		// 释放config数据库资源
		daos.PayDB.Close()

		done <- true
	}()
	time.Sleep(10 * time.Millisecond)
	log.Println("service ready")
	<-done
	log.Println("service exit")
}

func RegisterRoute() *chi.Mux{
	r := chi.NewRouter()

	//heath check
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_,_ = w.Write([]byte("trade"))
	})

	r.Route("/api", func(r chi.Router) {
		//支付模块
		r.Mount("/payment", payment.Routes())
	})

	return r
}
