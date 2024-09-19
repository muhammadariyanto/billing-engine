package loanService

import (
	"github.com/muhammadariyanto/billing-engine/internal/repository"
	"github.com/muhammadariyanto/billing-engine/internal/service"
)

type loanService struct {
	loanRepository     repository.ILoanRepository
	customerRepository repository.ICustomerRepository
}

func New(
	loanRepository repository.ILoanRepository,
	customerRepository repository.ICustomerRepository,
) service.ILoanService {
	return &loanService{
		loanRepository:     loanRepository,
		customerRepository: customerRepository,
	}
}
