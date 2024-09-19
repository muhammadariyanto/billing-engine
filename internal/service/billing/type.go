package billingService

import (
	"github.com/muhammadariyanto/billing-engine/internal/repository"
	"github.com/muhammadariyanto/billing-engine/internal/service"
)

type billingService struct {
	billingRepository repository.IBillingRepository
	loanRepository    repository.ILoanRepository
}

func New(
	billingRepository repository.IBillingRepository,
	loanRepository repository.ILoanRepository,
) service.IBillingService {
	return &billingService{
		billingRepository: billingRepository,
		loanRepository:    loanRepository,
	}
}
