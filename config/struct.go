package config

// Database  -
type Database struct {
	Host       string `json:"host" yaml:"host"`
	User       string `json:"user" yaml:"user"`
	Password   string `json:"password" yaml:"password"`
	Database   string `json:"database" yaml:"database"`
	Collection string `json:"collection" yaml:"collection"`
	SSL        bool   `json:"ssl" yaml:"ssl"`
	AppName    string `json:"appname" yaml:"appname"`
	PoolSize   uint64 `json:"poolsize" yaml:"poolsize"`
	ReplicaSet string `json:"replicaset" yaml:"replicaset"`
}

// Logger -
type Logger struct {
	Level   string `json:"level" yaml:"level"`
	Webhook string `json:"webhook" yaml:"webhook"`
}

// Nats -
type Nats struct {
	Host     string `json:"host" yaml:"host"`
	User     string `json:"user" yaml:"user"`
	Password string `json:"password" yaml:"password"`
	Name     string `json:"name" yaml:"name"`
}

type Consul struct {
	Host string `json:"host" yaml:"host"`
	Acl  string `json:"acl" yaml:"acl"`
}

