package billingService_test

import (
	"context"
	"errors"
	"github.com/google/uuid"
	loanModel "github.com/muhammadariyanto/billing-engine/internal/model/loan"
	. "github.com/muhammadariyanto/billing-engine/internal/service/billing"
	repositoryMock "github.com/muhammadariyanto/billing-engine/mocks/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func TestCreateSchedule(t *testing.T) {
	loanRepo := repositoryMock.NewILoanRepository(t)
	billingRepo := repositoryMock.NewIBillingRepository(t)
	svc := New(billingRepo, loanRepo)

	testCases := []struct {
		name      string
		wantErr   bool
		setupMock func()
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
			name:    "SUCCESS",
			wantErr: false,
			setupMock: func() {
				loanRepo.On("FindByID",
					mock.Anything,
					mock.AnythingOfType("string"),
				).Once().Return(&loanModel.Loan{
					Period: 7,
				}, nil)

				billingRepo.On("Insert",
					mock.Anything,
					mock.AnythingOfType("*billingModel.Billing"),
				).Times(7).Return(nil)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.setupMock()

			err := svc.CreateSchedule(context.Background(), uuid.NewString(), time.Now())
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
