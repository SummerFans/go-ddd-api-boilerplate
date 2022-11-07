package dependency

import (
	"database/sql"
	"go-ddd-api-boilerplate/domain/repository"
	"go-ddd-api-boilerplate/infrastructure/persistence/postgres"
	r "go-ddd-api-boilerplate/infrastructure/persistence/redis"

	"github.com/go-redis/redis"
)

func NewGoodsRepository(db interface{}) repository.GoodsRepository {

	switch connection := db.(type) {
	case *sql.DB:
		return postgres.NewGoodsRepository(connection)

	}
	return nil
}

func NewGoodsCacheRepository(db interface{}) repository.GoodsCacheRepository {

	switch connection := db.(type) {
	case *redis.Client:
		return r.NewGoodsCacheRepository(connection)

	}
	return nil
}
