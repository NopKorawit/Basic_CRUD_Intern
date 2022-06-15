package controller

import (
	"customer/handler"
	"customer/model"
	"customer/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RoomController struct{}

func NewBookController() *RoomController {
	return &RoomController{}
}

func (r *RoomController) ReadRoom(ctx *gin.Context) {
	id := ctx.Query("id")
	RoomConfig := repository.RoomRepository{}

	if result, err := RoomConfig.ReadCustomer(id); err != nil {
		handler.NewError(ctx, http.StatusNotFound, err)
		return
	} else {
		ctx.JSON(http.StatusOK, model.Controller{
			Message: "successful",
			Data:    result,
		})
	}
}
