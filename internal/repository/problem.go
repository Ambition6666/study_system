package repository

import (
	"studysystem/models"
	"studysystem/sql"
)

// 添加问题
func Add_problem(p *models.Problem) {
	db := sql.GetPgsql()
	db.Create(p)
}

// 删除问题
func Delete_problem(qid uint) {
	db := sql.GetPgsql()
	db.Where("id = ?", qid).Delete(&models.Problem{})
}

// 根据问题id查找问题
func Get_problem(qid uint) *models.Problem {
	v := new(models.Problem)
	db := sql.GetPgsql()
	db.Where("id = ? ", qid).Find(v)
	return v
}

// 根据视频id获取问题列表
func Get_problem_list(vid uint) []models.Problem {
	db := sql.GetPgsql()
	plist := make([]models.Problem, 0)
	db.Where("video_id = ?", vid).First(&plist)
	return plist
}
