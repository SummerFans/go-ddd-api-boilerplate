package presenter

import (
	"go-ddd-api-boilerplate/application/handler"

	"github.com/gin-gonic/gin"
)

func goodsRouter(router *gin.Engine, handler handler.GoodsHandler) {
	r := router.Group("/api/goods")
	{
		// 获取备货申请列表
		r.GET("/:id", func(c *gin.Context) {
			handler.GetGoods(c)
		})
	}
}
