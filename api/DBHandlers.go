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

		if err := rows.Scan(&order.OrderID, &order.OrderLength, &order.OrderWidth, &order.OrderHeight, &order.OrderWeight, &order.OrderStatus, &order.OrderConsigneeId, &order.OrderPickupID, &order.ConsigneeID, &order.ConsigneeName, &order.ConsigneePhoneNumber, &order.ConsigneeCountry, &order.ConsigneeAddress, &order.ConsigneePostal, &order.ConsigneeState, &order.ConsigneeCity, &order.ConsigneeProvince, &order.ConsigneeEmail, &order.PickupID, &order.PickupName, &order.PickupPhoneNumber, &order.PickupCountry, &order.PickupAddress, &order.PickupPostal, &order.PickupState, &order.PickupCity, &order.PickupProvince); err != nil {
			return nil, fmt.Errorf("Get All Orders DB:: %v", err)
		}

		orders = append(orders, order)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Get All Orders:: %v", err)
	}

	return orders, nil
}

func GetItemDB(db *sql.DB) ([]models.ItemDBParams, error) {
	var (
		items []models.ItemDBParams
	)
	rows, err := db.Query("SELECT * FROM item")
	if err != nil {
		return nil, fmt.Errorf("Get All Orders Items DB: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var item models.ItemDBParams

		if err := rows.Scan(&item.ItemID, &item.ItemDescription, &item.ItemCategory, &item.ItemSku, &item.ItemQuantity, &item.ItemPrice, &item.ItemCurrency, &item.ItemOrderID); err != nil {
			return nil, fmt.Errorf("Get All Orders Items DB: %v", err)
		}
		items = append(items, item)
	}
	return items, nil
}

func GetOrdersAdmin(db *sql.DB) (models.AdminOrdersParams, error) {
	var (
		adminOrdersParams models.AdminOrdersParams
		adminOrders       []models.AdminOrder
	)

	orders, _ := GetAllOrdersDB(db)
	for _, order := range orders {
		adminOrder := FormatDbToAdmin(db, order)
		adminOrders = append(adminOrders, adminOrder)
	}

	adminOrdersParams.Orders = adminOrders
	return adminOrdersParams, nil
}

func FormatDbToAdmin(db *sql.DB, order models.GetOrderDBParams) models.AdminOrder {
	var adminOrder models.AdminOrder

	adminOrder.OrderID = order.OrderID
	adminOrder.OrderLength = order.OrderLength
	adminOrder.OrderWidth = order.OrderWidth
	adminOrder.OrderHeight = order.OrderHeight
	adminOrder.OrderWeight = order.OrderWeight
	adminOrder.OrderStatus = order.OrderStatus

	adminConsignee := models.Consignee{ConsigneeID: order.ConsigneeID, ConsigneeName: order.ConsigneeName, ConsigneePhoneNumber: order.ConsigneePhoneNumber, ConsigneeCountry: order.ConsigneeCountry, ConsigneeAddress: order.ConsigneeAddress, ConsigneePostal: order.ConsigneePostal, ConsigneeState: order.ConsigneeState, ConsigneeCity: order.ConsigneeCity, ConsigneeProvince: order.ConsigneeProvince, ConsigneeEmail: order.ConsigneeEmail}
	adminOrder.Consignee = adminConsignee

	adminPickup := models.Pickup{PickupID: order.PickupID, PickupName: order.PickupName, PickupPhoneNumber: order.PickupPhoneNumber, PickupCountry: order.PickupCountry, PickupAddress: order.PickupAddress, PickupPostal: order.PickupPostal, PickupState: order.PickupState, PickupCity: order.PickupCity, PickupProvince: order.PickupProvince}
	adminOrder.Pickup = adminPickup

	var adminItems []models.AdminItem
	items, _ := GetItemDB(db)
	for _, item := range items {
		adminItem := models.AdminItem{ItemID: item.ItemID, ItemDescription: item.ItemDescription, ItemCategory: item.ItemCategory, ItemSku: item.ItemSku, ItemQuantity: item.ItemQuantity, ItemPrice: item.ItemPrice, ItemCurrency: item.ItemCurrency}
		adminItems = append(adminItems, adminItem)
	}

	adminOrder.Items = adminItems
	return adminOrder
}
