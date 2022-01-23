package config

type System struct {
	// 环境值 线上环境：public 开发环境：develop 开发环境日治为debug模式
	env  string `json:"env"`
	Addr int    `json:"addr"`
}
