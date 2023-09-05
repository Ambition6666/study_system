package models

import (
	"studysystem/config"

	"gorm.io/gorm"
)

// role
// 1-->普通成员
// 2-->管理员
type User struct {
	gorm.Model       `json:"-"`
	ID               int64  `json:"uid" gorm:"primarykey"`
	NickName         string `json:"nickname"`
	PassWord         string `json:"-"`
	Email            string `json:"email"`
	Avatar           string `json:"avatar"`
	IndividualResume string `json:"individual_resume"`
	Role             int    `json:"role"`
}

func (u *User) AfterFind(tx *gorm.DB) (err error) {
	dir := "http://" + config.ServerHost + ":" + config.ServerPort + "/identified/avatar/"
	if u.Avatar != dir {
		u.Avatar = dir
	}
	return nil
}
