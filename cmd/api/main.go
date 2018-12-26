package main

import (
	//	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	v1alpha1 "github.com/cloudflavor/optimus-api/pkg/apis/v1alpha1"
	"github.com/cloudflavor/optimus-api/pkg/middleware"
)

func main() {
	logrus.SetLevel(logrus.TraceLevel)
	logrus.SetFormatter(
		&logrus.JSONFormatter{},
	)
	certFile := os.Getenv("OPTIMUS_SSL_CERT_PATH")
	keyFile := os.Getenv("OPTIMUS_SSL_CERT_KEY_PATH")
	bindHost := os.Getenv("OPTIMUS_BIND_HOST")
	if bindHost == "" {
		bindHost = ":8000"
	}

	router := mux.NewRouter().StrictSlash(true)
	validContentTypes := "application/json"

	for _, route := range v1alpha1.Routes {
		handler := middleware.WrapFunctionality(
			route.Handler,
			middleware.Logging(route.Pattern),
			middleware.WrapContentType(validContentTypes),
			middleware.ValidateMethod(route.Method),
			middleware.ValidateContentType(validContentTypes),
		)
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

		logrus.WithFields(
			logrus.Fields{
				"path": route.Pattern,
				"name": route.Name,
			},
		).Debug("registered route")
	}

	httpServer := &http.Server{
		Handler:      router,
		Addr:         bindHost,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logrus.Infof("Starting server on %s", bindHost)
	logrus.Fatalf(
		"Server exited: %s",
		httpServer.ListenAndServeTLS(certFile, keyFile),
	)
}
