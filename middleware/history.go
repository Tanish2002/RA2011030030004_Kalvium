package middleware

import (
	"kalvium/models"
	"net/http"
	"time"
)

type Middleware struct {
	Model models.Model
}

func (m *Middleware) LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Capture request information
		logEntry := models.RequestLog{
			Method:    r.Method,
			URL:       r.URL.String(),
			Headers:   r.Header,
			Timestamp: time.Now(),
		}

		// Add log entry to the circular buffer
		m.Model.AddToRequestLog(logEntry)

		// Continue with the next handler
		next.ServeHTTP(w, r)
	})
}
