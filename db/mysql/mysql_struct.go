package mysql

import "time"

// Config -
type Config struct {
	Host        string
	User        string
	Password    string
	Database    string
	AppName     string
	OpenConns   int
	IdleConns   int
	MaxLifetime time.Duration
}
