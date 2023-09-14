package models

type MYProviderOrdersParams struct {
	Orders []MYOrder `json:"orders"`
}

type MYOrder struct {
	OrderID      int            `json:"order_id"`
	OrderDetails MYOrderDetails `json:"order_details"`
	Items        []MYItem       `json:"items"`
}

type MYOrderDetails struct {
	OrderLength float64   `json:"order_length"`
	OrderWidth  float64   `json:"order_width"`
	OrderHeight float64   `json:"order_height"`
	OrderWeight float64   `json:"order_weight"`
	Consignee   Consignee `json:"consignee"`
	Pickup      Pickup    `json:"pickup"`
}

type MYItem struct {
	ItemID          int     `json:"item_product_id"`
	ItemDescription string  `json:"item_desc"`
	ItemCategory    string  `json:"item_category"`
	ItemSku         string  `json:"item_sku"`
	ItemQuantity    int     `json:"item_quantity"`
	ItemPrice       float64 `json:"item_price_value"`
	ItemCurrency    string  `json:"item_price_currency"`
}
