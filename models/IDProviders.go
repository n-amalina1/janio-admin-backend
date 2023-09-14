package models

type IDProviderParams struct {
	Orders []IDOrder `json:"orders"`
}

type IDOrder struct {
	OrderID              int      `json:"order_id"`
	OrderLength          float64  `json:"order_length"`
	OrderWidth           float64  `json:"order_width"`
	OrderHeight          float64  `json:"order_height"`
	OrderWeight          float64  `json:"order_weight"`
	ConsigneeName        string   `json:"consignee_name"`
	ConsigneePhoneNumber string   `json:"consignee_number"`
	ConsigneeCountry     string   `json:"consignee_country"`
	ConsigneeAddress     string   `json:"consignee_address"`
	ConsigneePostal      int      `json:"consignee_postal"`
	ConsigneeState       string   `json:"consignee_state"`
	ConsigneeCity        string   `json:"consignee_city"`
	ConsigneeProvince    string   `json:"consignee_province"`
	ConsigneeEmail       string   `json:"consignee_email"`
	PickupName           string   `json:"pickup_contact_name"`
	PickupPhoneNumber    string   `json:"pickup_contact_number"`
	PickupCountry        string   `json:"pickup_country"`
	PickupAddress        string   `json:"pickup_address"`
	PickupPostal         int      `json:"pickup_postal"`
	PickupState          string   `json:"pickup_state"`
	PickupCity           string   `json:"pickup_city"`
	PickupProvince       string   `json:"pickup_province"`
	Items                []IDItem `json:"items"`
}

type IDItem struct {
	ItemID          int     `json:"item_product_id"`
	ItemDescription string  `json:"item_desc"`
	ItemCategory    string  `json:"item_category"`
	ItemSku         string  `json:"item_sku"`
	ItemQuantity    int     `json:"item_quantity"`
	ItemPrice       float64 `json:"item_price_value"`
	ItemCurrency    string  `json:"item_price_currency"`
}
