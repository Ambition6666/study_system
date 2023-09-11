package repository

import (
	"studysystem/models"
	"studysystem/sql"
)

// 添加视频
func AddVideo(v *models.Video) {
	db := sql.GetMysqlDB()
	db.Create(v)
}

// 删除视频
func DeleteVideo(vid uint) {
	db := sql.GetMysqlDB()
	db.Where("id = ?", vid).Delete(&models.Video{})
}

// 通过视频id查找视频
func SearchVideoByID(vid uint) *models.Video {
	db := sql.GetMysqlDB()
	v := new(models.Video)
	db.Where("id = ?", vid).Find(v)
	return v
}

// 通过视频标题查找视频
func SearchVideoByTitle(t string) *models.Video {
	db := sql.GetMysqlDB()
	v := new(models.Video)
	db.Where("title = ?", t).Find(v)
	return v
}

// 提取视频
func GetVideoList(r int, m int, s int) []models.Video {
	db := sql.GetMysqlDB()
	v := make([]models.Video, 0)
	if m == 0 && s == 0 {
		db.Where("line_type = ?", r).Find(&v)
	} else if s == 0 {
		db.Where("mini_type = ?", m).Find(&v)
	} else if m == 0 {
		db.Where("stage_type = ?", s).Find(&v)
	} else {
		db.Where("line_type = ? and mini_type = ? and stage_type = ?", r, m, s).Find(&v)
	}
	return v
}
