package model

// GetAllModels 获取并注册所有模型
func GetAllModels() []interface{} {
	return []interface{}{
		&Account{},
	}
}
