package sredis

import (
	"context"
	"time"

	"github.com/MQEnergy/mqshop/internal/vars"
	"github.com/redis/go-redis/v9"
)

const (
	EmployeeFmt     = "ely:%s" // 员工信息key
	OrganizationFmt = "org:%s" // 组织信息key
	ApplicationFmt  = "app:%s" // app配置
	MClientIdsFmt   = "mcs:%d" // 商家应用列表
)

func BuildAuthRedisKey(key string) string {
	return vars.Config.GetString("redis.prefix") + key
}

func Set(ctx context.Context, key string, value string) *redis.StatusCmd {
	rKey := BuildAuthRedisKey(key)
	return vars.Redis.Set(ctx, rKey, value, vars.Config.GetDuration("jwt.expire")*time.Second)
}

func Del(ctx context.Context, key string) *redis.IntCmd {
	rKey := BuildAuthRedisKey(key)
	return vars.Redis.Del(ctx, rKey)
}

func Get(ctx context.Context, key string) *redis.StringCmd {
	rKey := BuildAuthRedisKey(key)
	return vars.Redis.Get(ctx, rKey)
}

func Exists(ctx context.Context, key string) *redis.IntCmd {
	rKey := BuildAuthRedisKey(key)
	return vars.Redis.Exists(ctx, rKey)
}
