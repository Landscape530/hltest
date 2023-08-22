package httper

import (
	"awesomeProject/server/base/mistake"
	"github.com/gin-gonic/gin"
	"net/http"
)

func BindAndCheck(ctx *gin.Context, data interface{}) bool {
	//// 统一设置返回json格式， 因为gin在Bind发生异常的时候回自动加上header和400 所以需要提前设置json格式，不然返回的是text
	//ctx.Writer.Header().Add("Content-Type", "application/json; charset=utf-8")
	//
	//// 参数映射
	//if err := ctx.Bind(data); err != nil {
	//	// 当入参格式不正确的会出现，比如int传递为string
	//	HandleResponse(ctx, mistake.NewReqErr("Request Bind Error"), nil)
	//	return true
	//}
	//
	//// 去除结构体内部数据前后的空格
	//str.TrimStruct(data)

	return true
}

// HandleResponse 统一处理异常，统一处理日志，统一处理返回
func HandleResponse(c *gin.Context, err error, data interface{}) {
	// 如果没有错误，就是正常请求
	if err == nil {
		c.JSON(http.StatusOK, SucceedRespBodyAndData("操作成功", DumpContent(data)))
		return
	}

	// 处理请求参数异常
	if reqErr, ok := err.(*mistake.ReqErr); ok {
		c.JSON(http.StatusBadRequest, ErrRespBody(reqErr.Message))
		return
	}

	// 处理服务层异常
	if serviceErr, ok := err.(*mistake.ServiceErr); ok {
		//logger.Error(serviceErr.Message, "\n", serviceErr.Err, "\n", serviceErr.Stack)
		c.JSON(serviceErr.HTTPCode, ErrRespBody(serviceErr.Message))
		return
	}

	if unauthorizedErr, ok := err.(*mistake.StatusUnauthorizedErr); ok {
		c.JSON(http.StatusUnauthorized, ErrRespBody(unauthorizedErr.Error()))
		return
	}

	// 没有封装的错误
	//logger.Error(err)
	c.JSON(http.StatusInternalServerError, ErrRespBody("操作异常"))
	return
}
