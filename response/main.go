package response

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type H struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

type HError struct {
	Message   string `json:"message"`
	ErrorCode string `json:"errorCode"`
}

type ResponseInterface interface {
	Success(H)
	PanicError(int, HError)
}

type Response struct {
	Ctx *gin.Context
}

func New(ctx *gin.Context) ResponseInterface {
	var response ResponseInterface = Response{Ctx: ctx}
	return response
}

func (r Response) Success(h H) {
	if h.Message == "" {
		h.Message = "Success"
	}
	r.Ctx.JSON(http.StatusOK, h)
}

func (r Response) PanicError(httpStatusCode int, h HError) {
	if h.Message == "" {
		h.Message = "Request failed"
	}
	if h.ErrorCode == "" {
		h.ErrorCode = "4000"
	}
	panic(fmt.Sprintf("%s-%s::%s", strconv.Itoa(httpStatusCode), h.ErrorCode, h.Message))
}
