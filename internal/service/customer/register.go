package customerService

import (
	"context"
	customerModel "github.com/muhammadariyanto/billing-engine/internal/model/customer"
)

func (s *customerService) Register(ctx context.Context, customer *customerModel.Customer) error {
	return s.customerRepo.Insert(ctx, customer)
}
