package repository

import (
	"studysystem/models"
	"studysystem/sql"
)

func Get_route(line_type int) *models.Study_route {
	db := sql.GetMysqlDB()
	var route models.Study_route
	db.Where("line_type = ?", line_type).Find(&route)
	return &route
}
