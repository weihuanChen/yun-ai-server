package mapper

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"yinglian.com/yun-ai-server/internal/global"
	"yinglian.com/yun-ai-server/internal/model"
	"yinglian.com/yun-ai-server/internal/utils"
)

type AccountMapper interface {
	CreateAccount(account *model.Account, c *gin.Context) (*model.Account, error)
	GetAccountByEmail(email string, c *gin.Context) (*model.Account, error)
}
type AccountMapperImpl struct {
	DB *gorm.DB
}

func NewAccountMapperImpl(db *gorm.DB) *AccountMapperImpl {
	return &AccountMapperImpl{DB: db}
}

func (mapper *AccountMapperImpl) GetAccountByEmail(email string, c *gin.Context) (*model.Account, error) {
	utils.BizLogger(c).Infof("根据邮箱查找账号: %v", email)
	var acc model.Account
	result := global.DB.Where("email = ?", email).First(&acc)
	if result.Error != nil {
		// 如果是记录未找到的错误，返回 nil 和 nil
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil // 或者返回特定的业务错误
		}
		// 如果是其他错误，返回错误
		return nil, result.Error
	}
	return &acc, nil
}

func (mapper *AccountMapperImpl) CreateAccount(account *model.Account, c *gin.Context) (*model.Account, error) {
	utils.BizLogger(c).Infof("创建新账户:%v", account)
	if err := mapper.DB.Create(account).Error; err != nil {
		return nil, err
	}
	return account, nil
}
