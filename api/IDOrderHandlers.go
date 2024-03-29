package api

import (
	models "backend/models"
	"database/sql"
	"fmt"
)

func GetOrdersIDProvider(db *sql.DB) (models.IDProviderOrdersParams, error) {
	var (
		IDProviderOrdersParams models.IDProviderOrdersParams
		IDOrders               []models.IDOrder
	)

	orders, _ := GetAllOrdersDB(db, "Indonesia")
	for _, order := range orders {
		IDOrder := FormatDbToID(db, order)

		IDOrders = append(IDOrders, IDOrder)
	}

	IDProviderOrdersParams.Orders = IDOrders
	return IDProviderOrdersParams, nil
}

func FormatDbToID(db *sql.DB, order models.GetOrderDBParams) models.IDOrder {
	var IDOrder models.IDOrder
	IDOrder.OrderID = order.OrderID
	IDOrder.OrderLength = order.OrderLength
	IDOrder.OrderWidth = order.OrderWidth
	IDOrder.OrderHeight = order.OrderHeight
	IDOrder.OrderWeight = order.OrderWeight
	IDOrder.OrderStatus = order.OrderStatus
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
	itemOrders, err := GetItemDB(db, order.OrderID)
	if err != nil {
		fmt.Printf("Get Item Orders DB: %v", err)
	}
	for _, itemOrder := range itemOrders {
		IDItem := models.IDItem{ItemID: itemOrder.ItemID, ItemDescription: itemOrder.ItemDescription, ItemCategory: itemOrder.ItemCategory, ItemSku: itemOrder.ItemSku, ItemQuantity: itemOrder.ItemQuantity, ItemPrice: itemOrder.ItemPrice, ItemCurrency: itemOrder.ItemCurrency}
		IDItems = append(IDItems, IDItem)

	}
	IDOrder.Items = IDItems

	return IDOrder
}

func PostStatusIDProvider(db *sql.DB, status models.IDOrderStatus) (models.IDOrderStatus, error) {
	_, errS := UpdateStatusDB(db, status)
	if errS != nil {
		return models.IDOrderStatus{}, fmt.Errorf("update Status DB: %v", errS)
	}
	return status, nil
}
