package config

type Log struct {
	Path         string `json:"path"`
	Name         string `json:"name"`
	Type         string `json:"type"`
	Format       string `json:"format"`
	Level        string `json:"level"`
	MaxAge       int    `json:"maxAge"`       // 文件最大保存时间，单位为分钟
	RotationTime int    `json:"rotationTime"` // 日志切割时间间隔，单位为分钟
}
