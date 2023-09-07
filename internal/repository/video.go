package repository

import (
	"studysystem/models"
	"studysystem/sql"
)

// 添加视频
func AddVideo(v *models.Video) {
	pdb := sql.GetMysqlDB()
	pdb.Create(v)
}

// 删除视频
func DeleteVideo(id uint) {
	pdb := sql.GetMysqlDB()
	pdb.Where("id = ?", id).Delete(&models.Video{})
}
