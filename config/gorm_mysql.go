package config

type Mysql struct {
	Username     string `json:"username" yaml:"username"`
	Password     string `json:"password" yaml:"password"`
	Path         string `json:"path" yaml:"path"`
	Port         string `json:"port" yaml:"port"`
	Dbname       string `json:"dbname" yaml:"dbname"`
	Config       string `json:"config" yaml:"config"`
	MaxIdleConns int    `json:"maxIdleConns" yaml:"maxIdleConns"`
	MaxOpenConns int    `json:"maxOpenConns" yaml:"maxOpenConns"`
	LogMode      string `json:"logMode" yaml:"logMode"`
	LogLogrus    bool   `json:"logLogrus" yaml:"logLogrus"`
}

// Dsn 拼接mysql dsn 链接
func (m *Mysql) Dsn() string {
	//想要正确的处理 time.Time , 您需要带上 parseTime 参数,
	//要支持完整的 UTF-8 编码, 您需要将 charset=utf8 更改为 charset=utf8mb4, 详情见 gorm.cn。
	//user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local
	return m.Username + ":" + m.Password + "@tcp(" + m.Path + ":" + m.Port + ")/" + m.Dbname + "?" + m.Config
}
