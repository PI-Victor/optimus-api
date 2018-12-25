package v1alpha1

import (
	"encoding/json"
	"net/http"

	"github.com/cloudflavor/optimus-api/pkg/apis"
)

func index(w http.ResponseWriter, r *http.Request) {
	availableAPIVersions := apis.APIVersions

	json.NewEncoder(w).Encode(availableAPIVersions)
}

func getUser(w http.ResponseWriter, r *http.Request) {

}

func findUserByID(w http.ResponseWriter, r *http.Request) {

}

func createNamespace(w http.ResponseWriter, r *http.Request) {

}

func getNamespaces(w http.ResponseWriter, r *http.Request) {

}
