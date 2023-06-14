package utils

import (
	"api-gateway-service/pkg/redis"
	"context"
	"time"
)


func SetClientID(token string, id int) error {
	err := redis.RedisClient.Set(context.TODO(), token, id, 24 * time.Hour).Err()
	return err
}

func GetClientID(token string) (int, error) {
	cmd := redis.RedisClient.Get(context.TODO(), token)
	
	return cmd.Int()
}