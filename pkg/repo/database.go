package repo

import (
	"fmt"

	mongoadaptor "github.com/upper/db/v4/adapter/mongo"
	mysqladaptor "github.com/upper/db/v4/adapter/mysql"
	postgresqladaptor "github.com/upper/db/v4/adapter/postgresql"
	sqliteadaptor "github.com/upper/db/v4/adapter/sqlite"

	"github.com/open-privacy/opv/pkg/config"
	"github.com/upper/db/v4"
)

// NewDBSession creates a new DB Session
func NewDBSession() (session db.Session, err error) {
	var connectionURL db.ConnectionURL
	switch config.ENV.DBDriver {
	case "mysql":
		connectionURL, err = mysqladaptor.ParseURL(config.ENV.DBConnectionStr)
		if err != nil {
			return nil, err
		}
	case "sqlite":
		connectionURL, err = sqliteadaptor.ParseURL(config.ENV.DBConnectionStr)
		if err != nil {
			return nil, err
		}
	case "postgresql":
		connectionURL, err = postgresqladaptor.ParseURL(config.ENV.DBConnectionStr)
		if err != nil {
			return nil, err
		}
	case "mongo":
		connectionURL, err = mongoadaptor.ParseURL(config.ENV.DBConnectionStr)
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("invalid DBDriver %s", config.ENV.DBDriver)
	}

	return db.Open(config.ENV.DBDriver, connectionURL)
}

// MustNewDBSession creates a new DB Sesison or panic
func MustNewDBSession() db.Session {
	session, err := NewDBSession()
	if err != nil {
		panic(err)
	}
	return session
}
