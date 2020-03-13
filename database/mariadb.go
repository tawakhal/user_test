package database

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql" //mysql driver
	cs "github.com/tawakhal/user_test/constanta"
	mdl "github.com/tawakhal/user_test/model"
)

type dbManipulation struct {
	db *sql.DB
}

// NewDatabaseManipulation create connection database and list api which connecting with database
func NewDatabaseManipulation(conf ConfDB) (DBManipulation, []mdl.Err) {
	const (
		parentFunc = "NewDatabaseManipulation"
	)
	var errs []mdl.Err
	schemaURL := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true",
		conf.User,
		conf.Pass,
		conf.Host,
		conf.Schema,
	)
	db, err := sql.Open(conf.Driver, schemaURL)
	if err != nil {
		errs = append(errs, mdl.Err{Parent: parentFunc, Function: "Sql Open", Code: cs.ErrCodeDatabase, Error: err})
		return nil, errs
	}
	db.SetMaxOpenConns(conf.MaxOpenConn)
	db.SetMaxIdleConns(conf.MaxIdleConn)
	db.SetConnMaxLifetime(time.Duration(conf.ConnMaxLifeTime) * time.Minute)
	return &dbManipulation{db: db}, nil
}
