package domain

import (
	"context"
	"time"
)

type (
	OrderRepository interface {
		SaveOrder(context.Context, OrderData) (OrderData, error)
	}

	OrderData struct {
		OrderNumber  string
		CreationDate time.Time
		UpdationDate time.Time
		Status       string
	}
)
