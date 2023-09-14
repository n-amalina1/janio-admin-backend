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

type AdminItem struct {
	ItemID          int     `json:"item_id"`
	ItemDescription string  `json:"item_desc"`
	ItemCategory    string  `json:"item_category"`
	ItemSku         string  `json:"item_sku"`
	ItemQuantity    int     `json:"item_quantity"`
	ItemPrice       float64 `json:"item_price"`
	ItemCurrency    string  `json:"item_currency"`
}
