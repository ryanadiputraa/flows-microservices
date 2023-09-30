package db

import (
	"github.com/jmoiron/sqlx"
	"github.com/ryanadiputraa/flows/flows-microservices/transaction/config"

	_ "github.com/lib/pq"
)

const (
	maxOpenConns    = 60
	connMaxLifeTime = 120
	maxIdleConns    = 30
	connMaxIdleTime = 20
)

func NewDB(conf *config.Config) (*sqlx.DB, error) {
	db, err := sqlx.Open(conf.DB.Driver, conf.DB.DSN)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(maxOpenConns)
	db.SetConnMaxLifetime(connMaxLifeTime)
	db.SetMaxIdleConns(maxIdleConns)
	db.SetConnMaxIdleTime(connMaxIdleTime)

	return db, err
}
