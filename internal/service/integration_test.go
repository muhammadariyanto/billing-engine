package service_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/muhammadariyanto/billing-engine/constant"
	customerModel "github.com/muhammadariyanto/billing-engine/internal/model/customer"
	loanModel "github.com/muhammadariyanto/billing-engine/internal/model/loan"
	billingRepository "github.com/muhammadariyanto/billing-engine/internal/repository/billing"
	customerRepository "github.com/muhammadariyanto/billing-engine/internal/repository/customer"
	loanRepository "github.com/muhammadariyanto/billing-engine/internal/repository/loan"
	billingService "github.com/muhammadariyanto/billing-engine/internal/service/billing"
	customerService "github.com/muhammadariyanto/billing-engine/internal/service/customer"
	loanService "github.com/muhammadariyanto/billing-engine/internal/service/loan"
)

func TestIntegration_CreateSchedule(t *testing.T) {
	ctx := context.Background()

	customerRepo := customerRepository.New()
	billingRepo := billingRepository.New()
	loanRepo := loanRepository.New()

	validCustomer := &customerModel.Customer{
		ID:   uuid.NewString(),
		Name: "John Doe",
	}
	_ = customerRepo.Insert(ctx, validCustomer)

	validLoanID := uuid.NewString()

	testCases := []struct {
		name      string
		wantErr   bool
		loanID    string
		setupMock func()
	}{
		{
			name:      "ERROR: Loan is not found",
			wantErr:   true,
			setupMock: func() {},
		},
		{
			name:    "SUCCESS: Successfully to create schedule",
			wantErr: false,
			loanID:  validLoanID,
			setupMock: func() {
				_ = loanRepo.Insert(ctx, &loanModel.Loan{
					ID:           validLoanID,
					CustomerID:   validCustomer.ID,
					Name:         "Loan 1",
					Period:       50,
					Amount:       5_000_000.00,
					InterestRate: 0.1,
					TotalAmount:  5_000_000.00 + (5_000_000.00 * 0.1),
					Status:       constant.LoanStatusInProgress,
				})
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.setupMock()

			svc := billingService.New(billingRepo, loanRepo)
			err := svc.CreateSchedule(ctx, tc.loanID, time.Now())

			billings, errFetch := billingRepo.FetchAllByLoanID(ctx, tc.loanID)
			for _, billing := range billings {
				fmt.Println("schedule: ", billing)
			}

			if tc.wantErr {
				assert.Error(t, err)
				assert.Error(t, errFetch)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, billings)
			}
		})
	}

}

func TestIntegration_MakePayment(t *testing.T) {
	ctx := context.Background()

	customerRepo := customerRepository.New()
	billingRepo := billingRepository.New()
	loanRepo := loanRepository.New()
	svc := billingService.New(billingRepo, loanRepo)

	validCustomer := &customerModel.Customer{
		ID:   uuid.NewString(),
		Name: "John Doe",
	}
	_ = customerRepo.Insert(ctx, validCustomer)

	validLoanID1 := uuid.NewString()
	validLoanID2 := uuid.NewString()
	validLoanID3 := uuid.NewString()

	testCases := []struct {
		name          string
		wantErr       bool
		loanID        string
		paymentAmount float64
		setupMock     func()
		expectedError string
	}{
		{
			name:      "ERROR: Loan is not found",
			wantErr:   true,
			setupMock: func() {},
		},
		{
			name:    "ERROR: Loan already completed",
			wantErr: true,
			loanID:  validLoanID1,
			setupMock: func() {
				_ = loanRepo.Insert(ctx, &loanModel.Loan{
					ID:           validLoanID1,
					CustomerID:   validCustomer.ID,
					Name:         "Loan 1",
					Period:       50,
					Amount:       5_000_000.00,
					InterestRate: 0.1,
					TotalAmount:  5_000_000.00 + (5_000_000.00 * 0.1),
					Status:       constant.LoanStatusCompleted,
				})
			},
		},
		{
			name:          "ERROR: Payment amount is not match",
			wantErr:       true,
			loanID:        validLoanID2,
			paymentAmount: 0.00,
			setupMock: func() {
				_ = loanRepo.Insert(ctx, &loanModel.Loan{
					ID:           validLoanID2,
					CustomerID:   validCustomer.ID,
					Name:         "Loan 2",
					Period:       50,
					Amount:       5_000_000.00,
					InterestRate: 0.1,
					TotalAmount:  5_000_000.00 + (5_000_000.00 * 0.1),
					Status:       constant.LoanStatusInProgress,
				})

				_ = svc.CreateSchedule(ctx, validLoanID2, time.Now())
			},
		},
		{
			name:          "SUCCESS: Successfully make payment for the first billing",
			wantErr:       false,
			loanID:        validLoanID3,
			paymentAmount: 110_000.00,
			setupMock: func() {
				_ = loanRepo.Insert(ctx, &loanModel.Loan{
					ID:           validLoanID3,
					CustomerID:   validCustomer.ID,
					Name:         "Loan 2",
					Period:       50,
					Amount:       5_000_000.00,
					InterestRate: 0.1,
					TotalAmount:  5_000_000.00 + (5_000_000.00 * 0.1),
					Status:       constant.LoanStatusInProgress,
				})

				_ = svc.CreateSchedule(ctx, validLoanID3, time.Now())
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.setupMock()
			err := svc.MakePayment(ctx, tc.loanID, tc.paymentAmount, time.Now())

			billings, _ := billingRepo.FetchAllByLoanID(ctx, tc.loanID)
			if tc.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, billings)

				firstBilling := billings[0]
				assert.NotEmpty(t, firstBilling.PaymentDate)
			}
		})
	}

}

