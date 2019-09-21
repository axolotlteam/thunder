package config

// Database  -
type Database struct {
	Host       string `json:"host" yaml:"host"`
	User       string `json:"user" yaml:"user"`
	Password   string `json:"password" yaml:"password"`
	Database   string `json:"database" yaml:"database"`
	Collection string `json:"collection" yaml:"collection"`
	SSL        bool   `json:"ssl" yaml:"ssl"`
}

// Logger -
type Logger struct {
	Level   string `json:"level" yaml:"level"`
	Webhook string `json:"webhook" yaml:"webhook"`
}
