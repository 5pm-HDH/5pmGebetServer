package web

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func AssertStatus(t *testing.T, handler http.HandlerFunc, req *http.Request, assertStatus int) {
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != assertStatus {
		t.Errorf("handler %s returned wrong status code: got %v want %v",
			req.URL, status, assertStatus)
	}
}

func TestApiGetAuth(t *testing.T) {
	SetDatabase(testDatabase)
	handler := http.HandlerFunc(apiGet)

	// check StatusUnauthorized for no key
	req, _ := http.NewRequest("GET", "/api", nil)
	AssertStatus(t, handler, req, http.StatusUnauthorized)

	// check StatusUnauthorized for unknown key
	req, _ = http.NewRequest("GET", "/api?key="+"Wrong_Key", nil)
	AssertStatus(t, handler, req, http.StatusUnauthorized)

	// check StatusUnauthorized wrong key type
	req, _ = http.NewRequest("GET", "/api?key="+keyTestView, nil)
	AssertStatus(t, handler, req, http.StatusUnauthorized)

	// check StatusOK for master key
	req, _ = http.NewRequest("GET", "/api?key="+keyTestMaster, nil)
	AssertStatus(t, handler, req, http.StatusOK)

	// check StatusOK for valid/known key
	req, _ = http.NewRequest("GET", "/api?key="+keyTestPray, nil)
	AssertStatus(t, handler, req, http.StatusOK)
}

func TestApiGetKeyAuth(t *testing.T) {
	SetDatabase(testDatabase)
	handler := http.HandlerFunc(ApiKey)

	// check StatusUnauthorized for no key
	req, _ := http.NewRequest("GET", "/api/key", nil)
	AssertStatus(t, handler, req, http.StatusUnauthorized)

	// check StatusUnauthorized for unknown key
	req, _ = http.NewRequest("GET", "/api/key?key="+"Wrong_Key", nil)
	AssertStatus(t, handler, req, http.StatusUnauthorized)

	// check StatusUnauthorized wrong key type
	req, _ = http.NewRequest("GET", "/api/view?key="+keyTestPray, nil)
	AssertStatus(t, handler, req, http.StatusUnauthorized)
	req, _ = http.NewRequest("GET", "/api/view?key="+keyTestView, nil)
	AssertStatus(t, handler, req, http.StatusUnauthorized)

	// check StatusOK for master key
	req, _ = http.NewRequest("GET", "/api/view?key="+keyTestMaster, nil)
	AssertStatus(t, handler, req, http.StatusOK)

}

func TestApiGetViewAuth(t *testing.T) {
	SetDatabase(testDatabase)
	handler := http.HandlerFunc(ApiView)

	// check StatusUnauthorized for no key
	req, _ := http.NewRequest("GET", "/api/view", nil)
	AssertStatus(t, handler, req, http.StatusUnauthorized)

	// check StatusUnauthorized for unknown key
	req, _ = http.NewRequest("GET", "/api/view?key="+"Wrong_Key", nil)
	AssertStatus(t, handler, req, http.StatusUnauthorized)

	// check StatusUnauthorized wrong key type
	req, _ = http.NewRequest("GET", "/api/view?key="+keyTestPray, nil)
	AssertStatus(t, handler, req, http.StatusUnauthorized)

	// check StatusOK for master key
	req, _ = http.NewRequest("GET", "/api/view?key="+keyTestMaster, nil)
	AssertStatus(t, handler, req, http.StatusOK)

	// check StatusOK for valid/known key
	req, _ = http.NewRequest("GET", "/api/view?key="+keyTestView, nil)
	AssertStatus(t, handler, req, http.StatusOK)

}
