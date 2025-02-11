package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"yinglian.com/yun-ai-server/internal/model"
	"yinglian.com/yun-ai-server/internal/utils"
	"yinglian.com/yun-ai-server/pkg/serve/controller/account/dto"
	"yinglian.com/yun-ai-server/pkg/serve/mapper"
)

type AccountService interface {
	RegisterAccount(req *dto.RegisterAccReq, c *gin.Context) error
}
type AccountServiceImpl struct {
	AccountMapper mapper.AccountMapper
}

func NewAccountServiceImpl(mapper mapper.AccountMapper) *AccountServiceImpl {
	return &AccountServiceImpl{AccountMapper: mapper}
}

// 注册账号
func (service *AccountServiceImpl) RegisterAccount(req *dto.RegisterAccReq, c *gin.Context) error {
	// 校验重复注册
	exist, _ := checkAccExist(req, service, c)
	if exist {
		return fmt.Errorf("账号已注册")
	}
	// 密码加密
	newPassword, err := cryptoPsw(req.Password)
	if err != nil {
		return err
	}
	// 注册完成
	addAcc := &model.Account{

		UserId:   utils.GenNewSnowflakeId(),
		Email:    req.Email,
		Password: newPassword,
		Nickname: req.NikeName,
	}
	newAccount, err := service.AccountMapper.CreateAccount(addAcc, c)
	if err != nil || newAccount == nil {
		return fmt.Errorf("注册失败")
	}
	return nil
}
func cryptoPsw(originPsw string) (string, error) {
	cryptoPwd, err := utils.CryptPwd(originPsw)
	if err != nil {
		return "", err
	}
	if len(cryptoPwd) == 0 {
		return "", fmt.Errorf("加密失败")
	}
	return string(cryptoPwd), nil
}
func checkAccExist(req *dto.RegisterAccReq, service *AccountServiceImpl, c *gin.Context) (bool, error) {
	existAcc, err := service.AccountMapper.GetAccountByEmail(req.Email, c)
	if err != nil {
		utils.BizLogger(c).Errorf("校验重复注册时查询 DB 失败: %v", err)
		// 无法判断是否已经注册时, 不继续注册
		return true, err
	}
	return existAcc != nil, nil
}
