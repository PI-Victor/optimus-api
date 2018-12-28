package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

// wrappedRW convenience struct that aids middleware wrapping.
type wrappedRW struct {
	http.ResponseWriter
	httpResponseStatus int
}

// WriteHeader helps hijack the current http status code so that
// we can use for logging.
func (h *wrappedRW) WriteHeader(code int) {
	h.httpResponseStatus = code
	h.ResponseWriter.WriteHeader(code)
}

// NewMiddleware is a constructor for the middleware struct
func NewMiddleware() *Middleware {
	return &Middleware{}
}

// Middleware is a conveniece struct to converge all middleware
// functionality.
type Middleware struct{}

// Logging is the middleware that handles request response, path and method logging.
func (m *Middleware) Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wrappedRW := &wrappedRW{
			ResponseWriter:     w,
			httpResponseStatus: http.StatusOK,
		}
		start := time.Now()

		defer func() {
			logrus.WithFields(
				logrus.Fields{
					"method":   r.Method,
					"URI":      r.RequestURI,
					"response": wrappedRW.httpResponseStatus,
					"duration": time.Since(start),
				},
			).Info("received request")
		}()

		next.ServeHTTP(wrappedRW, r)
	})
}

// ValidateContentType will check that incoming requests match "application/json"
// content type.
func (m *Middleware) ValidateContentType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !(r.Method == "POST" || r.Method == "PUT" || r.Method == "PATCH") {
			next.ServeHTTP(w, r)
			return
		}
		if r.Header.Get("Content-Type") == "application/json" {
			next.ServeHTTP(w, r)
			return
		}
		http.Error(
			w,
			fmt.Sprintf(
				"Unsupported content type %q; expected one of %q",
				r.Header.Get("Content-Type"),
				"application/json",
			),
			http.StatusUnsupportedMediaType,
		)
	})
}

// WrapContentType sets the content-type for the response.
func (m *Middleware) WrapContentType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
