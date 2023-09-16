package models

type GetOrderDBParams struct {
	OrderID              int     `json:"order_id"`
	OrderLength          float64 `json:"order_length"`
	OrderWidth           float64 `json:"order_width"`
	OrderHeight          float64 `json:"order_height"`
	OrderWeight          float64 `json:"order_weight"`
	OrderStatus          string  `json:"order_status"`
	OrderConsigneeId     int     `json:"order_consignee_id"`
	OrderPickupID        int     `json:"order_pickup_id"`
	ConsigneeID          int     `json:"consignee_id"`
	ConsigneeName        string  `json:"consignee_name"`
	ConsigneePhoneNumber string  `json:"consignee_phone_number"`
	ConsigneeCountry     string  `json:"consignee_country"`
	ConsigneeAddress     string  `json:"consignee_address"`
	ConsigneePostal      int     `json:"consignee_postal"`
	ConsigneeState       string  `json:"consignee_state"`
	ConsigneeCity        string  `json:"consignee_city"`
	ConsigneeProvince    string  `json:"consignee_province"`
	ConsigneeEmail       string  `json:"consignee_email"`
	PickupID             int     `json:"pickup_id"`
	PickupName           string  `json:"pickup_name"`
	PickupPhoneNumber    string  `json:"pickup_phone_number"`
	PickupCountry        string  `json:"pickup_country"`
	PickupAddress        string  `json:"pickup_address"`
	PickupPostal         int     `json:"pickup_postal"`
	PickupState          string  `json:"pickup_state"`
	PickupCity           string  `json:"pickup_city"`
	PickupProvince       string  `json:"pickup_province"`
}

type ItemDBParams struct {
	ItemID          int     `json:"item_id"`
	ItemDescription string  `json:"item_desc"`
	ItemCategory    string  `json:"item_category"`
	ItemSku         string  `json:"item_sku"`
	ItemQuantity    int     `json:"item_quantity"`
	ItemPrice       float64 `json:"item_price"`
	ItemCurrency    string  `json:"item_currency"`
}

type ItemOrderDBParams struct {
	IOItemID        int     `json:"io_item_id"`
	IOOrderID       int     `json:"io_order_id"`
	ItemID          int     `json:"item_id"`
	ItemDescription string  `json:"item_desc"`
	ItemCategory    string  `json:"item_category"`
	ItemSku         string  `json:"item_sku"`
	ItemQuantity    int     `json:"item_quantity"`
	ItemPrice       float64 `json:"item_price"`
	ItemCurrency    string  `json:"item_currency"`
}

type User struct {
	UserID       int    `json:"id"`
	UserName     string `json:"name"`
	UserEmail    string `json:"email"`
	UserPassword string `json:"password"`
}

type Order struct {
	OrderID     int     `json:"order_id"`
	OrderLength float64 `json:"order_length"`
	OrderWidth  float64 `json:"order_width"`
	OrderHeight float64 `json:"order_height"`
	OrderWeight float64 `json:"order_weight"`
	OrderStatus string  `json:"order_status"`
}

type Consignee struct {
	ConsigneeID          int    `json:"consignee_id"`
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

type Pickup struct {
	PickupID          int    `json:"pickup_id"`
	PickupName        string `json:"pickup_name"`
	PickupPhoneNumber string `json:"pickup_phone_number"`
	PickupCountry     string `json:"pickup_country"`
	PickupAddress     string `json:"pickup_address"`
	PickupPostal      int    `json:"pickup_postal"`
	PickupState       string `json:"pickup_state"`
	PickupCity        string `json:"pickup_city"`
	PickupProvince    string `json:"pickup_province"`
}

type Item struct {
	ItemID          int     `json:"item_id"`
	ItemDescription string  `json:"item_desc"`
	ItemCategory    string  `json:"item_category"`
	ItemSku         string  `json:"item_sku"`
	ItemQuantity    int     `json:"item_quantity"`
	ItemPrice       float64 `json:"item_price"`
	ItemCurrency    string  `json:"item_currency"`
}
