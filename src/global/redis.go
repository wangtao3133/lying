package global

import (
	"config"
	"errors"
	"github.com/go-redis/redis"
)

var (
	redisClientMap      map[string]*redis.Client
	errInvalidRedisNode = errors.New("config redis node is nil")
)

const (
	captcha = "captcha"
	login   = "login"
	order   = "order"
)

func InitRedis(configRedis []config.RedisConfig) error {
	redisClientMap = make(map[string]*redis.Client)
	length := len(configRedis)
	for i := 0; i < length; i++ {
		node := configRedis[i].Name
		if len(node) == 0 {
			return errInvalidRedisNode
		}
		host := configRedis[i].Host
		password := configRedis[i].Password
		db := configRedis[i].DB
		redisClientMap[node] = redis.NewClient(&redis.Options{
			Addr:     host,
			Password: password,
			DB:       db,
		})

		_, err := redisClientMap[node].Ping().Result()
		if err != nil {
			return err
		}
	}
	return nil
}

// 获取验证码redis库
func GetCaptcha() *redis.Client {
	return redisClientMap[captcha]
}

// 获取登录redis库
func GetLogin() *redis.Client {
	return redisClientMap[login]
}

// 获取订单redis库
func GetOrder() *redis.Client {
	return redisClientMap[order]
}
