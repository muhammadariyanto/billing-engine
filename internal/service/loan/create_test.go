package loanService_test

import (
	"context"
	"errors"
	customerModel "github.com/muhammadariyanto/billing-engine/internal/model/customer"
	loanModel "github.com/muhammadariyanto/billing-engine/internal/model/loan"
	. "github.com/muhammadariyanto/billing-engine/internal/service/loan"
	repositoryMock "github.com/muhammadariyanto/billing-engine/mocks/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestCreateLoan(t *testing.T) {
	customerRepo := repositoryMock.NewICustomerRepository(t)
	loanRepo := repositoryMock.NewILoanRepository(t)
	svc := New(loanRepo, customerRepo, nil)

	testCases := []struct {
		name      string
		wantErr   bool
		setupMock func()
	}{
		{
			name:    "ERROR: Find customer error",
			wantErr: true,
			setupMock: func() {
				customerRepo.On("FindByID",
					mock.Anything,
					mock.AnythingOfType("string"),
				).Once().Return(nil, errors.New("some error"))
			},
		},
		{
			name:    "SUCCESS",
			wantErr: false,
			setupMock: func() {
				customerRepo.On("FindByID",
					mock.Anything,
					mock.AnythingOfType("string"),
				).Return(&customerModel.Customer{}, nil)

				loanRepo.On("Insert",
					mock.Anything,
					mock.AnythingOfType("*loanModel.Loan"),
				).Once().Return(nil)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.setupMock()

			err := svc.CreateLoan(context.Background(), &loanModel.Loan{})
			if tc.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			customerRepo.AssertExpectations(t)
			loanRepo.AssertExpectations(t)
		})
	}
}
