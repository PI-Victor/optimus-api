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
	certFile := os.Getenv("CERT_PATH")
	keyFile := os.Getenv("KEY_PATH")
	bindHost := os.Getenv("BIND_HOST")
	if bindHost == "" {
		bindHost = ":8000"
	}

	router := mux.NewRouter().StrictSlash(true)
	validContentTypes := []string{"content-type: application/json"}

	for _, route := range v1alpha1.Routes {
		var handler http.HandlerFunc
		// TODO: write an aggregator with a proper type to avoid doing this manually
		handler = middleware.Logging(route.Handler, route.Pattern)
		handler = middleware.ValidateContentType(route.Handler, validContentTypes)
		handler = middleware.ValidateMethod(route.Handler, route.Method)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
		logrus.Debugf("registered route: %#v", route)
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
