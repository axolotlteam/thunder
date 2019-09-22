package mongo

import (
	"crypto/tls"
	"net"
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
func Con(c Config) error {
	if err := con(c); err != nil {
		return err
	}
	return nil
}

func con(c Config) error {
	mconfig := &mgo.DialInfo{
		Addrs:         strings.Split(c.Host, ","),
		Username:      c.User,
		Password:      c.Password,
		Database:      c.Database,
		Timeout:       5 * time.Second,
		PoolLimit:     30,
		PoolTimeout:   120 * time.Second,
		AppName:       c.AppName,
		MaxIdleTimeMS: 10000,
	}

	if c.SSL {
		mconfig.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
			return tls.Dial("tcp", addr.String(), &tls.Config{
				InsecureSkipVerify: true,
			})
		}
	}

	s, err := mgo.DialWithInfo(mconfig)
	if err != nil {
		return err
	}
	m = s
	return nil
}
