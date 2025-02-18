package account

import (
	"github.com/gin-gonic/gin"
	"net/http"
	biz_err "yinglian.com/yun-ai-server/internal/error"
	"yinglian.com/yun-ai-server/internal/utils"
	"yinglian.com/yun-ai-server/pkg/serve/controller/dto"
	"yinglian.com/yun-ai-server/pkg/serve/service"
	"yinglian.com/yun-ai-server/pkg/vo"
	"yinglian.com/yun-ai-server/pkg/vo/account"
)

type AccountController struct {
	AccountService service.AccountService
}

func NewAccountController(userService service.AccountService) *AccountController {
	return &AccountController{AccountService: userService}
}
func (ctl *AccountController) RegisterAcc(c *gin.Context) {
	req := new(dto.RegisterAccReq)
	// 解析请求体到 req
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusOK, vo.Fail(biz_err.New(biz_err.BadRequest, err.Error()), nil, c))
		return
	}
	// 基本参数验证
	errors := utils.Validator(*req)
	if errors != nil {
		c.JSON(http.StatusOK,
			vo.Fail(biz_err.New(biz_err.BadRequest), errors, c))
		return
	}

	// 注册
	err := ctl.AccountService.RegisterAccount(req, c)
	if err != nil {
		c.JSON(http.StatusOK, vo.Fail(biz_err.New(biz_err.ServerError, err.Error()), nil, c))
		return
	}

	// 返回成功响应
	c.JSON(http.StatusOK, vo.Success(nil, "注册成功!", c))

}

// 登陆
func (ctl *AccountController) Login(c *gin.Context) {
	req := new(dto.LoginAccRequest)
	// 解析请求体
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusOK, vo.Fail(biz_err.New(biz_err.BadRequest, err.Error()), nil, c))
		return
	}
	// 基本参数验证
	errors := utils.Validator(*req)
	if errors != nil {
		c.JSON(http.StatusOK, vo.Fail(biz_err.New(biz_err.BadRequest), errors, c))
		return
	}
	// 调用loginAcc函数进行验证
	aToken, err := ctl.AccountService.LoginAcc(req, c)
	if err != nil {
		c.JSON(http.StatusOK, vo.Fail(biz_err.New(biz_err.ServerError, err.Error()), nil, c))
		return
	}
	c.JSON(http.StatusOK, vo.Success(account.LoginVO{
		Token: aToken,
	}, "登录成功", c))
}