func TestIntegration_GetOutstanding(t *testing.T) {
	ctx := context.Background()

	customerRepo := customerRepository.New()
	billingRepo := billingRepository.New()
	loanRepo := loanRepository.New()
	billingSvc := billingService.New(billingRepo, loanRepo)
	loanSvc := loanService.New(loanRepo, customerRepo, billingRepo)

	validCustomer := &customerModel.Customer{
		ID:   uuid.NewString(),
		Name: "John Doe",
	}
	_ = customerRepo.Insert(ctx, validCustomer)

	validLoanID1 := uuid.NewString()
	validLoanID2 := uuid.NewString()

	testCases := []struct {
		name                string
		wantErr             bool
		loanID              string
		setupMock           func()
		expectedOutstanding float64
	}{
		{
			name:      "ERROR: Outstanding balance is 0 because loan is not found",
			wantErr:   true,
			setupMock: func() {},
		},
		{
			name:    "SUCCESS: Outstanding balance is 0 because loan already completed",
			wantErr: false,
			loanID:  validLoanID1,
			setupMock: func() {
				_ = loanRepo.Insert(ctx, &loanModel.Loan{
					ID:           validLoanID1,
					CustomerID:   validCustomer.ID,
					Name:         "Loan 1",
					Period:       50,
					Amount:       5_000_000.00,
					InterestRate: 0.1,
					TotalAmount:  5_000_000.00 + (5_000_000.00 * 0.1),
					Status:       constant.LoanStatusCompleted,
				})
			},
			expectedOutstanding: 0.00,
		},
		{
			name:    "SUCCESS: Successfully get outstanding balance after make payment for the first billing",
			wantErr: false,
			loanID:  validLoanID2,
			setupMock: func() {
				_ = loanRepo.Insert(ctx, &loanModel.Loan{
					ID:           validLoanID2,
					CustomerID:   validCustomer.ID,
					Name:         "Loan 2",
					Period:       50,
					Amount:       5_000_000.00,
					InterestRate: 0.1,
					TotalAmount:  5_000_000.00 + (5_000_000.00 * 0.1),
					Status:       constant.LoanStatusInProgress,
				})

				_ = billingSvc.CreateSchedule(ctx, validLoanID2, time.Now())
			},
			expectedOutstanding: 5_500_000 - 110_000,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.setupMock()
			err := billingSvc.MakePayment(ctx, tc.loanID, 110_000.00, time.Now())
			outstandingBalance, err := loanSvc.GetOutstanding(ctx, tc.loanID)

			if tc.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)

				assert.Equal(t, tc.expectedOutstanding, outstandingBalance)
			}
		})
	}
}

func TestIntegration_IsDelinquent(t *testing.T) {
	ctx := context.Background()

	customerRepo := customerRepository.New()
	billingRepo := billingRepository.New()
	loanRepo := loanRepository.New()
	billingSvc := billingService.New(billingRepo, loanRepo)
	customerSvc := customerService.New(customerRepo, loanRepo, billingRepo)

	validCustomer := &customerModel.Customer{
		ID:   uuid.NewString(),
		Name: "John Doe",
	}
	_ = customerRepo.Insert(ctx, validCustomer)

	validCustomer2 := &customerModel.Customer{
		ID:   uuid.NewString(),
		Name: "Gabriella",
	}
	_ = customerRepo.Insert(ctx, validCustomer2)

	validLoanID := uuid.NewString()
	validLoanID2 := uuid.NewString()

	testCases := []struct {
		name         string
		customerID   string
		setupMock    func()
		isDelinquent bool
	}{
		{
			name:       "SUCCESS: Customer is delinquent",
			customerID: validCustomer.ID,
			setupMock: func() {
				_ = loanRepo.Insert(ctx, &loanModel.Loan{
					ID:           validLoanID,
					CustomerID:   validCustomer.ID,
					Name:         "Loan 1",
					Period:       50,
					Amount:       5_000_000.00,
					InterestRate: 0.1,
					TotalAmount:  5_000_000.00 + (5_000_000.00 * 0.1),
					Status:       constant.LoanStatusInProgress,
				})

				_ = billingSvc.CreateSchedule(ctx, validLoanID, time.Now().Add(-3*7*24*time.Hour)) // 3 weeks ago
			},
			isDelinquent: true,
		},
		{
			name:       "SUCCESS: Customer is not delinquent",
			customerID: validCustomer2.ID,
			setupMock: func() {
				_ = loanRepo.Insert(ctx, &loanModel.Loan{
					ID:           validLoanID2,
					CustomerID:   validCustomer2.ID,
					Name:         "Loan 2",
					Period:       50,
					Amount:       5_000_000.00,
					InterestRate: 0.1,
					TotalAmount:  5_000_000.00 + (5_000_000.00 * 0.1),
					Status:       constant.LoanStatusInProgress,
				})

				_ = billingSvc.CreateSchedule(ctx, validLoanID2, time.Now())
			},
			isDelinquent: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.setupMock()

			isDelinquent, _ := customerSvc.IsDelinquent(ctx, tc.customerID)
			assert.Equal(t, tc.isDelinquent, isDelinquent)
		})
	}
}
