package loanService

import (
	"github.com/muhammadariyanto/billing-engine/internal/repository"
	"github.com/muhammadariyanto/billing-engine/internal/service"
)

type loanService struct {
	loanRepository     repository.ILoanRepository
	customerRepository repository.ICustomerRepository
	billingRepository  repository.IBillingRepository
}

func New(
	loanRepository repository.ILoanRepository,
	customerRepository repository.ICustomerRepository,
	billingRepository repository.IBillingRepository,
) service.ILoanService {
	return &loanService{
		loanRepository:     loanRepository,
		customerRepository: customerRepository,
		billingRepository:  billingRepository,
	}
}
