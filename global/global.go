package global

import (
	"database/sql"
	"project/e-commerce/config"

	"github.com/go-redis/redis/v8"

	"go.uber.org/zap"
)

const DateFormat = "2006-01-02 15:04:05"

var (
	CONFIG config.Config
	LOG    *zap.Logger
	DB     *sql.DB
	REDIS  *redis.Client
)
