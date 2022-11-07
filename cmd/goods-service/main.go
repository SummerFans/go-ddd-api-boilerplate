package main

import (
	"context"
	"fmt"
	"go-ddd-api-boilerplate/application/handler"
	"go-ddd-api-boilerplate/dependency"
	"go-ddd-api-boilerplate/domain/interactor"
	"go-ddd-api-boilerplate/presenter"

	"github.com/sirupsen/logrus"
)

func main() {

	ctx := context.Background()

	log := logrus.New()

	db, err := dependency.NewPostgresConnection()
	if err != nil {
		fmt.Printf("%s", err.Error())
		return
	} else {
		defer dependency.Close(db)
	}

	goodsInteractor := interactor.NewGoodsInteractor(
		ctx,
		log,
		dependency.NewGoodsRepository(db),
		dependency.NewGoodsCacheRepository(db),
	)

	goodsHandler := handler.GoodsHandler{
		Handler: handler.Handler{
			Logger: log,
		},
		Interactor: goodsInteractor,
	}

	// build service

	// 构建服务
	httpServer := presenter.NewHttpServer(
		goodsHandler,
	)
	// run service
	httpServer.Run()
}
