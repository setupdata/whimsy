package config

type EMAIL struct {
	Host         string `json:"host" yaml:"host"`
	From         string `json:"from" yaml:"from"`
	SmtpPassword string `json:"smtpPassword" yaml:"smtpPassword"`
}
