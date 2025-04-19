package cache 

import (
	"magic_link/adapter/output/model"

	redis "github.com/redis/go-redis/v9"

	"context"
	"time"
)

var ctx = context.Background()

type cache struct {
	rdb *redis.Client
}

func NewCache(rdb *redis.Client) *cache {
	return &cache{
		rdb: rdb,
	}
}

func (c *cache) Set(user model.UserCode) error {
	return c.rdb.Set(ctx, user.Code, user.Email, 5*time.Minute).Err()
}

func (c *cache) Get(code string) (string, error) {
	user, err := c.rdb.Get(ctx, code).Result()
	return user, err 
}