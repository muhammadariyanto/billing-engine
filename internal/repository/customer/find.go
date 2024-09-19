package customerRepository

import (
	"context"
	"errors"
	customerModel "github.com/muhammadariyanto/billing-engine/internal/model/customer"
)

func (r *customerRepository) FindByID(ctx context.Context, customerID string) (*customerModel.Customer, error) {
	customer, ok := r.data[customerID]
	if !ok {
		return nil, errors.New("customer not found")
	}
	return customer, nil
}
