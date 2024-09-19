package customerService

import (
	"github.com/muhammadariyanto/billing-engine/internal/repository"
	"github.com/muhammadariyanto/billing-engine/internal/service"
)

type customerService struct {
	customerRepository repository.ICustomerRepository
	loanRepository     repository.ILoanRepository
	billingRepository  repository.IBillingRepository
}

func New(
	customerRepository repository.ICustomerRepository,
	loanRepository repository.ILoanRepository,
	billingRepository repository.IBillingRepository,
) service.ICustomerService {
	return &customerService{
		customerRepository: customerRepository,
		loanRepository:     loanRepository,
		billingRepository:  billingRepository,
	}
}
