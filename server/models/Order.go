package models

import (
	"database/sql"
	"gopkg.in/guregu/null.v4"
	"fmt"
	"strings"
)

type Order struct {
    TotalAmount null.Float `json:"TotalAmount"`
    Quantity string `json:"Quantity"`
    DeliveredQuantity string `json:"DeliveredQuantity"`
	DeliveredAmount null.Float `json:"DeliveredAmount"`
    UnitPrice null.Float `json:"UnitPrice"`
    OrderName string `json:"OrderName"`
	CompanyName string `json:"CompanyName"`
	CustomerName string `json:"CustomerName"`
	OrderDate string `json:"OrderDate"`
	FullCount int32 `json:"FullCount"`


}

func GetOrders(db *sql.DB, start, count int, startDate string, endDate string, searchText string) ([]Order, error) {
	var searchTextLike = ""
	if len(searchText) > 0{
		searchText = strings.Trim(searchText, "\"")	  
		searchTextLike = "%" + searchText + "%"
	}
	var query =   fmt.Sprintf(`Select *, count(*) OVER() AS full_count from (Select (NULLIF(it.price_per_unit, '')::decimal * it.quantity::decimal) as TotalAmount, it.quantity as Quantity, dl.delivered_quantity as DeliveredQuantity,
		  (NULLIF(it.price_per_unit, '')::decimal * dl.delivered_quantity::decimal) as DeliveredAmount, NULLIF(it.price_per_unit, '')::decimal as UnitPrice,
		  ord.order_name as OrderName , co.company_name as CompanyName, cu.name as CustomerName, ord.created_at as OrderDate
		  from items it join deliveries dl on dl.order_item_id = it.id
		  join orders ord on it.order_id = ord.id
		  join customers cu on cu.user_id = ord.customer_id
		  join companies co on co.company_id = cu.company_id where (ord.created_at >= '%s' AND ord.created_at <= '%s' ) AND ('""' = TRIM('%s') or ord.order_name like '%s' )) tbl LIMIT %d OFFSET %d`,
		  startDate, endDate, searchText, searchTextLike, count, start)
	fmt.Println(query)	  
	rows, err := db.Query(query)
  if err != nil {
		  return nil, err
	  }
  defer rows.Close()
  orders := []Order{}
  for rows.Next() {
		  var p Order
		  if err := rows.Scan(&p.TotalAmount, &p.Quantity,&p.DeliveredQuantity, &p.DeliveredAmount, &p.UnitPrice, &p.OrderName, &p.CompanyName, &p.CustomerName, &p.OrderDate, &p.FullCount); err != nil {
			  return nil, err
		  }
		  orders = append(orders, p)
	  }
  return orders, nil
  }