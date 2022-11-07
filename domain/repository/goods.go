package repository

import "go-ddd-api-boilerplate/domain/entity"

type GoodsRepository interface {
	Save(goods *entity.Goods) (int64, error)
}

type GoodsCacheRepository interface {
	SaveCache(goods *entity.Goods) (int64, error)
}
