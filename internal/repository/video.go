package repository

import (
	"context"
	"strconv"
	"studysystem/config"
	"studysystem/models"
	"studysystem/sql"

	"github.com/redis/go-redis/v9"
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

// 存储redis视频列表
func AddVideoList(line_type int, vid uint, t float64) error {
	rdb := sql.GetRedis()
	err := rdb.ZAdd(context.Background(), "studyroute"+strconv.FormatInt(int64(line_type), 10), redis.Z{Score: t, Member: vid}).Err()
	return err
}

// 取出全部redis视频列表
func GetVideoList(line_type int) ([]string, error) {
	rdb := sql.GetRedis()
	return rdb.ZRange(context.Background(), "studyroute"+strconv.FormatInt(int64(line_type), 10), 0, -1).Result()
}

// 分页展示视频列表
func GetVideoList_limit(line_type int, offset int64) ([]string, error) {
	rdb := sql.GetRedis()
	vals, err := rdb.ZRange(context.Background(), "studyroute"+strconv.FormatInt(int64(line_type), 10), offset*int64(config.PageNum), offset*int64(config.PageNum)+int64(config.PageNum)).Result()
	if len(vals) == 0 {
		return rdb.ZRange(context.Background(), "studyroute"+strconv.FormatInt(int64(line_type), 10), offset*int64(config.PageNum), -1).Result()
	}
	return vals, err
}

// 根据分数查询
func GetVideoListByScore(line_type int, offset int64, score float64) ([]string, error) {
	rdb := sql.GetRedis()
	// 初始化查询条件， Offset和Count用于分页
	op := redis.ZRangeBy{
		Min:    strconv.FormatFloat(score, 'e', 1, 64), // 最小分数
		Max:    strconv.FormatFloat(score, 'e', 1, 64), // 最大分数
		Offset: offset,                                 // 类似sql的limit, 表示开始偏移量
		Count:  int64(config.PageNum),                  // 一次返回多少数据
	}

	return rdb.ZRangeByScore(context.Background(), "studyroute"+strconv.FormatInt(int64(line_type), 10), &op).Result()
}

// 提取视频
func GetVideo(vids []string) []models.Video {
	db := sql.GetMysqlDB()
	v := make([]models.Video, 0)
	db.Where("id in ?", vids).Find(&v)
	return v
}
