package api

import (
	models "backend/models"
	"database/sql"
	"fmt"
)

func GetAllOrdersDB(db *sql.DB) (models.IDProviderParams, error) {
	var (
		IDProviderParams models.IDProviderParams
		IDOrders         []models.IDOrder
	)

	rows, err := db.Query("SELECT * FROM orders JOIN consignee ON orders.order_consignee_id = consignee.consignee_id JOIN pickup ON orders.order_pickup_id = pickup.pickup_id")
	if err != nil {
		return models.IDProviderParams{}, fmt.Errorf("Get All Orders: %v", err)
	}

	defer rows.Close()
	for rows.Next() {
		var order models.GetOrderParams

		if err := rows.Scan(&order.OrderID, &order.OrderLength, &order.OrderWidth, &order.OrderHeight, &order.OrderWeight, &order.OrderConsigneeId, &order.OrderPickupID, &order.ConsigneeID, &order.ConsigneeName, &order.ConsigneePhoneNumber, &order.ConsigneeCountry, &order.ConsigneeAddress, &order.ConsigneePostal, &order.ConsigneeState, &order.ConsigneeCity, &order.ConsigneeProvince, &order.ConsigneeEmail, &order.PickupID, &order.PickupName, &order.PickupPhoneNumber, &order.PickupCountry, &order.PickupAddress, &order.PickupPostal, &order.PickupState, &order.PickupCity, &order.PickupProvince); err != nil {
			return models.IDProviderParams{}, fmt.Errorf("Get All Orders:: %v", err)
		}

		IDOrder := FormatDbToID(db, order)
		IDOrders = append(IDOrders, IDOrder)
	}
	if err := rows.Err(); err != nil {
		return models.IDProviderParams{}, fmt.Errorf("Get All Orders:: %v", err)
	}
	IDProviderParams.Orders = IDOrders

	return IDProviderParams, nil
}

func GetItemDB(db *sql.DB) ([]models.Item, error) {
	var (
		items []models.Item
	)
	rows, err := db.Query("SELECT * FROM item")
	if err != nil {
		return nil, fmt.Errorf("Get All Orders Items: %v", err)
	}
	fmt.Println("Step 1")
	defer rows.Close()
	for rows.Next() {
		fmt.Println("Step 2")
		var item models.Item

		columns, _ := rows.Columns()
		for _, value := range columns {
			fmt.Printf("* %s\n", value)
		}

		if err := rows.Scan(&item.ItemID, &item.ItemDescription, &item.ItemCategory, &item.ItemSku, &item.ItemQuantity, &item.ItemPrice, &item.ItemCurrency, &item.ItemOrderID); err != nil {
			fmt.Println("Step 3")
			return nil, fmt.Errorf("Get All Orders Items: %v", err)
		}
		fmt.Println("Step 4")
		fmt.Printf("%+v\n", item)
		items = append(items, item)
	}
	return items, nil
}

func FormatDbToID(db *sql.DB, order models.GetOrderParams) models.IDOrder {
	var IDOrder models.IDOrder
	IDOrder.OrderID = order.OrderID
	IDOrder.OrderLength = order.OrderLength
	IDOrder.OrderWidth = order.OrderWidth
	IDOrder.OrderHeight = order.OrderHeight
	IDOrder.OrderWeight = order.OrderWeight
	IDOrder.ConsigneeName = order.ConsigneeName
	IDOrder.ConsigneePhoneNumber = order.ConsigneePhoneNumber
	IDOrder.ConsigneeCountry = order.ConsigneeCountry
	IDOrder.ConsigneeAddress = order.ConsigneeAddress
	IDOrder.ConsigneePostal = order.ConsigneePostal
	IDOrder.ConsigneeState = order.ConsigneeState
	IDOrder.ConsigneeCity = order.ConsigneeCity
	IDOrder.ConsigneeProvince = order.ConsigneeProvince
	IDOrder.ConsigneeEmail = order.ConsigneeEmail
	IDOrder.PickupName = order.PickupName
	IDOrder.PickupPhoneNumber = order.PickupPhoneNumber
	IDOrder.PickupCountry = order.PickupCountry
	IDOrder.PickupAddress = order.PickupAddress
	IDOrder.PickupPostal = order.PickupPostal
	IDOrder.PickupState = order.PickupState
	IDOrder.PickupCity = order.PickupCity
	IDOrder.PickupProvince = order.PickupProvince

	var IDItems []models.IDItem
	items, _ := GetItemDB(db)
	for _, item := range items {
		var IDItem models.IDItem
		IDItem.ItemID = item.ItemID
		IDItem.ItemDescription = item.ItemDescription
		IDItem.ItemCategory = item.ItemCategory
		IDItem.ItemSku = item.ItemSku
		IDItem.ItemQuantity = item.ItemQuantity
		IDItem.ItemPrice = item.ItemPrice
		IDItem.ItemCurrency = item.ItemCurrency
		IDItems = append(IDItems, IDItem)
	}

	IDOrder.Items = IDItems
	return IDOrder
}
