package response

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SendResponseOK(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	resp := Response{
		Message: "OK",
		Data:    data,
	}

	encodeJsonErr := json.NewEncoder(w).Encode(resp)
	if encodeJsonErr != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func SendResponseError(w http.ResponseWriter, statusCode int, errMessage error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	resp := Response{
		Message: errMessage.Error(),
	}

	encodeJsonErr := json.NewEncoder(w).Encode(resp)
	if encodeJsonErr != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
