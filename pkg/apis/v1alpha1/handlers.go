package v1alpha1

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/cloudflavor/optimus-api/pkg/apis"
)

func index(w http.ResponseWriter, r *http.Request) {
	availableAPIVersions := apis.APIVersions

	json.NewEncoder(w).Encode(availableAPIVersions)
	w.WriteHeader(http.StatusOK)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	user := &User{}
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		logrus.Errorf("could not parse request body: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(b, user)
	if err != nil {
		logrus.Errorf("could not unmarshall request body: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func findUserByID(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func createNamespace(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func getNamespaceByID(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func createJob(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func getJobByID(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func callGitHubWebHook(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func getMetricsByType(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
