package router

import (
	"cdvet/app/pkg/logger"
	"context"
	"net/http"
)

// loggerKey is the key used to store the logger in the context.
const LoggerKey Key = "logger"

// LoggingMiddleware is a middleware that logs requests and adds the logger to the context.
// It takes a logger as input and returns a new middleware.
func LoggingMiddleware(log logger.Logger) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Log the request.
			log.Infof("%s %s", r.Method, r.URL.Path)

			// Add the logger to the context.
			ctx := context.WithValue(r.Context(), LoggerKey, log)

			// Call the next handler.
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
