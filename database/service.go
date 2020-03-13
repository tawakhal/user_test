package database

import (
	mdl "github.com/tawakhal/user_test/model"
)

// DBManipulation is list interface of database
type DBManipulation interface {
	InsertUser(user mdl.User) []mdl.Err
}
