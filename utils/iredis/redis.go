package initialize

import (
	"context"
	"crmeb_go/config"
	"crmeb_go/utils/izap"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

func initRedisClient(redisCfg *config.Redis) (redis.UniversalClient, error) {
	var client redis.UniversalClient
	// 使用集群模式
	if redisCfg.UseCluster {
		client = redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:    redisCfg.ClusterAddrs,
			Password: redisCfg.Password,
		})
	} else {
		// 使用单例模式
		client = redis.NewClient(&redis.Options{
			Addr:     redisCfg.Addr,
			Password: redisCfg.Password,
			DB:       redisCfg.DB,
		})
	}
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		izap.Log.Error("redis connect ping failed, errorm:", zap.String("name", redisCfg.Name), zap.Error(err))

		return nil, err
	}

	izap.Log.Info("redis connect ping response:", zap.String("name", redisCfg.Name), zap.String("pong", pong))
	return client, nil
}

func Redis(config *config.Conf) redis.UniversalClient {
	redisClient, err := initRedisClient(&config.Redis)
	if err != nil {
		panic(err)
	}

	return redisClient
}

func RedisList(config *config.Conf) map[string]redis.UniversalClient {
	redisMap := make(map[string]redis.UniversalClient)

	for _, redisCfg := range config.RedisList {
		client, err := initRedisClient(&redisCfg)
		if err != nil {
			panic(err)
		}
		redisMap[redisCfg.Name] = client
	}

	return redisMap
}
