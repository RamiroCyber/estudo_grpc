package database

import (
	"gopkg.in/mgo.v2"
)

type Connection interface {
	Close()
	DB() *mgo.Database
}

type conn struct {
	session  *mgo.Session
	database *mgo.Database
}

func (c *conn) DB() *mgo.Database {
	return c.database
}

func (c *conn) Close() {
	c.session.Close()
}

func NewConnection(cfg Config) (Connection, error) {
	session, err := mgo.Dial(cfg.Dsn())
	if err != nil {
		panic(err)
	}
	return &conn{session: session, database: session.DB(cfg.DbName())}, nil
}
