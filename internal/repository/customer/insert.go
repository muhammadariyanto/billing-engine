package customerRepository

import (
	"context"
	"errors"
	"github.com/google/uuid"
	customerModel "github.com/muhammadariyanto/billing-engine/internal/model/customer"
	"time"
)

func (r *customerRepository) Insert(ctx context.Context, customer *customerModel.Customer) error {
	if customer.ID == "" {
		customer.ID = uuid.NewString()
	}
	customer.CreatedAt = time.Now()

	if _, exists := r.data[customer.ID]; exists {
		return errors.New("customer already exists")
	}
	r.data[customer.ID] = customer
	return nil
}
