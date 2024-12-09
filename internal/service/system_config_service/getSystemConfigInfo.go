package system_config_service

import (
	"crmeb_go/internal/data/redis_data"
	service_data "crmeb_go/internal/data/sevice_data"
	"crmeb_go/internal/model"
	"crmeb_go/internal/server"
	"crmeb_go/utils/izap"
	"errors"
	"github.com/redis/go-redis/v9"
	"github.com/samber/lo"
	"go.uber.org/zap"
	"time"
)

type GetSystemConfigInfoImpl interface {
	GetSystemConfigInfo(params service_data.GetSystemConfigParams) (data service_data.GetSystemConfigResult, err error)
}

type GetSystemConfigInfoService struct {
	svc *server.SvcContext
}

// NewSystemConfigService 新配置表 模型服务实例
func NewGetSystemConfigInfoService(svc *server.SvcContext) *GetSystemConfigInfoService {
	return &GetSystemConfigInfoService{svc: svc}
}

func (g GetSystemConfigInfoService) GetSystemConfigInfo(params service_data.GetSystemConfigParams) (data service_data.GetSystemConfigResult, err error) {
	// 如果同步配置没有开启
	if !g.svc.Conf.System.AsyncConfig {
		SystemConfig, err := g.svc.Repo.SystemConfigRepository.QueryByName(params.Ctx, params.Name)
		if err != nil {
			izap.Log.Error("EbSystemConfigRepository.QueryOne [err]:%v", zap.Error(err))

			return data, err
		}
		data.Name = SystemConfig.Name
		data.Value = SystemConfig.Value
		data.CreateTime = time.Unix(SystemConfig.CreatedAt, 0)

		return data, nil
	}

	// 检测redis是否为空
	exists, err := g.svc.RedisClient.Exists(params.Ctx, redis_data.ConfigList).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		izap.Log.Error("g.svc.RedisClient.HMGet [err]:%v", zap.Error(err))

		return data, err
	}

	if exists != 1 {
		// 将配置设置到redis中
		systemConfigs, err := g.svc.Repo.SystemConfigRepository.All(params.Ctx)
		if err != nil {
			izap.Log.Error("EbSystemConfigRepository.QueryAll [err]:%v", zap.Error(err))

			return data, err
		}

		lo.ForEach(systemConfigs, func(item *model.SystemConfig, index int) {
			err := g.svc.RedisClient.HMSet(params.Ctx, redis_data.ConfigList, item.Name, item.Value).Err()
			if err != nil {
				izap.Log.Error("RedisClient.HMSet [err]:%v", zap.Error(err))
			}
		})
	}

	value, err := g.svc.RedisClient.HMGet(params.Ctx, redis_data.ConfigList, params.Name).Result()
	if err != nil {
		izap.Log.Error("EbSystemConfigRepository.QueryAll [err]:%v", zap.Error(err))

		return data, err
	}

	data.Value = value[0].(string)
	return data, nil
}
