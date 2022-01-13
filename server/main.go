package main
import (
	"log"
	"net/http"
)


func main(){
	http.HandleFun("api/orders", ordersHandler)
	http.HandleFun("api/customers", customersHandler)
	http.HandleFun("api/items", itemsHandler)
	http.HandleFun("api/comapanies", itemsHandler)


}