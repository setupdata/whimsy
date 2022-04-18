package config

type Server struct {
	System System `json:"system" yaml:"system"`
	Log    Log    `json:"log" yaml:"log"`
	Mysql  Mysql  `json:"mysql" yaml:"mysql"`
	Redis  Redis  `json:"redis" yaml:"redis"`
	JWT    JWT    `json:"jwt" yaml:"jwt"`
	EMAIL  EMAIL  `json:"email" yaml:"email"`
}
