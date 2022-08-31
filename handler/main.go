package handler

import (
	"context"
	"errors"
	"net/http"

	"github.com/bankonly/goutils/response"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type HandlerFunc struct {
	Error      error
	Message    string
	ErrorCode  string
	StatusCode int
	Data       interface{}
}

type Options struct {
	UseSession bool
	DBClient   *mongo.Client
}

type HandlerContext struct {
	Context        *gin.Context
	SessionContext mongo.SessionContext
}

func GoHandler(handler func(HandlerContext), option Options) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		res := response.New(ctx)
		handlerOption := HandlerContext{Context: ctx}
		if option.UseSession {
			session, err := option.DBClient.StartSession()
			if err != nil {
				res.PanicError(http.StatusInternalServerError, response.HError{Message: err.Error(), ErrorCode: "5000"})
			}
			defer session.EndSession(context.Background())

			if _, errSession := session.WithTransaction(context.Background(), func(sessCtx mongo.SessionContext) (interface{}, error) {
				handlerOption.SessionContext = sessCtx
				handler(handlerOption)
				return nil, nil
			}); errSession != nil {
				session.AbortTransaction(context.Background())
				res.PanicError(http.StatusInternalServerError, response.HError{Message: errSession.Error(), ErrorCode: "5000"})
			}

		} else {
			handler(handlerOption)
		}
	}
}

func Error(statusCode int, errorMessage string, errorCode string) HandlerFunc {
	return HandlerFunc{Error: errors.New(errorMessage), ErrorCode: errorCode, StatusCode: statusCode}
}
