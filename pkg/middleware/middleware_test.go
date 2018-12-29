package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/gorilla/mux"

	v1alpha1 "github.com/cloudflavor/optimus-api/pkg/apis/v1alpha1"
)

func newRouter() *mux.Router {
	newMiddleWare := NewMiddleware()

	newRouter := mux.NewRouter().
		StrictSlash(true)

	newRouter.
		Use(
			newMiddleWare.Logging,
			newMiddleWare.ValidateContentType,
			newMiddleWare.WrapContentType,
		)

	for _, route := range v1alpha1.Routes {
		newRouter.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.Handler)
	}
	return newRouter
}

func TestMiddlewareFcuntionality(t *testing.T) {
	muxRouter := newRouter()

	rr := httptest.NewRecorder()
	newRequest, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("Failed to create new request on '/' endpoint %s ", err)
	}

	muxRouter.ServeHTTP(rr, newRequest)

	Convey("Index Endpoint should return 200", t, func() {
		So(rr.Code, ShouldEqual, http.StatusOK)
	})

	Convey("Response should have application/json content-type ", t, func() {
		So(rr.Header().Get("Content-Type"), ShouldEqual, "application/json")
	})

	Convey("POST to / should return 405", t, func() {
		rr := httptest.NewRecorder()
		newRequest, err := http.NewRequest("POST", "/", nil)
		if err != nil {
			t.Fatalf("Failed to create new request on '/' endpoint %s ", err)
		}

		muxRouter.ServeHTTP(rr, newRequest)
		So(rr.Code, ShouldEqual, http.StatusMethodNotAllowed)
	})

	Convey("Requests with wrong content-type should return 415", t, func() {
		rr := httptest.NewRecorder()
		newRequest, err := http.NewRequest("POST", "/v1alpha1/users", nil)
		if err != nil {
			t.Fatalf("Failed to create new request on '/' endpoint %s ", err)
		}
		newRequest.Header.Set("Content-Type", "application/html+text")
		muxRouter.ServeHTTP(rr, newRequest)
		So(rr.Code, ShouldEqual, http.StatusUnsupportedMediaType)
	})
}
