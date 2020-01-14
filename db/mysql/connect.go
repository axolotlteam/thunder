package mysql

import (
	"database/sql"
	"fmt"

	"github.com/axolotlteam/thunder/logger"
	_ "github.com/go-sql-driver/mysql"
)

var m *sql.DB

// M -
func M() (*sql.DB, error) {
	err := m.Ping()
	if err != nil {
		logger.WithError(err).Error("mysql connect failed")
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
	connStr := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true&loc=Local&timeout=30s",
		c.User,
		c.Password,
		c.Host,
		c.Database,
	)

	db, err := sql.Open("mysql", connStr)
	if err != nil {
		return err
	}

	db.SetMaxOpenConns(c.OpenConns)
	db.SetMaxIdleConns(c.IdleConns)
	db.SetConnMaxLifetime(c.MaxLifetime)

	m = db

	return nil
}
