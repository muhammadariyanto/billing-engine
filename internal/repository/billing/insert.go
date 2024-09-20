package billingRepository

import (
	"context"
	"errors"
	"github.com/google/uuid"
	billingModel "github.com/muhammadariyanto/billing-engine/internal/model/billing"
	"time"
)

func (r *billingRepository) Insert(ctx context.Context, billing *billingModel.Billing) error {
	r.mu.Lock()         // Lock before writing
	defer r.mu.Unlock() // Unlock after writing

	if billing.ID == "" {
		billing.ID = uuid.NewString()
	}
	billing.CreatedAt = time.Now()

	if _, exists := r.data[billing.ID]; exists {
		return errors.New("billing already exists")
	}
	r.data[billing.ID] = billing
	return nil
}
