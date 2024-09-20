package billingHandler

import (
	"github.com/go-playground/validator/v10"
	"github.com/muhammadariyanto/billing-engine/internal/service"
	"github.com/muhammadariyanto/billing-engine/port/handler"
)

type billingHandler struct {
	validate *validator.Validate

	billingSvc service.IBillingService
}

func New(billingSvc service.IBillingService) handler.IBillingHandler {
	return &billingHandler{
		validate:   validator.New(),
		billingSvc: billingSvc,
	}
}
