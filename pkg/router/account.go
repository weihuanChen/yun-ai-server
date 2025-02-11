package router

import (
	"github.com/gin-gonic/gin"
	"yinglian.com/yun-ai-server/pkg/serve/controller"
)

func accountRouter(ctl *controller.AccountController, rg *gin.RouterGroup) {
	rg.POST("/register", ctl.RegisterAcc)
}
