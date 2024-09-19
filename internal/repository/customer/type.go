package customerRepository

import (
	customerModel "github.com/muhammadariyanto/billing-engine/internal/model/customer"
	"github.com/muhammadariyanto/billing-engine/internal/repository"
)

type customerRepository struct {
	data map[string]*customerModel.Customer
}

func New() repository.ICustomerRepository {
	return &customerRepository{
		data: make(map[string]*customerModel.Customer),
	}
}
