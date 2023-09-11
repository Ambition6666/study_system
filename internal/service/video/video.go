package video

import (
	"studysystem/internal/repository"
	"studysystem/models"
	"studysystem/vo"
)

// 添加视频
func AddVideo(u *vo.Add_video_resquest) (int, string) {
	v := new(models.Video)
	v.Cover_url = u.Cover_url
	v.Description = u.Description
	v.Play_url = u.Play_url
	v.Title = u.Title
	v.Line_type = u.Line_type
	v.Mini_type = u.MiniType
	v.Stage_type = u.StageType
	repository.AddVideo(v)
	return 200, "添加成功"
}

// 删除视频
func DeleteVideo(vid uint) (int, string) {
	repository.DeleteVideo(vid)
	return 200, "删除成功"
}

// 获取视频列表
func GetVideoList(r int, m int, s int) (int, models.Study_route, []models.Video) {
	vlist := repository.GetVideoList(r, m, s)
	route := repository.Get_route(r)
	return 200, *route, vlist
}
