package problem

import (
	"studysystem/internal/repository"
	"studysystem/models"
	"studysystem/vo"
)

// 添加问题
func Add_problem(v *vo.Add_problem_request) (int, string) {
	p := new(models.Problem)
	p.Description = v.Description
	p.Answer = v.Answer
	p.Options = v.Options
	p.VideoID = v.VideoID
	p.ProblemType = v.ProblemType
	repository.Add_problem(p)
	return 200, "添加问题成功"
}

// 删除问题
func Delete_problem(id uint) (int, string) {
	repository.Delete_problem(id)
	return 200, "删除问题成功"
}
