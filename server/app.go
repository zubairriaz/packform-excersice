package main

import (
	"net/http"
	"github.com/zubairriaz/packform-excersice/server/models"
	_ "github.com/lib/pq"
	"strconv"
	"encoding/json"
	"fmt"
)
func setupCorsResponse(w *http.ResponseWriter, req *http.Request) {
    (*w).Header().Set("Access-Control-Allow-Origin", "*")
    (*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
    (*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")
}


func (a *App) getOrders(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: homePage")
    setupCorsResponse(&w, r)
    if (*r).Method == "OPTIONS" {
        return
    }
    count, _ := strconv.Atoi(r.FormValue("count"))
    start, _ := strconv.Atoi(r.FormValue("start"))
    startDate := strconv.Quote(r.FormValue("startDate"))
    endDate := strconv.Quote(r.FormValue("endDate"))
    searchText:= strconv.Quote(r.FormValue("searchText"))


    if count > 10 || count < 1 {
        count = 10
    }
    if start < 0 {
        start = 0
    }
    products, err := models.GetOrders(a.DB, start, count, startDate, endDate, searchText)
    if err != nil {
        respondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }

    respondWithJSON(w, http.StatusOK, products)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
    respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
    response, _ := json.Marshal(payload)

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    w.Write(response)
}