package api

import (
	"fmt"
	"net/http"
	"project/e-commerce/api/params"
	"project/e-commerce/db/sqlc"
	"project/e-commerce/global/response"
	"project/e-commerce/middleware"
	"project/e-commerce/token"

	"github.com/gin-gonic/gin"
)

func (server *Server) CreateOrder(ctx *gin.Context) {
	var req params.CreateOrderRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ErrResponse(ctx, http.StatusBadRequest, err)
		return
	}

	payload := ctx.MustGet(middleware.AuthorizationPayloadKey).(*token.Payload)

	arg := sqlc.OrderParams{
		UserID: payload.UserID,
		Item:   req.Item,
	}

	order, err := server.Store.OrderTx(ctx, arg)
	if err != nil {
		err = fmt.Errorf("order transaction failed: %v", err)
		response.ErrResponse(ctx, http.StatusBadRequest, err)
		return
	}

	response.SuccessResponse(ctx, order)
}

func (server *Server) ListOrder(ctx *gin.Context) {

	payload := ctx.MustGet(middleware.AuthorizationPayloadKey).(*token.Payload)

	orders, err := server.Store.ListOrder(ctx, payload.UserID)
	if err != nil {
		err = fmt.Errorf("order transaction failed: %v", err)
		response.ErrResponse(ctx, http.StatusBadRequest, err)
		return
	}
	rsp := params.ListOrderResponse{
		Orders: orders,
	}
	response.SuccessResponse(ctx, rsp)
}
