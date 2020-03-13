package api

import (
	uuid "github.com/tawakhal/user_test/extend/utils"
	mdl "github.com/tawakhal/user_test/model"
)

// UserService is list API of user
type UserService interface {
	AddUser(user mdl.User) (resp mdl.Respnose, err error)
	GetUserByID(userID uuid.UUID) (mdl.User, error)
}
