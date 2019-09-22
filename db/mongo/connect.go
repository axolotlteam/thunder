package mongo

import (
	"strings"
	"time"

	"github.com/axolotlteam/thunder/logger"
	"github.com/globalsign/mgo"
)

var m *mgo.Session

// ErrorList =
var (
	ErrNotFound = mgo.ErrNotFound
)

// M -
func M() *mgo.Session {
	if err := m.Ping(); err != nil {
		logger.WithError(err).Error("mongo connect failed")
	}
	return m.Copy()
}

// Close - close the database connect
func Close() {
	m.Close()
}

// Con - connect the  Mongo database
func Con(host, user, password, dbname, appname string) error {
	if err := con(host, user, password, dbname, appname); err != nil {
		return err
	}
	return nil
}

func con(host, user, password, dbname, appname string) error {
	mconfig := &mgo.DialInfo{
		Addrs:         strings.Split(host, ","),
		Username:      user,
		Password:      password,
		Database:      dbname,
		Timeout:       5 * time.Second,
		PoolLimit:     30,
		PoolTimeout:   120 * time.Second,
		AppName:       appname,
		MaxIdleTimeMS: 10000,
	}

	s, err := mgo.DialWithInfo(mconfig)
	if err != nil {
		return err
	}
	m = s

	return nil
}

// DevCon - use to connect the localhost database
func DevCon() error {
	return con("localhost:27017", "", "", "")
}
