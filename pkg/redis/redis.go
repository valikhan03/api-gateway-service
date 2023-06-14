package redis

import(
	"github.com/redis/go-redis/v9"

	"api-gateway-service/pkg/models"
)

var RedisClient *redis.Client

func InitRedisConn()  {
	opts := new(redis.Options)
	opts.Network = models.RedisConfigGlobal.Network
	opts.Addr = models.RedisConfigGlobal.Addr
	opts.ClientName = models.RedisConfigGlobal.ClientName
	opts.Username = models.RedisConfigGlobal.Username
	opts.DB = models.RedisConfigGlobal.DB

	RedisClient = redis.NewClient(opts)
}
