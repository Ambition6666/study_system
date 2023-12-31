package models

import (
	"github.com/lib/pq"
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

// 用户提交记录
type CommitRecord struct {
	gorm.Model
	Uid      int64         `json:"uid" gorm:"index"`
	Qid      uint          `json:"qid" gorm:"index"`
	MyAnswer pq.Int32Array `json:"my_answer" gorm:"type:int[]"`
	IsTrue   bool          `json:"is_true"`
}
