package middleware

import (
	"errors"
	"fmt"
	response "github.com/muhammadariyanto/billing-engine/util"
	"net/http"
	"sync"
)

var (
	idempotentStore = make(map[string]bool)
	storeMutex      sync.Mutex
)

// IdempotencyMiddleware ensures that requests with the same X-Request-Id are processed only once
func IdempotencyMiddleware() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Extract X-Request-Id from header
			requestID := r.Header.Get("X-Request-Id")
			if requestID == "" {
				response.SendResponseError(w, http.StatusBadRequest, errors.New("missing X-Request-Id header"))
				return
			}

			// Check if the request ID has already been processed
			storeMutex.Lock()
			if _, exists := idempotentStore[requestID]; exists {
				storeMutex.Unlock()

				// Return error (e.g., 409 Conflict) if the request was already processed
				response.SendResponseError(w, http.StatusConflict, fmt.Errorf("request with X-Request-Id '%s' has already been processed", requestID))
				return
			}

			// Mark the request ID as processed
			idempotentStore[requestID] = true
			storeMutex.Unlock()

			// Continue processing the request
			next.ServeHTTP(w, r)
		})
	}
}
