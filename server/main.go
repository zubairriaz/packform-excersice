package main

import (
	"log"
	"net/http"
	"github.com/joho/godotenv"
	"fmt"
	"os"
	"database/sql"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

)

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load("../.env")
  
	if err != nil {
	  log.Fatalf("Error loading .env file")
	}
  
	return os.Getenv(key)
  }

  type App struct {
	Router *mux.Router
	DB *sql.DB
   }
   
   func (a *App) Initialize(user, password, dbname string) {
	   connectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)
	  var err error
	   a.DB, err = sql.Open("postgres", connectionString)
	   if err != nil {
	   log.Fatal(err)
	   }
	  a.Router = mux.NewRouter()
	  }
	  
   
   func (a *App) Run(addr string) {
	   log.Fatal(http.ListenAndServe(addr, a.Router))
	  }  

func main(){
	a := App{}
	var Dbname = goDotEnvVariable("DBNAME");
    var Password = goDotEnvVariable("PASSWORD")
    var User = goDotEnvVariable("USER")
    a.Initialize(User, Password, Dbname)
	a.Router.HandleFunc("/api/orders", a.getOrders).Methods("GET")
    a.Run(":8001")
	


}
