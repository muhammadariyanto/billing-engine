package customerHandler

import (
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	response "github.com/muhammadariyanto/billing-engine/util"
	"net/http"
)

func (h *customerHandler) IsDelinquent(w http.ResponseWriter, r *http.Request) {
	customerID := chi.URLParam(r, "id")
	if err := uuid.Validate(customerID); err != nil {
		response.SendResponseError(w, http.StatusBadRequest, errors.New("customer id is required"))
		return
	}

	isDelinquent, err := h.customerSvc.IsDelinquent(r.Context(), customerID)
	if err != nil {
		response.SendResponseError(w, http.StatusInternalServerError, err)
		return
	}

	response.SendResponseOK(w, map[string]any{
		"is_delinquent": isDelinquent,
	})
}
