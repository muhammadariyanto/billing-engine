package billingHandler

import (
	"encoding/json"
	billingModel "github.com/muhammadariyanto/billing-engine/internal/model/billing"
	response "github.com/muhammadariyanto/billing-engine/util"
	"net/http"
)

func (h *billingHandler) MakePayment(w http.ResponseWriter, r *http.Request) {
	var payload billingModel.MakePaymentRequest
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		response.SendResponseError(w, http.StatusBadRequest, err)
		return
	}

	if err := h.validate.Struct(&payload); err != nil {
		response.SendResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	if err := h.billingSvc.MakePayment(r.Context(), payload.LoanID, payload.PaymentAmount, payload.PaymentDate); err != nil {
		response.SendResponseError(w, http.StatusInternalServerError, err)
		return
	}

	response.SendResponseOK(w, payload)
}
