package models

type AdminOrdersParams struct {
	Orders []AdminOrder `json:"orders"`
}
type AdminOrder struct {
	OrderID     int         `json:"order_id"`
	OrderLength float64     `json:"order_length"`
	OrderWidth  float64     `json:"order_width"`
	OrderHeight float64     `json:"order_height"`
	OrderWeight float64     `json:"order_weight"`
	OrderStatus string      `json:"order_status"`
	Consignee   Consignee   `json:"consignee"`
	Pickup      Pickup      `json:"pickup"`
	Items       []AdminItem `json:"items"`
}

type UpdateAdminOrder struct {
	OrderID     int         `json:"order_id"`
	OrderLength float64     `json:"order_length"`
	OrderWidth  float64     `json:"order_width"`
	OrderHeight float64     `json:"order_height"`
	OrderWeight float64     `json:"order_weight"`
	OrderStatus string      `json:"order_status"`
	Consignee   Consignee   `json:"consignee"`
	Pickup      Pickup      `json:"pickup"`
	Items       []AdminItem `json:"items"`
}

type PostAdminOrder struct {
	OrderLength int                `json:"order_length"`
	OrderWidth  int                `json:"order_width"`
	OrderHeight int                `json:"order_height"`
	OrderWeight int                `json:"order_weight"`
	OrderStatus string             `json:"order_status"`
	Consignee   PostAdminConsignee `json:"consignee"`
	Pickup      PostAdminPickup    `json:"pickup"`
	Items       []PostAdminItem    `json:"items"`
}

type DeleteAdminOrder struct {
	OrderID int `json:"orderId"`
}

type PostAdminConsignee struct {
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

type PostAdminPickup struct {
	PickupName        string `json:"pickup_name"`
	PickupPhoneNumber string `json:"pickup_phone_number"`
	PickupCountry     string `json:"pickup_country"`
	PickupAddress     string `json:"pickup_address"`
	PickupPostal      int    `json:"pickup_postal"`
	PickupState       string `json:"pickup_state"`
	PickupCity        string `json:"pickup_city"`
	PickupProvince    string `json:"pickup_province"`
}

type PostAdminItem struct {
	ItemDescription string  `json:"item_desc"`
	ItemCategory    string  `json:"item_category"`
	ItemSku         string  `json:"item_sku"`
	ItemQuantity    int     `json:"item_quantity"`
	ItemPrice       float64 `json:"item_price"`
	ItemCurrency    string  `json:"item_currency"`
}

type AdminItem struct {
	ItemID          int     `json:"item_id"`
	ItemDescription string  `json:"item_desc"`
	ItemCategory    string  `json:"item_category"`
	ItemSku         string  `json:"item_sku"`
	ItemQuantity    int     `json:"item_quantity"`
	ItemPrice       float64 `json:"item_price"`
	ItemCurrency    string  `json:"item_currency"`
}
