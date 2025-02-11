package global

import (
	"github.com/gomodule/redigo/redis"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// 日志
var (
	SysLog *log.Logger
)

// db
var (
	DB *gorm.DB
)

// redis
var (
	RDB *redis.Pool
)
