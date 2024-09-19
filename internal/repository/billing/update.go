package billingRepository

import (
	"context"
	"errors"
	billingModel "github.com/muhammadariyanto/billing-engine/internal/model/billing"
)

func (r *billingRepository) Update(ctx context.Context, billing *billingModel.Billing) error {
	if _, exists := r.data[billing.ID]; !exists {
		return errors.New("billing is not exists")
	}

	r.data[billing.ID] = billing
	return nil
}
