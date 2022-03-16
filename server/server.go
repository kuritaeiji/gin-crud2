package server

import (
	"gin-crud2/controller"

	"github.com/gin-gonic/gin"
)

func Init() {
	r := gin.Default()

	user := r.Group("/users")
	{
		ctr := controller.NewUserController()
		user.GET("", ctr.Index)
		user.GET("/:id", ctr.Show)
		user.POST("", ctr.Create)
	}

	r.Run()
}
