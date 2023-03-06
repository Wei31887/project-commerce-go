package params

import "project/e-commerce/db/sqlc"

type CreateOrderRequest struct {
	Item   []sqlc.RequestOrderItem `json:"request_order_item"`
}
