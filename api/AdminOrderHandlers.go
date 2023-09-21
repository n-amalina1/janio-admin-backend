package api

import (
	models "backend/models"
	"database/sql"
	"fmt"
)

func GetOrdersAdmin(db *sql.DB) (models.AdminOrdersParams, error) {
	var (
		adminOrdersParams models.AdminOrdersParams
		adminOrders       []models.AdminOrder
	)

	orders, _ := GetAllOrdersDB(db, "None")
	for _, order := range orders {
		adminOrder := FormatDbToAdmin(db, order)
		adminOrders = append(adminOrders, adminOrder)
	}

	adminOrdersParams.Orders = adminOrders
	return adminOrdersParams, nil
}

func UpdateOrderAdmin(db *sql.DB, order *models.UpdateAdminOrder) (models.UpdateAdminOrder, error) {

	updatedOrder, err := UpdateOrderDB(db, order)

	return updatedOrder, err
}

func PostOrderAdmin(db *sql.DB, order *models.PostAdminOrder) (*models.PostAdminOrder, error) {
	newOrder, err := PostOrderDB(db, order)

	return newOrder, err
}

func DeleteOrderAdmin(db *sql.DB, id int) (int, error) {
	orderId, err := DeleteOrderDB(db, id)

	return orderId, err
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
	itemOrders, err := GetItemDB(db, order.OrderID)
	if err != nil {
		fmt.Printf("Get Item Orders DB: %v", err)
	}
	for _, itemOrder := range itemOrders {
		adminItem := models.AdminItem{ItemID: itemOrder.ItemID, ItemDescription: itemOrder.ItemDescription, ItemCategory: itemOrder.ItemCategory, ItemSku: itemOrder.ItemSku, ItemQuantity: itemOrder.ItemQuantity, ItemPrice: itemOrder.ItemPrice, ItemCurrency: itemOrder.ItemCurrency}
		adminItems = append(adminItems, adminItem)
	}

	adminOrder.Items = adminItems
	return adminOrder
}
