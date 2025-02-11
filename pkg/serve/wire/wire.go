//go:build wireinject

package wire

import (
	"github.com/google/wire"
	"gorm.io/gorm"
	"yinglian.com/yun-ai-server/pkg/serve/controller"
	"yinglian.com/yun-ai-server/pkg/serve/mapper"
	"yinglian.com/yun-ai-server/pkg/serve/service"
)

var AccountSet = wire.NewSet(
	mapper.NewAccountMapperImpl,
	wire.Bind(new(mapper.AccountMapper), new(*mapper.AccountMapperImpl)),
	service.NewAccountServiceImpl,
	wire.Bind(new(service.AccountService), new(*service.AccountServiceImpl)),
	controller.NewAccountController,
)

func InitializeAccountController(db *gorm.DB) (*controller.AccountController, error) {
	wire.Build(AccountSet)
	return &controller.AccountController{}, nil
}
