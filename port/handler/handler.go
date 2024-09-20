package handler

import (
	"net/http"
)

type ICustomerHandler interface {
	Register(w http.ResponseWriter, r *http.Request)
	IsDelinquent(w http.ResponseWriter, r *http.Request)
}

type ILoanHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
	GetOutstanding(w http.ResponseWriter, r *http.Request)
}

type IBillingHandler interface {
	MakePayment(w http.ResponseWriter, r *http.Request)
	FetchAllByLoanID(w http.ResponseWriter, r *http.Request)
}
