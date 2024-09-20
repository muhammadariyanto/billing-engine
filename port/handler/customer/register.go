package customerHandler

import (
	"encoding/json"
	"github.com/google/uuid"
	customerModel "github.com/muhammadariyanto/billing-engine/internal/model/customer"
	response "github.com/muhammadariyanto/billing-engine/util"
	"net/http"
)

func (h *customerHandler) Register(w http.ResponseWriter, r *http.Request) {
	var payload customerModel.CreateCustomerRequest
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		response.SendResponseError(w, http.StatusBadRequest, err)
		return
	}

	if err := h.validate.Struct(&payload); err != nil {
		response.SendResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	customer := &customerModel.Customer{
		ID:   uuid.NewString(),
		Name: payload.Name,
	}

	if err := h.customerSvc.Register(r.Context(), customer); err != nil {
		response.SendResponseError(w, http.StatusInternalServerError, err)
		return
	}

	response.SendResponseOK(w, customer)
}
