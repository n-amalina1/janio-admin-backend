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
		return nil, fmt.Errorf("get All Orders DB: %v", err)
	}

	defer rows.Close()
	for rows.Next() {
		var order models.GetOrderDBParams

		if err := rows.Scan(&order.OrderID, &order.OrderLength, &order.OrderWidth, &order.OrderHeight, &order.OrderWeight, &order.OrderStatus, &order.OrderConsigneeId, &order.OrderPickupID, &order.ConsigneeID, &order.ConsigneeName, &order.ConsigneePhoneNumber, &order.ConsigneeCountry, &order.ConsigneeAddress, &order.ConsigneePostal, &order.ConsigneeState, &order.ConsigneeCity, &order.ConsigneeProvince, &order.ConsigneeEmail, &order.PickupID, &order.PickupName, &order.PickupPhoneNumber, &order.PickupCountry, &order.PickupAddress, &order.PickupPostal, &order.PickupState, &order.PickupCity, &order.PickupProvince); err != nil {
			return nil, fmt.Errorf("get All Orders DB:: %v", err)
		}

		orders = append(orders, order)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("get All Orders:: %v", err)
	}

	return orders, nil
}

func UpdateOrderDB(db *sql.DB, order *models.UpdateAdminOrder) (models.UpdateAdminOrder, error) {

	_, errO := db.Exec("UPDATE orders SET order_length=?, order_width=?, order_height=?, order_weight=?, order_status=? WHERE order_id=?", order.OrderLength, order.OrderWidth, order.OrderHeight, order.OrderWeight, order.OrderStatus, order.OrderID)
	if errO != nil {
		return models.UpdateAdminOrder{}, fmt.Errorf("update Order DB: %s", errO.Error())
	}

	_, errC := db.Exec("UPDATE consignee SET consignee_name=?, consignee_phone_number=?, consignee_country=?, consignee_address=?, consignee_postal=?, consignee_state=?, consignee_city=?, consignee_province=?, consignee_email=? WHERE consignee_id=?", order.Consignee.ConsigneeName, order.Consignee.ConsigneePhoneNumber, order.Consignee.ConsigneeCountry, order.Consignee.ConsigneeAddress, order.Consignee.ConsigneePostal, order.Consignee.ConsigneeState, order.Consignee.ConsigneeCity, order.Consignee.ConsigneeProvince, order.Consignee.ConsigneeEmail, order.Consignee.ConsigneeID)
	if errC != nil {
		return models.UpdateAdminOrder{}, fmt.Errorf("update Order Consignee DB: %s", errC.Error())
	}

	_, errP := db.Exec("UPDATE pickup SET pickup_name=?, pickup_phone_number=?, pickup_country=?, pickup_address=?, pickup_postal=?, pickup_state=?, pickup_city=?, pickup_province=? WHERE pickup_id=?", order.Pickup.PickupName, order.Pickup.PickupPhoneNumber, order.Pickup.PickupCountry, order.Pickup.PickupAddress, order.Pickup.PickupPostal, order.Pickup.PickupState, order.Pickup.PickupCity, order.Pickup.PickupProvince, order.Pickup.PickupID)
	if errP != nil {
		return models.UpdateAdminOrder{}, fmt.Errorf("update Order Pickup DB: %s", errP.Error())
	}

	for _, i := range order.Items {
		var exists bool
		row := db.QueryRow("SELECT EXISTS(SELECT 1 FROM item WHERE item_desc=? AND item_category=? AND item_sku=? AND item_quantity=? AND item_price=? AND item_currency=?", i.ItemDescription, i.ItemCategory, i.ItemSku, i.ItemQuantity, i.ItemPrice, i.ItemCurrency)
		if err := row.Scan(&exists); err == sql.ErrNoRows {
			if _, err := db.Exec("INSERT INTO item (item_desc, item_category, item_sku, item_quantity, item_price, item_currency) VALUES (?, ?, ?, ?, ?, ?)", i.ItemDescription, i.ItemCategory, i.ItemSku, i.ItemQuantity, i.ItemPrice, i.ItemCurrency); err != nil {
				return models.UpdateAdminOrder{}, fmt.Errorf("update Order Item DB: %s", err.Error())
			}
		} else {
			_, errI := db.Exec("UPDATE item SET item_desc=?, item_category=?, item_sku=?, item_quantity=?, item_price=?, item_currency=? WHERE item_id=?", i.ItemDescription, i.ItemCategory, i.ItemSku, i.ItemQuantity, i.ItemPrice, i.ItemCurrency, i.ItemID)
			if errI != nil {
				return models.UpdateAdminOrder{}, fmt.Errorf("update Order Item DB: %s", errI.Error())
			}
		}
	}

	return *order, nil
}

