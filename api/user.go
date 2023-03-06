package api

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"project/e-commerce/api/params"
	"project/e-commerce/db/sqlc"
	"project/e-commerce/global"
	"project/e-commerce/global/response"
	"project/e-commerce/middleware"
	"project/e-commerce/token"
	"project/e-commerce/utils"

	"github.com/gin-gonic/gin"
)

func (server *Server) createTokenSession(ctx *gin.Context, user sqlc.User) (*params.UserResponse, error) {
	// create token header
	accesToken, aceesPayload, err := server.TokenMaker.CreateToken(user.ID, global.CONFIG.JWT.AccessTokenDuration)
	if err != nil {
		return nil, errors.New("failed to create access token")
	}
	refreshToken, refreshPayload, err := server.TokenMaker.CreateToken(user.ID, global.CONFIG.JWT.RefreshTokenDuration)
	if err != nil {
		return nil, errors.New("failed to create refresh token")
	}

	arg := sqlc.CreateSessionParams{
		ID:           refreshPayload.Id,
		UserID:       user.ID,
		RefreshToken: refreshToken,
		ExpiredAt:    refreshPayload.ExpiresAt,
		UserAgent:    ctx.Request.UserAgent(),
		Ip:           ctx.ClientIP(),
		IsBlocked:    false,
	}

	_, err = server.Store.CreateSession(ctx, arg)
	if err != nil {
		return nil, err
	}

	// user response
	rsp := &params.UserResponse{
		UserID:                user.ID,
		UserName:              user.UserName,
		Email:                 user.Email,
		AccessToken:           accesToken,
		AccesstokenExpiredAt:  aceesPayload.ExpiresAt,
		RefreshToken:          refreshToken,
		RefreshTokenExpiredAt: refreshPayload.ExpiresAt,
		CreatedAt:             user.CreatedAt,
		ChangedPasswordAt:     user.ChangedPasswordAt,
	}
	return rsp, nil
}

// CreateUser
func (server *Server) CreateUser(ctx *gin.Context) {
	var req params.CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ErrResponse(ctx, http.StatusBadRequest, err)
		return
	}

	// validate email address
	_, err := server.Store.GetUserByEmail(ctx, req.Email)
	if err == nil {
		err := errors.New("email address existed")
		response.ErrResponse(ctx, http.StatusForbidden, err)
		return
	}

	hashedPd, err := utils.HashedPassword(req.Password)
	if err != nil {
		response.ErrResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	// query to create user 
	user, err := server.Store.CreateUser(ctx, sqlc.CreateUserParams{
		UserName:       req.UserName,
		HashedPassword: hashedPd,
		Email:          req.Email,
		FullName:       req.FullName,
	})
	if err != nil {
		err := fmt.Errorf("failed to create user: %v", err)
		response.ErrResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	// initialize user's cart
	_, err = server.Store.CreateCart(ctx, user.ID)
	if err != nil {
		err := fmt.Errorf("failed to initialize user's cart: %v", err)
		response.ErrResponse(ctx, http.StatusInternalServerError, err)
		return
	}


	// create token and session
	rsp, err := server.createTokenSession(ctx, user)
	if err != nil {
		response.ErrResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	response.SuccessResponse(ctx, rsp)
}

// UpdateUserInfo
func (server *Server) UpdateUserInfo(ctx *gin.Context) {
	var req params.UpdateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ErrResponse(ctx, http.StatusBadRequest, err)
		return
	}

	payload := ctx.MustGet(middleware.AuthorizationPayloadKey).(*token.Payload)
	if payload.UserID != req.UserID {
		response.ErrResponse(ctx, http.StatusForbidden, errors.New("user ID is not matched"))
	}

	arg := sqlc.UpdateUserInfoParams{
		ID:       req.UserID,
		FullName: req.FullName,
		Email:    req.Email,
		Gender:   req.Gender,
	}
	user, err := server.Store.UpdateUserInfo(ctx, arg)
	if err != nil {
		err := fmt.Errorf("failed to update user's information: %v", err)
		response.ErrResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	rsp := params.UpdateUserResponse{
		UserID:   user.ID,
		Email:    user.Email,
		Gender:   user.Gender,
		FullName: user.UserName,
	}
	response.SuccessResponse(ctx, rsp)
}

// UpdateUserPassword 
func (server *Server) UpdateUserPassword(ctx *gin.Context) {
	var req params.UpdateUserPasswordRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ErrResponse(ctx, http.StatusBadRequest, err)
		return
	}

	payload := ctx.MustGet(middleware.AuthorizationPayloadKey).(*token.Payload)
	if payload.UserID != req.UserID {
		response.ErrResponse(ctx, http.StatusForbidden, errors.New("user ID is not matched"))
	}

	hashedPd, err := utils.HashedPassword(req.Password)
	if err != nil {
		response.ErrResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	user, err := server.Store.UpdateUserPassword(ctx, sqlc.UpdateUserPasswordParams{
		ID:             req.UserID,
		HashedPassword: hashedPd,
	})
	if err != nil {
		err := fmt.Errorf("failed to update user's password: %v", err)
		response.ErrResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	rsp := params.UpdateUserResponse{
		UserID:   user.ID,
		Email:    user.Email,
		Gender:   user.Gender,
		FullName: user.UserName,
	}
	response.SuccessResponse(ctx, rsp)
}

// Login
func (server *Server) Login(ctx *gin.Context) {
	var req params.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ErrResponse(ctx, http.StatusBadRequest, err)
		return
	}

	// Get the user by email
	user, err := server.Store.GetUserByEmail(ctx, req.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			response.ErrResponse(ctx, http.StatusNotFound, errors.New("user not found"))
			return
		}
		response.ErrResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	// validate the user password
	err = utils.ComparedHashedPassword(user.HashedPassword, req.Password)
	if err != nil {
		response.ErrResponse(ctx, http.StatusNotFound, errors.New("password not matched"))
		return
	}

	rsp, err := server.createTokenSession(ctx, user)
	if err != nil {
		response.ErrResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	response.SuccessResponse(ctx, rsp)
}

func (server *Server) Logout(ctx *gin.Context) {

}
