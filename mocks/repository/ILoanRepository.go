// Code generated by mockery v2.43.0. DO NOT EDIT.

package mocks

import (
	context "context"

	loanModel "github.com/muhammadariyanto/billing-engine/internal/model/loan"
	mock "github.com/stretchr/testify/mock"
)

// ILoanRepository is an autogenerated mock type for the ILoanRepository type
type ILoanRepository struct {
	mock.Mock
}

// FetchUncompletedByCustomerID provides a mock function with given fields: ctx, customerID
func (_m *ILoanRepository) FetchUncompletedByCustomerID(ctx context.Context, customerID string) ([]*loanModel.Loan, error) {
	ret := _m.Called(ctx, customerID)

	if len(ret) == 0 {
		panic("no return value specified for FetchUncompletedByCustomerID")
	}

	var r0 []*loanModel.Loan
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]*loanModel.Loan, error)); ok {
		return rf(ctx, customerID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []*loanModel.Loan); ok {
		r0 = rf(ctx, customerID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*loanModel.Loan)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, customerID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByID provides a mock function with given fields: ctx, loanID
func (_m *ILoanRepository) FindByID(ctx context.Context, loanID string) (*loanModel.Loan, error) {
	ret := _m.Called(ctx, loanID)

	if len(ret) == 0 {
		panic("no return value specified for FindByID")
	}

	var r0 *loanModel.Loan
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*loanModel.Loan, error)); ok {
		return rf(ctx, loanID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *loanModel.Loan); ok {
		r0 = rf(ctx, loanID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*loanModel.Loan)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, loanID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Insert provides a mock function with given fields: ctx, loan
func (_m *ILoanRepository) Insert(ctx context.Context, loan *loanModel.Loan) error {
	ret := _m.Called(ctx, loan)

	if len(ret) == 0 {
		panic("no return value specified for Insert")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *loanModel.Loan) error); ok {
		r0 = rf(ctx, loan)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: ctx, loan
func (_m *ILoanRepository) Update(ctx context.Context, loan *loanModel.Loan) error {
	ret := _m.Called(ctx, loan)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *loanModel.Loan) error); ok {
		r0 = rf(ctx, loan)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewILoanRepository creates a new instance of ILoanRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewILoanRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *ILoanRepository {
	mock := &ILoanRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
