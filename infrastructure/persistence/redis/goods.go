package redis

import (
	"go-ddd-api-boilerplate/domain/entity"
	"go-ddd-api-boilerplate/domain/repository"

	"github.com/go-redis/redis"
)

type goodsCacheRepository struct {
	redis *redis.Client
}

func (r *goodsCacheRepository) SaveCache(goods *entity.Goods) (int64, error) {

	// r.redis.Set("key", "value", 0)

	// _, err := r.db.Exec("INSERT INTO table (id) VALUES (?)", "1")

	// if err != nil {
	// 	return 0, err
	// }

	return 0, nil
}

func NewGoodsCacheRepository(redis *redis.Client) repository.GoodsCacheRepository {
	return &goodsCacheRepository{redis: redis}
}
