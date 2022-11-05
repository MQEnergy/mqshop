package global

import (
	"github.com/MQEnergy/mqshop/config"
	"github.com/MQEnergy/mqshop/pkg/lib"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB      // Mysql数据库
	Logger *lib.Logger   // 日志
	Redis  *redis.Client // redis连接池
	Router *gin.Engine   // 路由
	Cfg    *config.Conf  // yaml配置
)
