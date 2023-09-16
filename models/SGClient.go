package models

type ClientOrder struct {
	OrderDetails OrderDetails `json:"order_details"`
	Address      Address      `json:"address"`
	Items        []ClientItem `json:"items"`
}

type OrderDetails struct {
	OrderLength float64 `json:"order_length"`
	OrderWidth  float64 `json:"order_width"`
	OrderHeight float64 `json:"order_height"`
	OrderWeight float64 `json:"order_weight"`
}

type Address struct {
	ClientConsignee ClientConsignee `json:"consignee"`
	ClientPickup    ClientPickup    `json:"pickup"`
}

type ClientConsignee struct {
	ConsigneeName        string `json:"consignee_name"`
	ConsigneePhoneNumber string `json:"consignee_phone_number"`
	ConsigneeCountry     string `json:"consignee_country"`
	ConsigneeAddress     string `json:"consignee_address"`
	ConsigneePostal      int    `json:"consignee_postal"`
	ConsigneeState       string `json:"consignee_state"`
	ConsigneeCity        string `json:"consignee_city"`
	ConsigneeProvince    string `json:"consignee_province"`
	ConsigneeEmail       string `json:"consignee_email"`
}

type ClientPickup struct {
	PickupName        string `json:"pickup_name"`
	PickupPhoneNumber string `json:"pickup_phone_number"`
	PickupCountry     string `json:"pickup_country"`
	PickupAddress     string `json:"pickup_address"`
	PickupPostal      int    `json:"pickup_postal"`
	PickupState       string `json:"pickup_state"`
	PickupCity        string `json:"pickup_city"`
	PickupProvince    string `json:"pickup_province"`
}

type ClientItem struct {
	ItemDescription string  `json:"item_desc"`
	ItemCategory    string  `json:"item_category"`
	ItemSku         string  `json:"item_sku"`
	ItemQuantity    int     `json:"item_quantity"`
	ItemPrice       float64 `json:"item_price"`
	ItemCurrency    string  `json:"item_currency"`
}

type ClientToDBOrder struct {
	OrderLength float64         `json:"order_length"`
	OrderWidth  float64         `json:"order_width"`
	OrderHeight float64         `json:"order_height"`
	OrderWeight float64         `json:"order_weight"`
	OrderStatus string          `json:"order_status"`
	Consignee   ClientConsignee `json:"consignee"`
	Pickup      ClientPickup    `json:"pickup"`
	Items       []ClientItem    `json:"items"`
}
