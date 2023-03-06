package params

import "project/e-commerce/db/sqlc"

type AddProductParams struct {
	ProductId int64 `json:"product_id" binding:"required"`
	Count  int64 `json:"count" binding:"required"`
}

type AddCartItemRequest struct {
	AddItems []AddProductParams `json:"add_items" binding:"required"`
}

type DeleteCartItemRequest struct {
	ProductId int64 `json:"product_id" binding:"required"`
}

type CartItemResponse struct {
	Products []sqlc.ListCartItemRow `json:"products"`
}
