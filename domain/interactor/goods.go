package interactor

import (
	"context"
	"go-ddd-api-boilerplate/domain/repository"

	"github.com/sirupsen/logrus"
)

type GoodsInteractor interface {
}

type goodsInteractor struct {
	ctx                  context.Context
	Logger               *logrus.Logger
	goodsRepository      repository.GoodsRepository
	goodsCacheRepository repository.GoodsCacheRepository
}

func NewGoodsInteractor(
	ctx context.Context,
	log *logrus.Logger,
	goodsRepository repository.GoodsRepository,
	goodsCacheRepository repository.GoodsCacheRepository,
) GoodsInteractor {
	return &goodsInteractor{
		ctx:                  ctx,
		Logger:               log,
		goodsRepository:      goodsRepository,
		goodsCacheRepository: goodsCacheRepository,
	}
}
