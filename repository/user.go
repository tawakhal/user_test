package repository

import (
	"database/sql"

	"github.com/rs/zerolog/log"
	cs "github.com/tawakhal/user_test/constanta"
)

/*
CREATE TABLE mst_user (
	user_id BIGINT(20) NOT NULL AUTO_INCREMENT,
	user_name VARCHAR(32) NOT NULL,
	password VARCHAR(256) NOT NULL,
	email VARCHAR(64) NOT NULL,
	avatar VARCHAR(50) NOT NULL,
	status INT(12) NOT NULL,
	PRIMARY KEY (`user_id`))
	ENGINE = InnoDB;
*/
const (
	insertUser = `INSERT INTO mst_user (user_name,password,email,avatar,status) VALUES (?,?,?,?,?)`
)

// User is model table of user
type User struct {
	UserID   int32
	UserName string
	Password string
	Email    string
	Avatar   string
	Status   cs.Status
}

// Insert is insert to mst_user
func (user *User) Insert() (int64, error) {
	const (
		msgFunc = `User Insert`
	)

	result, err := db.Exec(insertUser, user.UserName, user.Password, user.Email, user.Avatar, user.Status)
	if err != nil && err != sql.ErrNoRows {
		log.Err(err).Msg(msgFunc)
		return 0, err
	}
	return checkResult(result), nil
}
