package middleware

import (
	"errors"
	"github.com/muhammadariyanto/billing-engine/config"
	response "github.com/muhammadariyanto/billing-engine/util"
	"net/http"
)

func InternalServiceMiddleware(config *config.Config) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// check internal api key
			internalApiKey := r.Header.Get("INTERNAL-API-KEY")
			if config.InternalApiKey != internalApiKey {
				response.SendResponseError(w, http.StatusUnauthorized, errors.New("unauthorized"))
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
