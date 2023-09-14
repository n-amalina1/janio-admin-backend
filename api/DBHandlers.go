package api

import (
	models "backend/models"
	"database/sql"
	"fmt"
)

func GetAllOrdersDB(db *sql.DB) ([]models.GetOrderDBParams, error) {
	var (
		orders []models.GetOrderDBParams
	)

	rows, err := db.Query("SELECT * FROM orders JOIN consignee ON orders.order_consignee_id = consignee.consignee_id JOIN pickup ON orders.order_pickup_id = pickup.pickup_id")
	if err != nil {
		return nil, fmt.Errorf("Get All Orders DB: %v", err)
	}

	defer rows.Close()
	for rows.Next() {
		var order models.GetOrderDBParams

		if err := rows.Scan(&order.OrderID, &order.OrderLength, &order.OrderWidth, &order.OrderHeight, &order.OrderWeight, &order.OrderConsigneeId, &order.OrderPickupID, &order.ConsigneeID, &order.ConsigneeName, &order.ConsigneePhoneNumber, &order.ConsigneeCountry, &order.ConsigneeAddress, &order.ConsigneePostal, &order.ConsigneeState, &order.ConsigneeCity, &order.ConsigneeProvince, &order.ConsigneeEmail, &order.PickupID, &order.PickupName, &order.PickupPhoneNumber, &order.PickupCountry, &order.PickupAddress, &order.PickupPostal, &order.PickupState, &order.PickupCity, &order.PickupProvince); err != nil {
			return nil, fmt.Errorf("Get All Orders DB:: %v", err)
		}

		orders = append(orders, order)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Get All Orders:: %v", err)
	}

	return orders, nil
}

func GetItemDB(db *sql.DB) ([]models.Item, error) {
	var (
		items []models.Item
	)
	rows, err := db.Query("SELECT * FROM item")
	if err != nil {
		return nil, fmt.Errorf("Get All Orders Items DB: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var item models.Item

		columns, _ := rows.Columns()
		for _, value := range columns {
			fmt.Printf("* %s\n", value)
		}

		if err := rows.Scan(&item.ItemID, &item.ItemDescription, &item.ItemCategory, &item.ItemSku, &item.ItemQuantity, &item.ItemPrice, &item.ItemCurrency, &item.ItemOrderID); err != nil {
			return nil, fmt.Errorf("Get All Orders Items DB: %v", err)
		}
		items = append(items, item)
	}
	return items, nil
}