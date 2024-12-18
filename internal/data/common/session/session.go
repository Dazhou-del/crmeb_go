package session

import (
	"crmeb_go/config"
	"crmeb_go/internal/container/repository"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type SvcContext struct {
	RedisClient     redis.UniversalClient
	Logger          *zap.Logger
	Repo            *repository.Container
	Conf            config.Conf
	Gorm            *gorm.DB
	RedisClientList map[string]redis.UniversalClient
	//MQClient *rocketmq.Client
}

type LoginUserInfo struct {
	UserId          int64    `json:"user_id,omitempty"` // 用户id
	PermissionsList []string `json:"permissionsList"`   // 权限标识
}
