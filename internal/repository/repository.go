package repository

import (
	"context"
	billingModel "github.com/muhammadariyanto/billing-engine/internal/model/billing"
	customerModel "github.com/muhammadariyanto/billing-engine/internal/model/customer"
	loanModel "github.com/muhammadariyanto/billing-engine/internal/model/loan"
)

type ICustomerRepository interface {
	Insert(ctx context.Context, customer *customerModel.Customer) error
	FindByID(ctx context.Context, customerID string) (*customerModel.Customer, error)
}

type ILoanRepository interface {
	Insert(ctx context.Context, loan *loanModel.Loan) error
	FindByID(ctx context.Context, loanID string) (*loanModel.Loan, error)
}

type IBillingRepository interface {
	Insert(ctx context.Context, billing *billingModel.Billing) error
	FetchAllByLoanID(ctx context.Context, loanID string) ([]*billingModel.Billing, error)
	Update(ctx context.Context, billing *billingModel.Billing) error
}
