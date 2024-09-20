package port

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/muhammadariyanto/billing-engine/config"
	"github.com/muhammadariyanto/billing-engine/port/handler"
	customMiddleware "github.com/muhammadariyanto/billing-engine/port/middleware"
	"net/http"
	"time"
)

type RouterModule struct {
	Cfg *config.Config

	// API Route Modules
	CustomerHandler handler.ICustomerHandler
	LoanHandler     handler.ILoanHandler
	BillingHandler  handler.IBillingHandler
}

func NewRouter(module RouterModule) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(900 * time.Second))

	r.Mount("/debug", middleware.Profiler())

	r.Route("/api", func(r chi.Router) {
		r.Use(customMiddleware.InternalServiceMiddleware(module.Cfg))

		r.Route("/customer", func(r chi.Router) {
			r.Post("/register", module.CustomerHandler.Register)
			r.Get("/{id}/is-delinquent", module.CustomerHandler.IsDelinquent)
		})

		r.Route("/loan", func(r chi.Router) {
			r.Post("/apply", module.LoanHandler.Create)
			r.Get("/outstanding", module.LoanHandler.GetOutstanding)
		})

		r.Route("/billing", func(r chi.Router) {
			r.Post("/payment", module.BillingHandler.MakePayment)
		})
	})

	return r
}
