package config

type System struct {
	// 环境值 线上环境：public 开发环境：develop 开发环境日治为debug模式
	Env  string `json:"env"`
	Addr string `json:"addr"`
}
