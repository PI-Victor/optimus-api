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
	httpResponseCode := http.StatusOK

	user := &User{}
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logrus.Errorf("Couldn't parse request body: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(b, user)
	if err != nil {
		logrus.Errorf("Couldn't unmarshall request body: %s", err)
		httpResponseCode = http.StatusInternalServerError
		w.WriteHeader(httpResponseCode)
		return
	}

	w.WriteHeader(httpResponseCode)
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
