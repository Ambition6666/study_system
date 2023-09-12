package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

// 问题
type Problem struct {
	gorm.Model
	VideoID     uint           `json:"video_id" gorm:"index"`
	Description string         `json:"description"`
	Options     pq.StringArray `json:"options" gorm:"type:text[]"`
	Answer      pq.Int32Array  `json:"-" gorm:"type:int[]"`
	ProblemType int            `json:"problem_type"`
	CodeID      int64          `json:"code_id"` //选择题id默认为0,如果是编程题,则它存的是问题id
}
