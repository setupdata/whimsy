package config

type Server struct {
	System System `json:"system"`
	Log    Log    `json:"log"`
	Mysql  Mysql  `json:"mysql"`
}
