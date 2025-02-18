package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"yinglian.com/yun-ai-server/internal/global"
	"yinglian.com/yun-ai-server/internal/model"
	"yinglian.com/yun-ai-server/internal/utils"
	"yinglian.com/yun-ai-server/pkg/serve/controller/dto"
	"yinglian.com/yun-ai-server/pkg/serve/mapper"
)

type AccountService interface {
	RegisterAccount(req *dto.RegisterAccReq, c *gin.Context) error
	LoginAcc(req *dto.LoginAccRequest, c *gin.Context) (string, error)
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

func (service *AccountServiceImpl) LoginAcc(req *dto.LoginAccRequest, c *gin.Context) (string, error) {
	// 查找用户
	user, err := service.AccountMapper.GetAccountByEmail(req.Email, c)
	if err != nil {
		return "", fmt.Errorf("服务器内部错误: %v", err)
	}
	if user == nil {
		return "", fmt.Errorf("用户不存在，请先注册")
	}
	// 校验密码
	if !checkPassword(user.Password, req.Password) {
		return "", fmt.Errorf("密码错误")
	}
	return getAuthToken(user.UserId)
}

func checkPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

// getAuthToken
func getAuthToken(userId int64) (string, error) {
	// 生成鉴权token
	accToken, err := utils.GenAccessToken(userId)
	if err != nil {
		return "", err
	}
	// 写入缓存 jwt无状态，只能等待时间过期后自动过期，加入缓存校验
	accKey := global.AccAuthTokenCachePrefix + strconv.FormatInt(userId, 10)
	conn := global.RDB.Get()
	defer conn.Close()
	// 储存过期时间 token
	_, err = conn.Do(global.SetCmd, accKey, accToken, global.ExCmd, global.AccAuthTokenCacheExpire)
	if err != nil {
		return "", err
	}
	return accToken, nil
}
