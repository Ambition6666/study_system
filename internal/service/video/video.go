package video

import (
	"studysystem/internal/repository"
	"studysystem/models"
	"studysystem/vo"
)

func AddVideo(u *vo.Add_video_resquest) (int, string) {
	v := new(models.Video)
	v.Cover_url = u.Cover_url
	v.Description = u.Description
	v.Play_url = u.Play_url
	v.Title = u.Title
	v.Type = u.Type
	repository.AddVideo(v)
	return 200, "添加成功"
}
