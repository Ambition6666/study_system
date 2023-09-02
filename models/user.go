package models

import (
	"studysystem/config"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID               int64  `json:"uid" gorm:"primarykey"`
	NickName         string `json:"nickname"`
	PassWord         string `json:"password"`
	Email            string `json:"email"`
	Avatar           string `json:"avatar"`
	IndividualResume string `json:"individual_resume"`
}

func (u *User) AfterFind(tx *gorm.DB) (err error) {
	dir := "http://" + config.ServerHost + ":" + config.ServerPort + "/identified/avatar/"
	if u.Avatar != dir {
		u.Avatar = dir
	}
	return nil
}
