package main

import (
	"testing"
   "net/http"
   "net/http/httptest"
   "os"
)

var a App
func TestMain(m *testing.M) {
	var Dbname = goDotEnvVariable("DBNAME");
    var Password = goDotEnvVariable("PASSWORD")
    var User = goDotEnvVariable("USER")
    a.Initialize(User, Password, Dbname)
	code := m.Run()
	os.Exit(code)
}
   
func TestAPI(t *testing.T) {

    req, err := http.NewRequest("GET", "/api/orders?count=1", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(a.getOrders)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `[{"TotalAmount":13.454,"Quantity":"10","DeliveredQuantity":"5","DeliveredAmount":6.727,"UnitPrice":1.3454,"OrderName":"PO #001-I","CompanyName":"Roga \u0026 Kopyta","CustomerName":"Ivan Ivanovich","OrderDate":"2020-01-02T15:34:12Z","FullCount":15}]`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
