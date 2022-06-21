package controller

import (
	"customer/handler"
	"customer/model"
	"customer/repository"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CustomerController struct{}

func NewCustomerController() *CustomerController {
	return &CustomerController{}
}

func (c *CustomerController) ReadCustomer(ctx *gin.Context) {
	id := ctx.Query("id")
	CustomerConfig := repository.CustomerRepository{}

	if result, err := CustomerConfig.ReadCustomer(id); err != nil {
		handler.NewError(ctx, http.StatusNotFound, err)
		return
	} else {
		ctx.JSON(http.StatusOK, model.Controller{
			Message: "successful",
			Data:    result,
		})
	}
}

func (c *CustomerController) ReadAllCustomer(ctx *gin.Context) {
	CustomerConfig := repository.CustomerRepository{}

	if result, err := CustomerConfig.ReadAllCustomer(); err != nil {
		handler.NewError(ctx, http.StatusNotFound, err)
		return
	} else {
		ctx.JSON(http.StatusOK, model.Controller{
			Message: "successful",
			Data:    result,
		})
	}
}

func (c *CustomerController) CreateCustomer(ctx *gin.Context) {
	var CustomerData model.CustomerModel
	if err := ctx.ShouldBindJSON(&CustomerData); err != nil {
		handler.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	CustomerConfig := repository.CustomerRepository{}

	if err := CustomerConfig.CreateCustomer(CustomerData); err != nil {
		handler.NewError(ctx, http.StatusBadRequest, err)
		return
	} else {
		ctx.JSON(http.StatusOK, model.Controller{Message: "successful", Data: nil})
	}
}

func (c *CustomerController) UpdateCustomer(ctx *gin.Context) {
	var updateCustomer model.CustomerModel
	if err := ctx.ShouldBindJSON(&updateCustomer); err != nil {
		handler.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	CustomerConfig := repository.CustomerRepository{}

	if err := CustomerConfig.UpdateCustomer(updateCustomer); err != nil {
		handler.NewError(ctx, http.StatusNotFound, err)
		return
	} else {
		ctx.JSON(http.StatusOK, model.Controller{
			Message: "successful",
			Data:    model.QueryEffected{Row: int64(updateCustomer.Id)},
		})
	}
}

func (c *CustomerController) DeleteCustomer(ctx *gin.Context) {
	idStr := ctx.Param("id")
	if idStr == "" {
		handler.NewError(ctx, http.StatusBadRequest, errors.New("id is undefine"))
	}
	CustomerConfig := repository.CustomerRepository{}

	if err := CustomerConfig.DeleteCustomer(idStr); err != nil {
		handler.NewError(ctx, http.StatusNotFound, err)
		return
	} else {
		ctx.JSON(http.StatusOK, model.Controller{
			Message: "successful",
			Data:    idStr,
		})
	}

}
