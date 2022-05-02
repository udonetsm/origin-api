package db

import (
	"database/sql"
	"origin-api/getconf"

	"github.com/udonetsm/help/helper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func sqlDb() *sql.DB {
	sqldb, err := sql.Open("pgx", getconf.Storeconf)
	helper.Errors(err, "sqlopen")
	sqldb.SetConnMaxIdleTime(2)
	sqldb.SetConnMaxLifetime(2)
	sqldb.SetMaxIdleConns(5)
	sqldb.SetMaxOpenConns(5)
	return sqldb
}

func gormdb(sqldb *sql.DB) *gorm.DB {
	gormdb, err := gorm.Open(postgres.New(postgres.Config{Conn: sqldb}), &gorm.Config{})
	helper.Errors(err, "gormopen")
	return gormdb
}
