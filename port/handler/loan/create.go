package loanHandler

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/muhammadariyanto/billing-engine/constant"
	loanModel "github.com/muhammadariyanto/billing-engine/internal/model/loan"
	response "github.com/muhammadariyanto/billing-engine/util"
	"net/http"
)

func (h *loanHandler) Create(w http.ResponseWriter, r *http.Request) {
	var payload loanModel.ApplyLoanRequest
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		response.SendResponseError(w, http.StatusBadRequest, err)
		return
	}

	if err := h.validate.Struct(&payload); err != nil {
		response.SendResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	// Create loan entity
	loan := &loanModel.Loan{
		ID:           uuid.NewString(),
		CustomerID:   payload.CustomerID,
		Name:         payload.Name,
		Period:       payload.Period,
		Amount:       payload.Amount,
		InterestRate: payload.InterestRate,
		TotalAmount:  payload.Amount + (payload.Amount * payload.InterestRate),
		Status:       constant.LoanStatusInProgress,
	}
	if err := h.loanSvc.CreateLoan(r.Context(), loan); err != nil {
		response.SendResponseError(w, http.StatusInternalServerError, err)
		return
	}

	// Create schedule
	if err := h.billingSvc.CreateSchedule(r.Context(), loan.ID, payload.StartDate); err != nil {
		response.SendResponseError(w, http.StatusInternalServerError, err)
		return
	}

	// Fetch all schedule
	billings, err := h.billingSvc.FetchAllByLoanID(r.Context(), loan.ID)
	if err != nil {
		response.SendResponseError(w, http.StatusInternalServerError, err)
		return
	}

	response.SendResponseOK(w, map[string]any{
		"loan":     loan,
		"billings": billings,
	})
}
