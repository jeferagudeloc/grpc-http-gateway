package domain

type (
	OrderRepository interface {
		SaveOrder(OrderData) (OrderData, error)
	}

	Product struct {
		SKU  string
		Name string
	}

	OrderData struct {
		OrderNumber string
		Products    []Product
	}
)
