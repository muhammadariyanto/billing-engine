package billingService_test

import (
	"context"
	"github.com/google/uuid"
	billingModel "github.com/muhammadariyanto/billing-engine/internal/model/billing"
	. "github.com/muhammadariyanto/billing-engine/internal/service/billing"
	repositoryMock "github.com/muhammadariyanto/billing-engine/mocks/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestFetchAllByLoanID(t *testing.T) {
	billingRepo := repositoryMock.NewIBillingRepository(t)
	svc := New(billingRepo, nil)

	testCases := []struct {
		name      string
		wantErr   bool
		setupMock func()
	}{
		{
			name:    "SUCCESS",
			wantErr: false,
			setupMock: func() {
				billingRepo.On("FetchAllByLoanID",
					mock.Anything,
					mock.AnythingOfType("string"),
				).Once().Return([]*billingModel.Billing{
					{
						ID: uuid.NewString(),
					},
				}, nil)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.setupMock()

			resp, err := svc.FetchAllByLoanID(context.Background(), uuid.NewString())
			if tc.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, resp)
			}

			billingRepo.AssertExpectations(t)
		})
	}
}
