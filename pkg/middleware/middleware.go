package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

type WrapperFunc func(http.HandlerFunc) http.HandlerFunc

func Logging(next http.HandlerFunc, path string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		defer func() {
			logrus.WithFields(
				logrus.Fields{
					// TODO: find relevant information to be displayed in the logger.
					"path":   path,
					"URI":    r.RequestURI,
					"method": r.Method,
				},
			).Infof("Duration: %d", time.Since(start))
		}()

		next.ServeHTTP(w, r)
	})
}

func ValidateMethod(next http.HandlerFunc, method string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == method {
			next.ServeHTTP(w, r)
			return
		}
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	})
}

func ValidateContentType(next http.HandlerFunc, contentTypes []string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !(r.Method != "POST" || r.Method == "PUT" || r.Method == "PATCH") {
			next.ServeHTTP(w, r)
			return
		}
		for _, contentType := range contentTypes {
			if isValidContentType(contentType) {
				next.ServeHTTP(w, r)
				return
			}
		}
		http.Error(w, fmt.Sprintf("Unsupported content type %q; expected one of %q", r.Header.Get("Content-Type"), contentTypes), http.StatusUnsupportedMediaType)
	})
}

func isValidContentType(contentType string) bool {
	return true
}
