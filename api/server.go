package api

import (
	"project/e-commerce/config"
	db "project/e-commerce/db/sqlc"
	"project/e-commerce/global"
	"project/e-commerce/middleware"
	"project/e-commerce/token"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Store      db.Store
	TokenMaker *token.JwtMaker
	Router     *gin.Engine
}

func NewServer(config config.Config, store db.Store) *Server {
	maker := token.NewJWTMaker(config.JWT.SigningKey)
	server := &Server{
		Store:      store,
		TokenMaker: maker,
	}

	server.setupRouter()
	return server
}

func (server *Server) setupRouter() {
	router := gin.New()
	router.Use()
	router.Use(gin.Recovery())
	router.Static("/static", "static")

	router.Use(middleware.Log(global.LOG))
	router.Use(middleware.Cors())

	server.Router = router

	server.apiGroup()
}

func (server *Server) apiGroup() {
	// basic api
	server.Router.POST("/login", server.Login)
	server.Router.POST("/user", server.CreateUser)
	// renew token
	server.Router.POST("token/renew", server.RenewAccessToken)

	server.Router.POST("/product", server.ListProduct)
	server.Router.POST("/product/category", server.ListAllProductCategory)
	// api with admission
	server.authorizationGroup()
}

// func (server *Server) adminGroup() {
// 	group := server.Router.Group("/admin").Use(middleware.Authorization(server.TokenMaker))
// 	group.POST("/products", server.CreateProduct)
// }

func (server *Server) authorizationGroup() {
	authorGroup := server.Router.Group("/").Use(middleware.Authorization(server.TokenMaker))
	{
		authorGroup.POST("/logout", server.Logout)

		authorGroup.POST("/order", server.CreateOrder)

		authorGroup.POST("/user/info/update", server.UpdateUserInfo)
		authorGroup.POST("/user/password/update", server.UpdateUserPassword)

		authorGroup.POST("/cart/add", server.AddCartItem)
		authorGroup.POST("/cart/delete", server.DeleteCartItem)
		authorGroup.POST("/cart/list", server.ListCartItems)

	}
}

func (server *Server) RunServer(address string) error {
	return server.Router.Run(address)
}