func PostOrderDB(db *sql.DB, newOrder *models.PostAdminOrder) (*models.PostAdminOrder, error) {
	order := *newOrder

	var existsC bool
	var idC int64
	rowC := db.QueryRow("SELECT EXISTS(SELECT 1 FROM consignee WHERE consignee_name=? AND  consignee_phone_number=? AND consignee_country=? AND consignee_address=? AND consignee_postal=? AND  consignee_state=? AND consignee_city=? AND consignee_province=? AND consignee_email=? ", order.Consignee.ConsigneeName, order.Consignee.ConsigneePhoneNumber, order.Consignee.ConsigneeCountry, order.Consignee.ConsigneeAddress, order.Consignee.ConsigneePostal, order.Consignee.ConsigneeState, order.Consignee.ConsigneeCity, order.Consignee.ConsigneeProvince, order.Consignee.ConsigneeEmail)
	if err := rowC.Scan(&existsC); err == sql.ErrNoRows {
		resC, _ := db.Query("SELECT consignee_id FROM consignee WHERE consignee_name=? AND  consignee_phone_number=? AND consignee_country=? AND consignee_address=? AND consignee_postal=? AND  consignee_state=? AND consignee_city=? AND consignee_province=? AND consignee_email=? ", order.Consignee.ConsigneeName, order.Consignee.ConsigneePhoneNumber, order.Consignee.ConsigneeCountry, order.Consignee.ConsigneeAddress, order.Consignee.ConsigneePostal, order.Consignee.ConsigneeState, order.Consignee.ConsigneeCity, order.Consignee.ConsigneeProvince, order.Consignee.ConsigneeEmail)
		for resC.Next() {
			err = resC.Scan(&idC)
			if err != nil {
				return nil, fmt.Errorf("post Order Pickup DB: %s", err.Error())
			}
		}

	} else {

		resC, errC := db.Exec("INSERT INTO consignee (consignee_name, consignee_phone_number, consignee_country, consignee_address, consignee_postal, consignee_state, consignee_city, consignee_province, consignee_email) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)", order.Consignee.ConsigneeName, order.Consignee.ConsigneePhoneNumber, order.Consignee.ConsigneeCountry, order.Consignee.ConsigneeAddress, order.Consignee.ConsigneePostal, order.Consignee.ConsigneeState, order.Consignee.ConsigneeCity, order.Consignee.ConsigneeProvince, order.Consignee.ConsigneeEmail)
		if errC != nil {
			return nil, fmt.Errorf("post Order Consignee DB: %s", errC.Error())
		}
		id, _ := resC.LastInsertId()
		idC = id
	}

	var existsP bool
	var idP int64
	rowP := db.QueryRow("SELECT EXISTS(SELECT 1 FROM pickup WHERE pickup_name=? AND pickup_phone_number=? AND pickup_country=? AND pickup_address=? AND pickup_postal=? AND pickup_state=? AND pickup_city=? AND pickup_province=?", order.Pickup.PickupName, order.Pickup.PickupPhoneNumber, order.Pickup.PickupCountry, order.Pickup.PickupAddress, order.Pickup.PickupPostal, order.Pickup.PickupState, order.Pickup.PickupCity, order.Pickup.PickupProvince)
	if err := rowP.Scan(&existsP); err == sql.ErrNoRows {

		resP, _ := db.Query("SELECT pickup_id FROM pickup WHERE pickup_name=? AND pickup_phone_number=? AND pickup_country=? AND pickup_address=? AND pickup_postal=? AND pickup_state=? AND pickup_city=? AND pickup_province=?", order.Pickup.PickupName, order.Pickup.PickupPhoneNumber, order.Pickup.PickupCountry, order.Pickup.PickupAddress, order.Pickup.PickupPostal, order.Pickup.PickupState, order.Pickup.PickupCity, order.Pickup.PickupProvince)
		for resP.Next() {
			err = resP.Scan(&idP)
			if err != nil {
				return nil, fmt.Errorf("post Order Pickup DB: %s", err.Error())
			}
		}
	} else {

		resP, errP := db.Exec("INSERT INTO pickup (pickup_name, pickup_phone_number, pickup_country, pickup_address, pickup_postal, pickup_state, pickup_city, pickup_province) VALUES (?, ?, ?, ?, ?, ?, ?, ?)", order.Pickup.PickupName, order.Pickup.PickupPhoneNumber, order.Pickup.PickupCountry, order.Pickup.PickupAddress, order.Pickup.PickupPostal, order.Pickup.PickupState, order.Pickup.PickupCity, order.Pickup.PickupProvince)
		if errP != nil {
			return nil, fmt.Errorf("post Order Pickup DB: %s", errP.Error())
		}
		id, _ := resP.LastInsertId()
		idP = id

	}

	resO, errO := db.Exec("INSERT INTO orders (order_length, order_width, order_height, order_weight, order_status, order_consignee_id, order_pickup_id) VALUES (?, ?, ?, ?, ?, ?, ?)", order.OrderLength, order.OrderWidth, order.OrderHeight, order.OrderWeight, order.OrderStatus, idC, idP)
	if errO != nil {
		return nil, fmt.Errorf("post Order DB: %s", errO.Error())
	}
	idO, _ := resO.LastInsertId()

	for _, i := range order.Items {
		var exists bool
		var idI int64
		row := db.QueryRow("SELECT EXISTS(SELECT 1 FROM item WHERE item_desc=? AND item_category=? AND item_sku=? AND item_quantity=? AND item_price=? AND item_currency=?", i.ItemDescription, i.ItemCategory, i.ItemSku, i.ItemQuantity, i.ItemPrice, i.ItemCurrency)
		if err := row.Scan(&exists); err == sql.ErrNoRows {
			if _, err := db.Exec("INSERT INTO item (item_desc, item_category, item_sku, item_quantity, item_price, item_currency) VALUES (?, ?, ?, ?, ?, ?)", i.ItemDescription, i.ItemCategory, i.ItemSku, i.ItemQuantity, i.ItemPrice, i.ItemCurrency); err != nil {
				return nil, fmt.Errorf("post Order Item DB: %s", err.Error())
			}
		} else {
			resI, _ := db.Query("SELECT item_id FROM item WHERE item_desc=? AND item_category=? AND item_sku=? AND item_quantity=? AND item_price=? AND item_currency=?", i.ItemDescription, i.ItemCategory, i.ItemSku, i.ItemQuantity, i.ItemPrice, i.ItemCurrency)
			for resI.Next() {
				err = resI.Scan(&idI)
				if err != nil {
					return nil, fmt.Errorf("post Order Pickup DB: %s", err.Error())
				}
			}
			_, errIO := db.Exec("INSERT INTO item_order (io_item_id, io_order_id) VALUES (?, ?)", idI, idO)
			if errIO != nil {
				return nil, fmt.Errorf("post Order Item DB: %s", errIO.Error())
			}

		}
	}

	return newOrder, nil
}

