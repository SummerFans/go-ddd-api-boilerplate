package handler

import (
	"encoding/json"
	"net/http"
	"os"

	"go-ddd-api-boilerplate/domain/throwable"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ErrorResponse struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
	Error   error  `json:"-"`
}

type SuccessResponse struct {
	Code int64       `json:"code"`
	Data interface{} `json:"data"`
}

type Handler struct {
	Logger *logrus.Logger
}

func (h *Handler) Debug(format string, args ...interface{}) {

	if debug := os.Getenv("DEBUG"); debug == "true" {
		h.Logger.Debugf(format+"\n", args...)
	}
}

func (h *Handler) Error(c *gin.Context, err error) {
	var statusCode int
	var message string

	c.Writer.Header().Set("Content-Type", "application/json; charset=UTF-8")

	switch e := err.(type) {
	case *throwable.NotFound, throwable.NotFound:
		statusCode = http.StatusNotFound
		message = e.Error()
	case *throwable.Unauthorized, throwable.Unauthorized:
		statusCode = http.StatusUnauthorized
		message = e.Error()
	case *json.UnsupportedTypeError, *json.UnmarshalTypeError, *json.SyntaxError:
		statusCode = http.StatusBadRequest
		message = "Request body is invalid"
	default:
		statusCode = http.StatusInternalServerError
		message = e.Error()
	}

	h.Debug("%s", err.Error())

	errorResponse := &ErrorResponse{
		Code:    int64(statusCode),
		Message: message,
		Error:   err,
	}

	h.Respond(c, statusCode, errorResponse)
}

func (h *Handler) Respond(c *gin.Context, code int, src interface{}) {
	var body []byte
	var err error

	c.Writer.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if body, err = json.Marshal(src); err != nil {
		errorBody := "{\"status\": 500, \"message\": \"Something happened wrong during generating response\"}"
		c.Writer.WriteHeader(http.StatusInternalServerError)
		c.Writer.Write([]byte(errorBody))
		return
	}

	c.Writer.WriteHeader(code)
	c.Writer.Write(body)
}
