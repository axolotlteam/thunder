package mssql

import (
	"database/sql"
	"strings"

	"github.com/axolotlteam/thunder/logger"

	// mssql
	_ "github.com/denisenkom/go-mssqldb"
)

var m *sql.DB

// M -
func M() (*sql.DB, error) {
	err := m.Ping()
	if err != nil {
		logger.WithError(err).Error("mssql connect failed")
		return nil, err
	}
	return m, nil
}

// Con -
func Con(c Config) error {
	if err := con(c); err != nil {
		return err
	}
	return nil
}

func con(c Config) error {
	var conf []string
	conf = append(conf, "server="+c.Host)
	conf = append(conf, "user id="+c.User)
	conf = append(conf, "password="+c.Password)
	conf = append(conf, "port="+c.Port)
	conf = append(conf, "encrypt=disable")
	conf = append(conf, "database="+c.Database)
	conn := strings.Join(conf, ";")
	db, err := sql.Open("sqlserver", conn)
	if err != nil {
		return err
	}

	db.SetMaxOpenConns(c.OpenConns)
	db.SetMaxIdleConns(c.IdleConns)
	db.SetConnMaxLifetime(c.MaxLifetime)

	m = db

	return nil
}
