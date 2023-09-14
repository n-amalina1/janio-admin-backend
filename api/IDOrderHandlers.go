package api

import (
	models "backend/models"
	"database/sql"
)

func GetOrdersIDProvider(db *sql.DB) (models.IDProviderOrdersParams, error) {
	var (
		IDProviderOrdersParams models.IDProviderOrdersParams
		IDOrders               []models.IDOrder
	)

	orders, _ := GetAllOrdersDB(db)
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
		IDItem := models.IDItem{ItemID: item.ItemID, ItemDescription: item.ItemDescription, ItemCategory: item.ItemCategory, ItemSku: item.ItemSku, ItemQuantity: item.ItemQuantity, ItemPrice: item.ItemPrice, ItemCurrency: item.ItemCurrency}
		IDItems = append(IDItems, IDItem)
	}

	IDOrder.Items = IDItems
	return IDOrder
}
