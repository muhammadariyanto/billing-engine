package customerService

import (
	"github.com/muhammadariyanto/billing-engine/internal/repository"
	"github.com/muhammadariyanto/billing-engine/internal/service"
)

type customerService struct {
	customerRepo repository.ICustomerRepository
}

func New(customerRepo repository.ICustomerRepository) service.ICustomerService {
	return &customerService{
		customerRepo: customerRepo,
	}
}
