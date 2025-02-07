package db

import (
	"log"
	"yinglian.com/yun-ai-server/internal/global"
	"yinglian.com/yun-ai-server/internal/model"
)

func AutoMigrate() {
	// 确保数据库初始化
	if global.DB == nil {
		log.Fatal("数据库初始化失败，无法执行自动迁移...")
	}
	// 自动迁移模型
	err := global.DB.AutoMigrate(model.GetAllModels()...)
	if err != nil {
		log.Fatalf("数据库自动迁移失败: %v", err)
	}
	log.Println("数据库自动迁移成功")
}
