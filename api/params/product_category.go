package params

import (
	"database/sql"
	"project/e-commerce/db/sqlc"
)

type CreateProductCategoryReqest struct {
	Name string        `json:"name" binding:"required"`
	Sort sql.NullInt32 `json:"sort" binding:"required"`
}

type CreateProductCategoryResponse struct {
	ID   int64         `json:"id"`
	Name string        `json:"name"`
	Sort sql.NullInt32 `json:"sort"`
}

type ListProductCategoryReqest struct {
	CategoryId int64 `json:"category_id" binding:"required"`
}

type ListProductCategoryReponse struct {
	Categories []sqlc.ProductCategory `json:"categories"`
}