func PostOrdersClientToDb(db *sql.DB, order models.ClientToDBOrder) (models.ClientToDBOrder, error) {

	var existsC bool
	var idC int64
	rowC := db.QueryRow("SELECT EXISTS(SELECT 1 FROM consignee WHERE consignee_name=? AND  consignee_phone_number=? AND consignee_country=? AND consignee_address=? AND consignee_postal=? AND  consignee_state=? AND consignee_city=? AND consignee_province=? AND consignee_email=? ", order.Consignee.ConsigneeName, order.Consignee.ConsigneePhoneNumber, order.Consignee.ConsigneeCountry, order.Consignee.ConsigneeAddress, order.Consignee.ConsigneePostal, order.Consignee.ConsigneeState, order.Consignee.ConsigneeCity, order.Consignee.ConsigneeProvince, order.Consignee.ConsigneeEmail)
	if err := rowC.Scan(&existsC); err == sql.ErrNoRows {
		resC, _ := db.Query("SELECT consignee_id FROM consignee WHERE consignee_name=? AND  consignee_phone_number=? AND consignee_country=? AND consignee_address=? AND consignee_postal=? AND  consignee_state=? AND consignee_city=? AND consignee_province=? AND consignee_email=? ", order.Consignee.ConsigneeName, order.Consignee.ConsigneePhoneNumber, order.Consignee.ConsigneeCountry, order.Consignee.ConsigneeAddress, order.Consignee.ConsigneePostal, order.Consignee.ConsigneeState, order.Consignee.ConsigneeCity, order.Consignee.ConsigneeProvince, order.Consignee.ConsigneeEmail)
		for resC.Next() {
			err = resC.Scan(&idC)
			if err != nil {
				return models.ClientToDBOrder{}, fmt.Errorf("post Order Pickup DB: %s", err.Error())
			}
		}

	} else {

		resC, errC := db.Exec("INSERT INTO consignee (consignee_name, consignee_phone_number, consignee_country, consignee_address, consignee_postal, consignee_state, consignee_city, consignee_province, consignee_email) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)", order.Consignee.ConsigneeName, order.Consignee.ConsigneePhoneNumber, order.Consignee.ConsigneeCountry, order.Consignee.ConsigneeAddress, order.Consignee.ConsigneePostal, order.Consignee.ConsigneeState, order.Consignee.ConsigneeCity, order.Consignee.ConsigneeProvince, order.Consignee.ConsigneeEmail)
		if errC != nil {
			return models.ClientToDBOrder{}, fmt.Errorf("post Order Consignee DB: %s", errC.Error())
		}
		id, _ := resC.LastInsertId()
		idC = id
	}

	var existsP bool
	var idP int64
	rowP := db.QueryRow("SELECT EXISTS(SELECT 1 FROM pickup WHERE pickup_name=? AND pickup_phone_number=? AND pickup_country=? AND pickup_address=? AND pickup_postal=? AND pickup_state=? AND pickup_city=? AND pickup_province=?", order.Pickup.PickupName, order.Pickup.PickupPhoneNumber, order.Pickup.PickupCountry, order.Pickup.PickupAddress, order.Pickup.PickupPostal, order.Pickup.PickupState, order.Pickup.PickupCity, order.Pickup.PickupProvince)
	if err := rowP.Scan(&existsP); err == sql.ErrNoRows {

		resP, _ := db.Query("SELECT pickup_id FROM pickup WHERE pickup_name=? AND pickup_phone_number=? AND pickup_country=? AND pickup_address=? AND pickup_postal=? AND pickup_state=? AND pickup_city=? AND pickup_province=?", order.Pickup.PickupName, order.Pickup.PickupPhoneNumber, order.Pickup.PickupCountry, order.Pickup.PickupAddress, order.Pickup.PickupPostal, order.Pickup.PickupState, order.Pickup.PickupCity, order.Pickup.PickupProvince)
		for resP.Next() {
			err = resP.Scan(&idP)
			if err != nil {
				return models.ClientToDBOrder{}, fmt.Errorf("post Order Pickup DB: %s", err.Error())
			}
		}
	} else {

		resP, errP := db.Exec("INSERT INTO pickup (pickup_name, pickup_phone_number, pickup_country, pickup_address, pickup_postal, pickup_state, pickup_city, pickup_province) VALUES (?, ?, ?, ?, ?, ?, ?, ?)", order.Pickup.PickupName, order.Pickup.PickupPhoneNumber, order.Pickup.PickupCountry, order.Pickup.PickupAddress, order.Pickup.PickupPostal, order.Pickup.PickupState, order.Pickup.PickupCity, order.Pickup.PickupProvince)
		if errP != nil {
			return models.ClientToDBOrder{}, fmt.Errorf("post Order Pickup DB: %s", errP.Error())
		}
		id, _ := resP.LastInsertId()
		idP = id

	}

	resO, errO := db.Exec("INSERT INTO orders (order_length, order_width, order_height, order_weight, order_status, order_consignee_id, order_pickup_id) VALUES (?, ?, ?, ?, ?, ?, ?)", order.OrderLength, order.OrderWidth, order.OrderHeight, order.OrderWeight, order.OrderStatus, idC, idP)
	if errO != nil {
		return models.ClientToDBOrder{}, fmt.Errorf("post Order DB: %s", errO.Error())
	}
	idO, _ := resO.LastInsertId()

	for _, i := range order.Items {
		var exists bool
		var idI int64
		row := db.QueryRow("SELECT EXISTS(SELECT 1 FROM item WHERE item_desc=? AND item_category=? AND item_sku=? AND item_quantity=? AND item_price=? AND item_currency=?", i.ItemDescription, i.ItemCategory, i.ItemSku, i.ItemQuantity, i.ItemPrice, i.ItemCurrency)
		if err := row.Scan(&exists); err == sql.ErrNoRows {
			if _, err := db.Exec("INSERT INTO item (item_desc, item_category, item_sku, item_quantity, item_price, item_currency) VALUES (?, ?, ?, ?, ?, ?)", i.ItemDescription, i.ItemCategory, i.ItemSku, i.ItemQuantity, i.ItemPrice, i.ItemCurrency); err != nil {
				return models.ClientToDBOrder{}, fmt.Errorf("post Order Item DB: %s", err.Error())
			}
		} else {
			resI, _ := db.Query("SELECT item_id FROM item WHERE item_desc=? AND item_category=? AND item_sku=? AND item_quantity=? AND item_price=? AND item_currency=?", i.ItemDescription, i.ItemCategory, i.ItemSku, i.ItemQuantity, i.ItemPrice, i.ItemCurrency)
			for resI.Next() {
				err = resI.Scan(&idI)
				if err != nil {
					return models.ClientToDBOrder{}, fmt.Errorf("post Order Pickup DB: %s", err.Error())
				}
			}
			_, errIO := db.Exec("INSERT INTO item_order (io_item_id, io_order_id) VALUES (?, ?)", idI, idO)
			if errIO != nil {
				return models.ClientToDBOrder{}, fmt.Errorf("post Order Item DB: %s", errIO.Error())
			}

		}

	}
	return order, nil
}

