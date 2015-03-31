// Unit test for our simple web server

package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer(t *testing.T) {
	request, _ := http.NewRequest("GET", "", nil)
	recorder := httptest.NewRecorder()
	handler := getHomeHandler()
	handler.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Error("Status not 200")
	}

	content := recorder.Body.String()
	if content != TITLE {
		t.Error("Unexpected content: %v", recorder.Body)
	}
}
