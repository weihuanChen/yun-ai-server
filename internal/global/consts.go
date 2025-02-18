package global

// 缓存常量模块
const (
	SetCmd = "SET"
	ExCmd  = "EX"
	GetCmd = "GET"
)

// 账号模块常量
const (
	Header                  = "Authorization"
	BearerPrefix            = "Bearer "
	AccAuthTokenCachePrefix = "ACC_AUTH_TOKEN_CACHE_PREFIX:"
	AccAuthTokenCacheExpire = 82800 // 23 小时, 略小于 token 的实际有效期
	LocalsUserIdKey         = "userId"
)
