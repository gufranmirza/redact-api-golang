package middlewares

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gufranmirza/redact-api-golang/src/models"
)

// Logging middleware logs the incoming requests
func Logging(logger *log.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			t1 := time.Now().UTC()
			defer func() {
				requestID, ok := r.Context().Value(models.HdrRequestID).(string)
				if !ok {
					requestID = "unknown"
				}
				logger.Printf("%s: %s  Method: %s URL: %s RemoteAddr: %s UserAgent: %s Latency: %v ", models.HdrRequestID, requestID, r.Method, r.URL.Path, r.RemoteAddr, r.UserAgent(), time.Since(t1))
			}()
			next.ServeHTTP(w, r)
		})
	}
}

// Tracing adds a TracingID to each requests
func Tracing(logger *log.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := r.Header.Get("X-Request-Id")
			if requestID == "" {
				requestID = fmt.Sprintf("%d", time.Now().UnixNano())
			}
			ctx := context.WithValue(r.Context(), models.HdrRequestID, requestID)
			w.Header().Set("X-Request-Id", requestID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
