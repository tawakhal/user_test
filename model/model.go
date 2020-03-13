package model

import (
	"time"

	cs "github.com/tawakhal/user_test/constanta"
	uuid "github.com/tawakhal/user_test/extend/utils"
)

// User is struct of user
type User struct {
	UserID     uuid.UUID
	UserName   string
	UserSecret string
	Mobile     string
	Status     cs.Status
	CreatedBy  uuid.UUID
	CreatedAt  time.Time
	UpdatedBy  uuid.UUID
	UpdatedAt  time.Time
}

// Respnose is standart response for this
type Respnose struct {
	Data     string
	Code     cs.RespCode
	InfoCode string
	Errors   []Err
}

// Err is information of error
type Err struct {
	Parent   string
	Function string
	Code     cs.ErrCode
	Error    error
}
