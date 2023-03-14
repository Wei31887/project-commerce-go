package api

import (
	"net/http"
	"project/e-commerce/api/params"
	"project/e-commerce/db/sqlc"
	"project/e-commerce/global/response"

	"github.com/gin-gonic/gin"
)

func (server *Server) CreateProduct(ctx *gin.Context) {
	var req params.CreateProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ErrResponse(ctx, http.StatusBadRequest, err)
		return
	}

	arg := sqlc.CreateProductParams(req)
	product, err := server.Store.CreateProduct(ctx, arg)
	if err != nil {
		response.ErrResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	rsp := params.ProductResponse{
		Product: product,
	}
	response.SuccessResponse(ctx, rsp)
}

func (server *Server) GetProduct(ctx *gin.Context) {
	var req params.GetProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ErrResponse(ctx, http.StatusBadRequest, err)
		return
	}

	product, err := server.Store.GetProduct(ctx, req.ProductId)
	if err != nil {
		response.ErrResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	rsp := params.ProductResponse{
		Product: product,
	}
	response.SuccessResponse(ctx, rsp)
}

func (server *Server) ListProduct(ctx *gin.Context) {
	var req params.ListProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ErrResponse(ctx, http.StatusBadRequest, err)
		return
	}

	arg := sqlc.ListProductParams{
		Offset: req.Offset,
		Limit:  req.Limit,
	}
	products, err := server.Store.ListProduct(ctx, arg)
	if err != nil {
		response.ErrResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	rsp := params.ListProductResponse{
		Product: products,
	}
	response.SuccessResponse(ctx, rsp)
}

func (server *Server) ListHitProduct(ctx *gin.Context) {
	var req params.ListHitProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ErrResponse(ctx, http.StatusBadRequest, err)
		return
	}
	products, err := server.Store.ListHitProduct(ctx, req.Limit)
	if err != nil {
		response.ErrResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	rsp := params.ListProductResponse{
		Product: products,
	}
	response.SuccessResponse(ctx, rsp)
}

func (server *Server) ListProductByType(ctx *gin.Context) {
	var req params.ListProductByTypeRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ErrResponse(ctx, http.StatusBadRequest, err)
		return
	}

	arg := sqlc.ListProductByTypeParams(req)
	products, err := server.Store.ListProductByType(ctx, arg)
	if err != nil {
		response.ErrResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	rsp := params.ListProductResponse{
		Product: products,
	}
	response.SuccessResponse(ctx, rsp)
}
