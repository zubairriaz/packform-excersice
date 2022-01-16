package main

import (
	"testing"
   "net/http"
   "net/http/httptest"
   "github.com/gorilla/mux"
   )


   func TestEmptyTable(t *testing.T) {

    req, _ := http.NewRequest("GET", "/api/orders", nil)
    response := executeRequest(req)

    checkResponseCode(t, http.StatusOK, response.Code)

    if body := response.Body.String(); body != "[]" {
        t.Errorf("Expected an empty array. Got %s", body)
    }
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	a :=mux.NewRouter()
    rr := httptest.NewRecorder()
    a.ServeHTTP(rr, req)
	return rr
}


func checkResponseCode(t *testing.T, expected, actual int) {
    if expected != actual {
        t.Errorf("Expected response code %d. Got %d\n", expected, actual)
    }
}