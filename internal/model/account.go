package model

import "yinglian.com/yun-ai-server/internal/model/base"

type Account struct {
	base.BasedModel
	UserId   int64  `gorm:"type:bigint;not null;uniqueIndex" json:"user_id"`        // 用户 ID
	Email    string `gorm:"type:varchar(64);default:null;uniqueIndex" json:"email"` // 邮箱
	Password string `gorm:"type:varchar(255);not null" json:"password"`             // 加密密码
	Nickname string `gorm:"type:varchar(255);default:null" json:"nickname"`         // 昵称
}

func (Account) TableName() string {
	return "account"
}
