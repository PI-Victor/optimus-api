package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	v1alpha1 "github.com/cloudflavor/optimus-api/pkg/apis/v1alpha1"
	"github.com/cloudflavor/optimus-api/pkg/database"
	"github.com/cloudflavor/optimus-api/pkg/middleware"
)

func main() {
	logrus.SetLevel(logrus.TraceLevel)
	logrus.SetFormatter(
		&logrus.JSONFormatter{},
	)
	logrus.SetLevel(logrus.InfoLevel)

	if envLevel := os.Getenv("OPTIMUS_LOG_LEVEL"); envLevel != "" {
		logLevel, err := strconv.Atoi(envLevel)
		if err != nil {
			logrus.Warnf("failed to set custom log level: %s", err)
		}
		logrus.SetLevel(logrus.Level(logLevel))
	}

	certFile := os.Getenv("OPTIMUS_SSL_CERT_PATH")
	keyFile := os.Getenv("OPTIMUS_SSL_CERT_KEY_PATH")
	bindHost := os.Getenv("OPTIMUS_BIND_HOST")
	dbURI := os.Getenv("OPTIMUS_DB_URI")
	if bindHost == "" {
		bindHost = ":8000"
	}

	newMiddleWare := middleware.NewMiddleware()
	router := mux.NewRouter().StrictSlash(true)
	router.Use(
		newMiddleWare.Logging,
		newMiddleWare.ValidateContentType,
		newMiddleWare.WrapContentType,
	)

	for _, route := range v1alpha1.Routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.Handler)

		logrus.WithFields(
			logrus.Fields{
				"path": route.Pattern,
				"name": route.Name,
			},
		).Debug("route registered")
	}

	type optimus struct {
		database   *database.Database
		httpServer *http.Server
	}
	dbClient, err := database.NewDbConnection(dbURI)
	// TODO: make the API wait for the DB connection and not fail.
	if err != nil {
		logrus.Fatalf("An error occured while connecting to the database")
	}
	newApp := optimus{
		database: dbClient,
		httpServer: &http.Server{
			Handler:      router,
			Addr:         bindHost,
			WriteTimeout: 15 * time.Second,
			ReadTimeout:  15 * time.Second,
		},
	}

	go func() {
		logrus.Infof("starting server on %s", bindHost)
		logrus.Fatalf("server exited: %s", newApp.httpServer.ListenAndServeTLS(certFile, keyFile))
	}()
	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)

	<-c
	// NOTE: is this ok, should it be more or less?
	wait := time.Second * 3
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	go func() {
		err := newApp.httpServer.Shutdown(ctx)
		if err != nil {
			logrus.Fatalf("failed to close gracefully, %s", err)
		}
	}()

	<-ctx.Done()
	logrus.Info("shutting down server")
	os.Exit(0)
}
