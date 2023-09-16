package api

import (
	models "backend/models"
	"database/sql"
	"fmt"
)

func GetOrdersMYProvider(db *sql.DB) (models.MYProviderOrdersParams, error) {
	var (
		MYProviderOrdersParams models.MYProviderOrdersParams
		MYOrders               []models.MYOrder
	)

	orders, _ := GetAllOrdersDB(db)
	for _, order := range orders {
		MYOrder := FormatDbToMY(db, order)
		MYOrders = append(MYOrders, MYOrder)
	}

	MYProviderOrdersParams.Orders = MYOrders
	return MYProviderOrdersParams, nil
}

func FormatDbToMY(db *sql.DB, order models.GetOrderDBParams) models.MYOrder {
	var MYOrder models.MYOrder
	MYOrder.OrderID = order.OrderID

	MYConsignee := models.Consignee{ConsigneeID: order.ConsigneeID, ConsigneeName: order.ConsigneeName, ConsigneePhoneNumber: order.ConsigneePhoneNumber, ConsigneeCountry: order.ConsigneeCountry, ConsigneeAddress: order.ConsigneeAddress, ConsigneePostal: order.ConsigneePostal, ConsigneeState: order.ConsigneeState, ConsigneeCity: order.ConsigneeCity, ConsigneeProvince: order.ConsigneeProvince, ConsigneeEmail: order.ConsigneeEmail}

	MYPickup := models.Pickup{PickupID: order.PickupID, PickupName: order.PickupName, PickupPhoneNumber: order.PickupPhoneNumber, PickupCountry: order.PickupCountry, PickupAddress: order.PickupAddress, PickupPostal: order.PickupPostal, PickupState: order.PickupState, PickupCity: order.PickupCity, PickupProvince: order.PickupProvince}

	MYOrderDetails := models.MYOrderDetails{OrderLength: order.OrderLength, OrderWidth: order.OrderWidth, OrderHeight: order.OrderHeight, OrderWeight: order.OrderWeight, Consignee: MYConsignee, Pickup: MYPickup}
	MYOrder.OrderDetails = MYOrderDetails

	var MYItems []models.MYItem
	itemOrders, err := GetItemDB(db, order.OrderID)
	if err != nil {
		fmt.Printf("Get Item Orders DB: %v", err)
	}
	for _, itemOrder := range itemOrders {
		MYItem := models.MYItem{ItemID: itemOrder.ItemID, ItemDescription: itemOrder.ItemDescription, ItemCategory: itemOrder.ItemCategory, ItemSku: itemOrder.ItemSku, ItemQuantity: itemOrder.ItemQuantity, ItemPrice: itemOrder.ItemPrice, ItemCurrency: itemOrder.ItemCurrency}
		MYItems = append(MYItems, MYItem)
	}

	MYOrder.Items = MYItems
	return MYOrder
}
