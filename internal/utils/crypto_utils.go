package utils

import (
	"github.com/bwmarrin/snowflake"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"time"
	"yinglian.com/yun-ai-server/internal/global"
)

func CryptPwd(password string) ([]byte, error) {
	pwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return pwd, err
}
func NewRand() int {
	// 创建一个新的随机数生成器 (随机种子), 并使用当前时间的纳秒数作为种子
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(900000) + 100000
}
func GenNewSnowflakeId() int64 {
	node, err := snowflake.NewNode(1)
	if err != nil {
		global.SysLog.Errorf("创建雪花节点失败: %v", err)
	}
	return node.Generate().Int64()
}
