package middleware

import (
	"errors"
	"net/http"
	"project/e-commerce/token"

	"github.com/gin-gonic/gin"
)

var (
	AccessToken = "access_token"
	AuthorizationPayloadKey = "authorization_payload"
)

func Authorization(tokenMaker *token.JwtMaker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// get the token header
		token := ctx.GetHeader(AccessToken)
		if len(token) == 0 {
			err := errors.New("failed to get access token")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		}
	
		// verify the token header
		payload, err := tokenMaker.VerifyToken(token) 
		if err != nil { 
			err := errors.New("failed to verify access token")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.Set(AuthorizationPayloadKey, payload)
		ctx.Next()
	}
}