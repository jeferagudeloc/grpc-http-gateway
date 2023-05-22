package domain

type Order struct {
	OrderNumber int64 `json:"orderNumber"`
}

type OrderRequest struct {
	OrderNumber int64 `json:"orderNumber"`
}
