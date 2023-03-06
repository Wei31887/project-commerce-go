package api

import (
	"net/http"
	"project/e-commerce/api/params"
	"project/e-commerce/global/response"

	"github.com/gin-gonic/gin"
)

func (server *Server) ListProductCategory(ctx *gin.Context) {
	var req params.ListProductCategoryReqest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ErrResponse(ctx, http.StatusBadRequest, err)
		return
	}

	categories, err := server.Store.ListProductCategory(ctx, req.CategoryId)
	if err != nil {
		response.ErrResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	rsp := params.ListProductCategoryReponse{
		Categories: categories,
	}
	response.SuccessResponse(ctx, rsp)
}

func (server *Server) ListAllProductCategory(ctx *gin.Context) {

	categories, err := server.Store.ListAllProductCategory(ctx)
	if err != nil {
		response.ErrResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	rsp := params.ListProductCategoryReponse{
		Categories: categories,
	}
	response.SuccessResponse(ctx, rsp)
}
