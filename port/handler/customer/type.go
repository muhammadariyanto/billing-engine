package customerHandler

import (
	"github.com/go-playground/validator/v10"
	"github.com/muhammadariyanto/billing-engine/internal/service"
	"github.com/muhammadariyanto/billing-engine/port/handler"
)

type customerHandler struct {
	validate    *validator.Validate
	customerSvc service.ICustomerService
}

func New(customerSvc service.ICustomerService) handler.ICustomerHandler {
	return &customerHandler{
		validate:    validator.New(),
		customerSvc: customerSvc,
	}
}
