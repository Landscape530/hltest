package router

import (
	"awesomeProject/server/operation/hltest/controller"
	"github.com/gin-gonic/gin"
)

func InitRouter(port string) {
	r := gin.New()

	v1 := r.Group("/hl/api/v1")
	controller.RegisterHltestInitRouter(v1)

	if err := r.Run(":" + port); err != nil {
		panic(err)
	}
}
