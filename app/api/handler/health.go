// Actual logic for the APIs.
package handler

import (
	"net/http"
)

// @description Health check endpoint.
// @summary Health check.
// @tags health
// @produce plain
// @success 200 {string} string "OK"
// @router /v1/health [get]
// Health is a simple health check endpoint.
func Health(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("OK"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
