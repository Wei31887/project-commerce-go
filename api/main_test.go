package api

import (
	"os"
	"project/e-commerce/config"
	db "project/e-commerce/db/sqlc"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}

func newTestServer(config config.Config, store db.Store) *Server {

	server := NewServer(config, store)

	return server
}