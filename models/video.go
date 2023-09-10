package models

import "gorm.io/gorm"

type Video struct {
	gorm.Model
	Play_url    string  `json:"play_url"`
	Cover_url   string  `json:"cover_url"`
	Description string  `json:"description"`
	Title       string  `json:"title" gorm:"type:varchar(255);index"`
	Type        float64 `json:"type"`
}
type Study_route struct {
	gorm.Model  `json:"-"`
	Line_type   int    `json:"line_type" gorm:"index"` //根据该类型去redis查找video列表
	Description string `json:"description"`
}
