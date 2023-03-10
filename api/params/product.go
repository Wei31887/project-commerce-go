package params

import "project/e-commerce/db/sqlc"

type CreateProductRequest struct {
	Name        string `json:"name" binding:"required"`
	CategoryID  int64  `json:"category_id" binding:"required"`
	Image       string `json:"image" binding:"required"`
	Stock       int64  `json:"stock" binding:"required"`
	Sell        int64  `json:"sell" binding:"required"`
	Price       string `json:"price" binding:"required"`
	OnSell      string `json:"on_sell" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type GetProductRequest struct {
	ProductId int64 `json:"product_id" binding:"required"`
}

type ProductResponse struct {
	Product sqlc.Product `json:"product"`
}

type ListProductRequest struct {
	Limit  int32 `json:"limit" binding:"required,min=1"`
	Offset int32 `json:"offset"`
}

type ListProductByTypeRequest struct {
	CategoryID int64 `json:"category_id" binding:"required"`
	Limit      int32 `json:"limit" binding:"required,min=1"`
	Offset     int32 `json:"offset" binding:"required,min=1"`
}

type ListProductResponse struct {
	Product []sqlc.Product `json:"product"`
}
