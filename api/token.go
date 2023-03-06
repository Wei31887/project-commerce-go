package api

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"project/e-commerce/api/params"
	"project/e-commerce/global"
	"project/e-commerce/global/response"
	"time"

	"github.com/gin-gonic/gin"
)

func (server *Server) RenewAccessToken(ctx *gin.Context) {
	var req params.RenewAccessTokenRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ErrResponse(ctx, http.StatusBadRequest, err)
		return
	}

	// verify refresh token
	refreshPayload, err := server.TokenMaker.VerifyToken(req.RefreshToken)
	if err != nil {
		err = fmt.Errorf("can't verify refresh token: %v", err)
		response.ErrResponse(ctx, http.StatusBadRequest, err)
		return
	}

	// get the session by refresh token
	session, err := server.Store.GetSession(ctx, refreshPayload.Id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = errors.New("failed to get session")
			response.ErrResponse(ctx, http.StatusForbidden, err)
			return
		}
		response.ErrResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	// verify session
	if session.IsBlocked {
		err = fmt.Errorf("blocked session")
		response.ErrResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	if session.UserID != refreshPayload.UserID {
		err = fmt.Errorf("incorrect session user")
		response.ErrResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	if session.RefreshToken != req.RefreshToken {
		err = fmt.Errorf("mismatched session token")
		response.ErrResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	if time.Now().After(session.ExpiredAt) {
		err := fmt.Errorf("expired session")
		response.ErrResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	accessToken, accessPayload, err := server.TokenMaker.CreateToken(
		refreshPayload.UserID,
		global.CONFIG.JWT.AccessTokenDuration,
	)
	if err != nil {
		response.ErrResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	rsp := params.RenewAccessTokenResponse{
		AccessToken:          accessToken,
		AccessTokenExpiresAt: accessPayload.ExpiresAt,
	}
	response.SuccessResponse(ctx, rsp)
}
