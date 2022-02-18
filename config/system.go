package config

type System struct {
	// 环境值 线上环境：public 开发环境：develop 开发环境日治为debug模式
	Env    string `json:"env"`    // 环境
	Ip     string `json:"ip"`     // 监听ip 0.0.0.0
	Addr   string `json:"addr"`   // 端口
	DbType string `json:"dbType"` // 数据库类型
	GinMod string `json:"ginMod"` // gin模式 debug,release,test 默认为release
}
