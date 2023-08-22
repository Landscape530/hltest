package controller

import "github.com/gin-gonic/gin"

// 项目初始化测试
func RegisterHltestInitRouter(r *gin.RouterGroup) {
	r.POST("/sys-init", AddInfo)
}
