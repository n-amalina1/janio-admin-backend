package api

import (
	"backend/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetClientNewOrders() []models.ClientOrder {
	apiUrl := "http://localhost:8008/new/orders"
	res, err := http.Get(apiUrl)
	if err != nil {
		fmt.Printf("Get Client New Orders: %v", err.Error())
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Get Client New Orders: %v", err.Error())
	}

	var order []models.ClientOrder
	json.Unmarshal(body, &order)

	return order
}

func PostOrdersClient(db *sql.DB, orders *[]models.ClientToDBOrder) ([]models.ClientToDBOrder, error) {

	for _, o := range *orders {
		_, err := PostOrdersClientToDb(db, o)
		if err != nil {
			return []models.ClientToDBOrder{}, err
		}
	}

	return *orders, nil
}

func FormatClientToDB(orders []models.ClientOrder) []models.ClientToDBOrder {
	var ordersC []models.ClientToDBOrder

	for _, o := range orders {
		var order models.ClientToDBOrder

		order.OrderLength = o.OrderDetails.OrderLength
		order.OrderWidth = o.OrderDetails.OrderWidth
		order.OrderHeight = o.OrderDetails.OrderHeight
		order.OrderWeight = o.OrderDetails.OrderWeight
		order.OrderStatus = "Pending"
		order.Consignee = o.Address.Consignee
		order.Pickup = o.Address.Pickup
		order.Items = o.Items

		ordersC = append(ordersC, order)
	}
	return ordersC
}
