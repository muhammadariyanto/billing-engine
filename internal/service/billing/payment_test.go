package billingService_test

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/muhammadariyanto/billing-engine/constant"
	billingModel "github.com/muhammadariyanto/billing-engine/internal/model/billing"
	loanModel "github.com/muhammadariyanto/billing-engine/internal/model/loan"
	. "github.com/muhammadariyanto/billing-engine/internal/service/billing"
	repositoryMock "github.com/muhammadariyanto/billing-engine/mocks/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func TestMakePayment(t *testing.T) {
	loanRepo := repositoryMock.NewILoanRepository(t)
	billingRepo := repositoryMock.NewIBillingRepository(t)
	svc := New(billingRepo, loanRepo)

	testCases := []struct {
		name          string
		wantErr       bool
		setupMock     func()
		paymentAmount float64
	}{
		{
			name:    "ERROR: Find loan not found",
			wantErr: true,
			setupMock: func() {
				loanRepo.On("FindByID",
					mock.Anything,
					mock.AnythingOfType("string"),
				).Once().Return(nil, errors.New("loan not found"))
			},
		},
		{
			name:    "ERROR: Loan already completed",
			wantErr: true,
			setupMock: func() {
				loanRepo.On("FindByID",
					mock.Anything,
					mock.AnythingOfType("string"),
				).Once().Return(&loanModel.Loan{Status: constant.LoanStatusCompleted}, nil)
			},
		},
		{
			name:    "ERROR: fetch unpaid billing by loan",
			wantErr: true,
			setupMock: func() {
				loanRepo.On("FindByID",
					mock.Anything,
					mock.AnythingOfType("string"),
				).Return(&loanModel.Loan{Status: constant.LoanStatusInProgress, Period: 1}, nil)

				billingRepo.On("FetchUnpaidByLoanID",
					mock.Anything,
					mock.AnythingOfType("string"),
				).Once().Return(nil, errors.New("some error"))
			},
		},
		{
			name:    "ERROR: payment amount is not match",
			wantErr: true,
			setupMock: func() {
				billingRepo.On("FetchUnpaidByLoanID",
					mock.Anything,
					mock.AnythingOfType("string"),
				).Return([]*billingModel.Billing{{
					TotalAmount: 1_000_000.00,
					Sequence:    1,
				}}, nil)
			},
		},
		{
			name:    "ERROR: update billing",
			wantErr: true,
			setupMock: func() {
				billingRepo.On("Update",
					mock.Anything,
					mock.AnythingOfType("*billingModel.Billing"),
				).Once().Return(errors.New("error update billing"))
			},
			paymentAmount: 1_000_000.00,
		},
		{
			name:    "SUCCESS",
			wantErr: false,
			setupMock: func() {
				billingRepo.On("Update",
					mock.Anything,
					mock.AnythingOfType("*billingModel.Billing"),
				).Once().Return(nil)

				loanRepo.On("Update",
					mock.Anything,
					mock.AnythingOfType("*loanModel.Loan"),
				).Once().Return(nil)
			},
			paymentAmount: 1_000_000.00,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.setupMock()

			err := svc.MakePayment(context.Background(), uuid.NewString(), tc.paymentAmount, time.Now())
			if tc.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			billingRepo.AssertExpectations(t)
			loanRepo.AssertExpectations(t)
		})
	}
}
