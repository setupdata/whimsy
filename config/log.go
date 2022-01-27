package config

type Log struct {
	Path   string `json:"path"`
	Name   string `json:"name"`
	Format string `json:"format"`
	Level  string `json:"level"`
}
