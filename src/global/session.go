package global

import (
	"time"

	"github.com/dxvgef/sessions"
)

// session引擎
var Session *sessions.Engine

// SetSessions 设置session引擎
func SetSessions() (err error) {
	// 创建session管理器
	Session, err = sessions.NewEngine(&sessions.Config{
		CookieName:                LocalConfig.Session.CookieName,                            // cookie中的sessionID名称
		HttpOnly:                  LocalConfig.Session.HTTPOnly,                              // 仅允许HTTP读取，js无法读取
		Domain:                    "",                                                        // 作用域名，留空则自动获取当前域名
		Path:                      "/",                                                       // 作用路径
		MaxAge:                    LocalConfig.Session.MaxAge * 60,                           // 最大生命周期（秒）
		IdleTime:                  time.Duration(LocalConfig.Session.IdleTime) * time.Minute, // 空闲超时时间
		Secure:                    LocalConfig.Session.Secure,                                // 启用HTTPS
		DisableAutoUpdateIdleTime: false,                                                     // 禁止自动更新空闲时间
		RedisAddr:                 LocalConfig.Session.RedisAddr,                             // Redis地址
		RedisDB:                   LocalConfig.Session.RedisDB,                               // Redis数据库
		RedisPassword:             "",                                                        // Redis密码
		RedisKeyPrefix:            LocalConfig.Session.RedisKeyPrefix,                        // Redis中的键名前缀，必须
		Key:                       LocalConfig.Session.Key,                                   // 用于加密sessionID的密钥，密钥的长度16,24,32对应AES-128,AES-192,AES-256算法
	})
	return
}
