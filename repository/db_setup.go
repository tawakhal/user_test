package repository

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql" // mysql
	"github.com/rs/zerolog/log"
	"github.com/tawakhal/user_test/extend/conf"
)

var db *sql.DB

// Setup is setup database connection
func Setup() error {
	var err error
	const (
		failedConnect = "Failed Connect To Database"
		dbError       = "Databases Error"
	)

	schemaURL := fmt.Sprintf(
		getDBConfig(),
		conf.DBConf.User,
		conf.DBConf.Password,
		conf.DBConf.Host+":"+strconv.Itoa(conf.DBConf.Port),
		conf.DBConf.DBName,
	)

	db, err = sql.Open(conf.DBConf.DBType, schemaURL)
	if err != nil {
		log.Err(err).Msg(failedConnect)
		return err
	}

	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(30 * time.Minute)

	return nil
}

func getDBConfig() string {
	const (
		mysql = "%s:%s@tcp(%s)/%s?parseTime=true&loc=Local"
	)
	switch conf.DBConf.DBType {
	case "mysql":
		return mysql
	default:
		return mysql
	}
}
