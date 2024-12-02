package eb_system_config_service

import (
	"crmeb_go/internal/data/redis_data"
	service_data "crmeb_go/internal/data/sevice_data"
	"crmeb_go/internal/model"
	"crmeb_go/internal/server"
	"errors"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type GetSystemConfigInfoImpl interface {
	GetSystemConfigInfo(params service_data.GetSystemConfigParams) (data service_data.GetSystemConfigResult, err error)
}

type GetSystemConfigInfoService struct {
	svc *server.SvcContext
}

// NewEbSystemConfigService 新配置表 模型服务实例
func NewGetSystemConfigInfoService(svc *server.SvcContext) *GetSystemConfigInfoService {
	return &GetSystemConfigInfoService{svc: svc}
}

func (g GetSystemConfigInfoService) GetSystemConfigInfo(params service_data.GetSystemConfigParams) (data service_data.GetSystemConfigResult, err error) {
	var (
		ebSystemConfig model.EbSystemConfig
		//ebSystemConfigList []model.EbSystemConfig
	)

	// 如果同步配置没有开启
	if !g.svc.Conf.System.AsyncConfig {

		data.Name = ebSystemConfig.Name
		data.Value = ebSystemConfig.Value
		data.CreateTime = ebSystemConfig.CreateTime

		return data, nil
	}

	//检测redis是否为空
	exists, err := g.svc.RedisClient.Exists(params.Ctx, redis_data.ConfigList).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		g.svc.Logger.Error("g.svc.RedisClient.HMGet [err]:%v", zap.Error(err))

		return data, err
	}

	if exists != 1 {
		// 将配置设置到redis中
		//err = g.svc.Repo.EbSystemConfigRepository.QueryAll(params.Ctx, "", []any{}, &modelList)
		//if err != nil {
		//	g.svc.Logger.Error("EbSystemConfigRepository.QueryAll [err]:%v", zap.Error(err))
		//
		//	return data, err
		//}
		//
		//lo.ForEach(modelList, func(item eb_system_config_model.EbSystemConfigModel, index int) {
		//	err := g.svc.RedisClient.HMSet(params.Ctx, redis_data.ConfigList, item.Name, item.Value).Err()
		//	if err != nil {
		//		g.svc.Logger.Error("RedisClient.HMSet [err]:%v", zap.Error(err))
		//	}
		//})
	}

	value, err := g.svc.RedisClient.HMGet(params.Ctx, redis_data.ConfigList, params.Name).Result()
	if err != nil {
		g.svc.Logger.Error("EbSystemConfigRepository.QueryAll [err]:%v", zap.Error(err))

		return data, err
	}

	data.Value = value[0].(string)

	return data, nil
}
