package loanRepository

import (
	loanModel "github.com/muhammadariyanto/billing-engine/internal/model/loan"
	"github.com/muhammadariyanto/billing-engine/internal/repository"
)

type loanRepository struct {
	data map[string]*loanModel.Loan
}

func New() repository.ILoanRepository {
	return &loanRepository{
		data: make(map[string]*loanModel.Loan),
	}
}
