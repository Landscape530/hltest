package controller

import (
	"awesomeProject/server/base/auth"
	"awesomeProject/server/repo/hltest/val"
	"awesomeProject/util/tools/httper"
	"github.com/gin-gonic/gin"
	//"github.com/qbox/qmatrix/util/special/httper"
)

func AddInfo(ctx *gin.Context) {
	addInfoReq := &val.InitConnectionReq{}
	//if httper.Bind
	user, _ := ctx.Get(auth.CtxLoginUserKey)
	addInfoReq.User = user.(*auth.UserInfo)

	httper.HandleResponse(ctx, nil, map[string]interface{}{"id": 1})
}
