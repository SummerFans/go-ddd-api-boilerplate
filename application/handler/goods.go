package handler

import (
	"go-ddd-api-boilerplate/domain/interactor"

	"github.com/gin-gonic/gin"
)

type GoodsHandler struct {
	Handler
	Interactor interactor.GoodsInteractor
}

func (h *GoodsHandler) GetGoods(c *gin.Context) {
	// ...
	successRes := SuccessResponse{
		Code: 0,
		Data: nil,
	}

	h.Respond(c, 200, successRes)
}
