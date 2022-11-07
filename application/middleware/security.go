package middleware

import (
	"go-ddd-api-boilerplate/application/handler"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type securityMiddleware struct {
	handler.Handler
}

func (s *securityMiddleware) Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}

func NewSecurityMiddleware(log *logrus.Logger) *securityMiddleware {
	return &securityMiddleware{
		Handler: handler.Handler{
			Logger: log,
		},
	}
}
