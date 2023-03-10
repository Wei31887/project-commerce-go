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

func (server *Server) ListCartItems(ctx *gin.Context) {
	// get payload from middlewares
	payload := ctx.MustGet(middleware.AuthorizationPayloadKey).(*token.Payload)

	cartItems, err := server.Store.ListCartItem(ctx, payload.UserID)
	if err != nil {
		response.ErrResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	response.SuccessResponse(ctx, params.CartItemResponse{
		Products: cartItems,
	})
}

func (server *Server) AddCartItem(ctx *gin.Context) {
	var req params.AddCartItemRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ErrResponse(ctx, http.StatusBadRequest, err)
		return
	}

	// get payload from middlewares
	payload := ctx.MustGet(middleware.AuthorizationPayloadKey).(*token.Payload)

	for _, item := range req.AddItems {
		_, err := server.Store.CreateCartItem(ctx, sqlc.CreateCartItemParams{
			CartID: payload.UserID,
			Count:  item.Count,
			ProductID: item.ProductId,
		})
		if err != nil {
			err := fmt.Errorf("faild to add item (product_id: %v, user_id: %d) to cart: %v",
				item.ProductId,
				payload.UserID,
				err,
			)
			response.ErrResponse(ctx, http.StatusInternalServerError, err)
			return
		}
	}

	msg := "add cart item successfully"
	response.MsgResponse(ctx, http.StatusOK, msg)
}

func (server *Server) DeleteCartItem(ctx *gin.Context) {
	var req params.DeleteCartItemRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ErrResponse(ctx, http.StatusBadRequest, err)
		return
	}

	// get payload from middlewares
	payload := ctx.MustGet(middleware.AuthorizationPayloadKey).(*token.Payload)

	err := server.Store.DeleteCartItem(ctx, sqlc.DeleteCartItemParams{
		ID: req.CartItemId,
		CartID: payload.UserID,
	})
	if err != nil {
		response.ErrResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	msg := "delete cart item successfully"
	response.MsgResponse(ctx, http.StatusOK, msg)
}
