package main

import (
	"customer/controller"
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	v1 := r.Group("api")
	{
		br := controller.NewCustomerController()

		c := v1.Group("/customer")
		{
			c.GET("/id", br.ReadCustomer)
			c.GET("", br.ReadAllCustomer)
			c.POST("", br.CreateCustomer)
			c.PUT("", br.UpdateCustomer)
			c.DELETE(":Id", br.DeleteCustomer)
		}
		// c2 := v1.Group("/sadsa")
		fmt.Println(c)
	}

	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8086")

}
