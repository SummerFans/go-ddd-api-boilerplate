package presenter

import (
	"context"
	"flag"
	"fmt"
	"go-ddd-api-boilerplate/application/handler"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

type httpServer struct {
	router *gin.Engine
}

func (s *httpServer) Run() {

	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "gracefully waiting for live connections")
	flag.Parse()

	var port string
	if port = os.Getenv("PORT"); len(port) == 0 {
		port = "20100"
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", port),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      s.router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	go func() {
		if err := srv.Shutdown(ctx); err != nil {
			log.Println(err)
		}
	}()
	<-ctx.Done()

	log.Println("Shutting down")
	os.Exit(0)
}

func loadRouter(
	goodsHandler handler.GoodsHandler,
) *gin.Engine {

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery()) // catch panic	recover
	router.Use(gzip.Gzip(gzip.DefaultCompression))

	//goods router
	goodsRouter(router, goodsHandler)

	return router
}

func NewHttpServer(
	goodsHandler handler.GoodsHandler,
) Server {
	router := loadRouter(
		goodsHandler,
	)

	return &httpServer{router: router}
}
