package loanHandler

import (
	"github.com/go-playground/validator/v10"
	"github.com/muhammadariyanto/billing-engine/internal/service"
	"github.com/muhammadariyanto/billing-engine/port/handler"
)

type loanHandler struct {
	validate *validator.Validate

	loanSvc    service.ILoanService
	billingSvc service.IBillingService
}

func New(
	loanSvc service.ILoanService,
	billingSvc service.IBillingService,
) handler.ILoanHandler {
	return &loanHandler{
		validate:   validator.New(),
		loanSvc:    loanSvc,
		billingSvc: billingSvc,
	}
}
