package billingHandler

import (
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	response "github.com/muhammadariyanto/billing-engine/util"
	"net/http"
)

func (h *billingHandler) FetchAllByLoanID(w http.ResponseWriter, r *http.Request) {
	loanID := chi.URLParam(r, "id")
	if err := uuid.Validate(loanID); err != nil {
		response.SendResponseError(w, http.StatusBadRequest, errors.New("loan id is required"))
		return
	}

	// Fetch all schedule
	billings, err := h.billingSvc.FetchAllByLoanID(r.Context(), loanID)
	if err != nil {
		response.SendResponseError(w, http.StatusInternalServerError, err)
		return
	}

	response.SendResponseOK(w, billings)
}
