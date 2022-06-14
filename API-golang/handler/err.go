package handler

import "github.com/gin-gonic/gin"

func NewError(ctx *gin.Context, status int, err error) {
	er := HTTPError{
		Code:    status,
		Message: err.Error(),
	}
	ctx.JSON(status, er)
}

type HTTPError struct {
	Code    int
	Message string
}

func PanicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}
