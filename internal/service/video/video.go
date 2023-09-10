package video

import (
	"fmt"
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
	v.Type = u.Type
	repository.AddVideo(v)
	return 200, "添加成功"
}

// 删除视频
func DeleteVideo(vid uint) (int, string) {
	repository.DeleteVideo(vid)
	return 200, "删除成功"
}

// 获取视频列表
func GetVideoList(line_type int, t float64, offset int64) (int, models.Study_route, []models.Video) {
	list, err := repository.GetVideoListByScore(line_type, offset, t)
	if err != nil {
		fmt.Println(err)
		return 500, models.Study_route{}, nil
	}
	route := repository.Get_route(line_type)
	vlist := repository.GetVideo(list)
	return 200, *route, vlist
}
