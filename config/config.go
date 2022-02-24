package config

type Server struct {
	System System `json:"system" yaml:"system"`
	Log    Log    `json:"log" yaml:"log"`
	Mysql  Mysql  `json:"mysql" yaml:"mysql"`
}