func DeleteOrderDB(db *sql.DB, orderId int) (int, error) {
	_, errIO := db.Exec("DELETE FROM item_order WHERE io_order_id = ?", orderId)
	if errIO != nil {
		return 0, fmt.Errorf("delete Order DB: %v", errIO)
	}

	_, errO := db.Exec("DELETE FROM orders WHERE order_id = ?", orderId)
	if errO != nil {
		return 0, fmt.Errorf("delete Order DB: %v", errO)
	}

	return orderId, nil
}

func GetItemDB(db *sql.DB, orderId int) ([]models.ItemOrderDBParams, error) {
	var (
		itemOrders []models.ItemOrderDBParams
	)
	rows, err := db.Query("SELECT * FROM item_order JOIN item ON item_order.io_item_id = item.item_id  WHERE io_order_id = ? ", orderId)
	if err != nil {
		return nil, fmt.Errorf("get All Item Orders DB: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var itemOrder models.ItemOrderDBParams

		if err := rows.Scan(&itemOrder.IOItemID, &itemOrder.IOOrderID, &itemOrder.ItemID, &itemOrder.ItemDescription, &itemOrder.ItemCategory, &itemOrder.ItemSku, &itemOrder.ItemQuantity, &itemOrder.ItemPrice, &itemOrder.ItemCurrency); err != nil {
			return nil, fmt.Errorf("get All Items Orders DB: %v", err)
		}
		itemOrders = append(itemOrders, itemOrder)
	}
	return itemOrders, nil
}
