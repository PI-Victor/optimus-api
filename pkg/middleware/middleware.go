package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

type funcWrapper func(http.HandlerFunc) http.HandlerFunc

func Logging(path string) funcWrapper {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			defer func() {
				logrus.WithFields(
					logrus.Fields{
						"path":   path,
						"URI":    r.RequestURI,
						"Method": r.Method,
					},
				).Infof("Duration: %d", time.Since(start))
			}()

			next(w, r)
		}
	}
}

func ValidateMethod(method string) funcWrapper {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			if r.Method == method {
				next(w, r)
				return
			}
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		}
	}
}

func ValidateContentType(contentTypes ...string) funcWrapper {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			if !(r.Method != "POST" || r.Method == "PUT" || r.Method == "PATCH") {
				next(w, r)
				return
			}
			for _, contentType := range contentTypes {
				if isValidContentType(r.Header.Get("Content-Type"), contentType) {
					next(w, r)
					return
				}
			}
			http.Error(w, fmt.Sprintf("Unsupported content type %q; expected one of %q", r.Header.Get("Content-Type"), contentTypes), http.StatusUnsupportedMediaType)
		}
	}
}

func WrapContentType(contentType string) funcWrapper {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", contentType)
			next(w, r)
		}
	}
}

func isValidContentType(currentContentType string, contentTypes ...string) bool {
	for _, contentType := range contentTypes {
		if currentContentType == contentType {
			return true
		}
		continue
	}
	return false
}

// WrapFunctionality will add all the necessary middleware to the route's handler.
func WrapFunctionality(handlerFunc http.HandlerFunc, wrappers ...funcWrapper) http.HandlerFunc {
	for _, wrapper := range wrappers {
		handlerFunc = wrapper(handlerFunc)
	}
	return handlerFunc
}
