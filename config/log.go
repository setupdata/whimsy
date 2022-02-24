package config

type Log struct {
	Path         string `json:"path" yaml:"path"`
	Name         string `json:"name" yaml:"name"`
	Type         string `json:"type" yaml:"type"`
	Format       string `json:"format" yaml:"format"`
	Level        string `json:"level" yaml:"level"`
	MaxAge       int    `json:"maxAge" yaml:"maxAge"`             // 文件最大保存时间，单位为分钟
	RotationTime int    `json:"rotationTime" yaml:"rotationTime"` // 日志切割时间间隔，单位为分钟
}
