package _var

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
	RC *redis.Client
)

var (
	VIDEO_MAX_LENGTH_LIMIT int64 = 1024 * 1024 * 128 // 128M
)
