package payment

import (
	"github.com/go-chi/chi"
	"test/pay/app/payment/https"
)

func Routes() *chi.Mux {
	router := chi.NewRouter()

	// /api/payment/notify
	router.Post("/notify", https.Notify)


	return router
}