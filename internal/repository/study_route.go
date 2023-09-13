package repository

import (
	"studysystem/models"
	"studysystem/sql"
)

// 获取学习路线
func Get_route(line_type int) *models.Study_route {
	db := sql.GetMysqlDB()
	r := new(models.Study_route)
	db.Where("line_type = ?", line_type).Find(r)
	return r
}

// 添加学习路线
func Add_route(v *models.Study_route) {
	db := sql.GetMysqlDB()
	db.Create(v)
}
