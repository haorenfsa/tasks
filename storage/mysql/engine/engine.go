package engine

import (
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"

	"github.com/jmoiron/sqlx"
)

// Default engine for mysql
type Default struct {
	*sqlx.DB
	mysql.Config
}

// DSN abbr. for Database Source Network
type DSN struct {
	Address  string
	UserName string
	Password string
	DBName   string
}

// NewDefault simple constructor
func NewDefault(dsn DSN) (*Default, error) {
	conf := mysql.NewConfig()
	conf.ParseTime = true
	conf.AllowCleartextPasswords = true

	conf.User = dsn.UserName
	conf.Addr = dsn.Address
	conf.Passwd = dsn.Password
	conf.DBName = dsn.DBName

	log.Print("full DSN: ", conf.FormatDSN())
	db, err := sqlx.Connect("mysql", conf.FormatDSN())
	if err != nil {
		return nil, fmt.Errorf("open dsn failed: %w", err)
	}
	return &Default{DB: db}, nil
}
