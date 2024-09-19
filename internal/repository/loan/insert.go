package loanRepository

import (
	"context"
	"errors"
	"github.com/google/uuid"
	loanModel "github.com/muhammadariyanto/billing-engine/internal/model/loan"
	"time"
)

func (r *loanRepository) Insert(ctx context.Context, loan *loanModel.Loan) error {
	if loan.ID == "" {
		loan.ID = uuid.NewString()
	}
	loan.CreatedAt = time.Now()

	if _, exists := r.data[loan.ID]; exists {
		return errors.New("loan already exists")
	}
	r.data[loan.ID] = loan
	return nil
}
