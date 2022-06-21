package main

import (
	"customer/controller"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	v1 := r.Group("api")
	{
		nc := controller.NewCustomerController()

		c := v1.Group("/customer")
		{
			c.GET("/id", nc.ReadCustomer)
			c.GET("", nc.ReadAllCustomer)
			c.POST("", nc.CreateCustomer)
			c.PUT("", nc.UpdateCustomer)
			c.DELETE(":Id", nc.DeleteCustomer)
		}

		rb := controller.NewBookController()

		r := v1.Group("/room")

		{
			r.GET("/id", rb.ReadRoom)
			r.GET("", rb.ReadAllRoom)
			// r.POST("", rb.CreateCustomer)
			// r.PUT("", rb.UpdateCustomer)
			// r.DELETE(":Id", rb.DeleteCustomer)
		}
	}

	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8086")

}
