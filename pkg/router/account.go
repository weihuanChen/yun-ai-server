package router

import (
	"github.com/gin-gonic/gin"
	"yinglian.com/yun-ai-server/pkg/serve/controller/account"
)

func accountRouter(ctl *account.AccountController, rg *gin.RouterGroup) {
	rg.POST("/register", ctl.RegisterAcc)
	rg.POST("/login", ctl.Login)
}
