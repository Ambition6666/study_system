package studyroute

import (
	"studysystem/internal/repository"
	"studysystem/models"
)

func Add_study_route(dsc string, lt int) (int, string) {
	v := new(models.Study_route)
	v.Description = dsc
	v.Line_type = lt
	repository.Add_route(v)
	return 200, "创建学习路线成功"
}
