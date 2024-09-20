package service

import (
	"context"
	billingModel "github.com/muhammadariyanto/billing-engine/internal/model/billing"
	customerModel "github.com/muhammadariyanto/billing-engine/internal/model/customer"
	loanModel "github.com/muhammadariyanto/billing-engine/internal/model/loan"
	"time"
)

type IBillingService interface {
	CreateSchedule(ctx context.Context, loanID string, startDate time.Time) error
	MakePayment(ctx context.Context, loanID string, paymentAmount float64, paymentDate time.Time) error
	FetchAllByLoanID(ctx context.Context, loanID string) ([]*billingModel.Billing, error)
}

type ILoanService interface {
	CreateLoan(ctx context.Context, loan *loanModel.Loan) error
	GetOutstanding(ctx context.Context, loanID string) (float64, error)
}

type ICustomerService interface {
	Register(ctx context.Context, customer *customerModel.Customer) error
	IsDelinquent(ctx context.Context, customerID string) (bool, error)
}
