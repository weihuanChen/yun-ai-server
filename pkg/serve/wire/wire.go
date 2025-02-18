//go:build wireinject

package wire

import (
	"github.com/google/wire"
	"gorm.io/gorm"
	"yinglian.com/yun-ai-server/pkg/serve/controller/account"
	"yinglian.com/yun-ai-server/pkg/serve/mapper"
	"yinglian.com/yun-ai-server/pkg/serve/service"
)

var AccountSet = wire.NewSet(
	mapper.NewAccountMapperImpl,
	wire.Bind(new(mapper.AccountMapper), new(*mapper.AccountMapperImpl)),
	service.NewAccountServiceImpl,
	wire.Bind(new(service.AccountService), new(*service.AccountServiceImpl)),
	account.NewAccountController,
)

func InitializeAccountController(db *gorm.DB) (*account.AccountController, error) {
	wire.Build(AccountSet)
	return &account.AccountController{}, nil
}
