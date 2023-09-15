package repository

import (
	"context"
	"strconv"
	"studysystem/logs"
	"studysystem/models"
	"studysystem/sql"
)

// 添加问题
func Add_problem(p *models.Problem) {
	db := sql.GetPgsql()
	db.Create(p)
	logs.SugarLogger.Info(p.ID)
	v := SearchVideoByID(p.VideoID)
	Add_stage_problem(v.Line_type, v.Stage_type, p.ProblemType, p.ID)
}

// 添加阶段题
func Add_stage_problem(l int, s int, p int, q uint) {
	rdb := sql.GetRedis()
	err := rdb.SAdd(context.Background(), "l"+strconv.Itoa(l)+"s"+strconv.Itoa(s)+"p"+strconv.Itoa(p), q).Err()
	if err != nil {
		logs.SugarLogger.Debugf("存问题:%v", err)
	}
}

// 随机从题目里抽取题
func Get__rand_problem(l int, s int, p int, num int64) []string {
	rdb := sql.GetRedis()
	res, err := rdb.SRandMemberN(context.Background(), "l"+strconv.Itoa(l)+"s"+strconv.Itoa(s)+"p"+strconv.Itoa(p), num).Result()
	if err != nil {
		logs.SugarLogger.Debugln("获取题目:", err)
		return res
	}
	return res
}

// 获取题目
func Get__rand_problem_list(qlist []string) []models.Problem {
	pdb := sql.GetPgsql()
	v := make([]models.Problem, 0)
	pdb.Where("id in ? ", qlist).Find(&v)
	return v
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
	db.Where("video_id = ?", vid).Find(&plist)
	return plist
}
