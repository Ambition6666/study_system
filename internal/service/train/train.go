package train

import (
	"studysystem/internal/repository"
	"studysystem/internal/service/pool"
	"studysystem/models"
)

// 获取题目
func Get_problem_list(vid uint) (int, []models.Problem) {
	plist := repository.Get_problem_list(vid)
	return 200, plist
}

// 判断题目是否正确
func JudgeProblem(val ...any) any {
	intf := val[0].([]any)
	alist := intf[0].([]int32)
	qid := intf[1].(uint)
	uid := intf[2].(int64)
	istrue := true
	v := new(models.CommitRecord)
	p := repository.Get_problem(qid)
	myanswer := make([]int32, len(p.Options))
	for _, val := range p.Answer {
		myanswer[val] = 1
	}
	if len(alist) == 0 {
		v.IsTrue = false
		v.MyAnswer = myanswer
		v.Qid = qid
		v.Uid = uid
		repository.CreateCommitRecord(v)
		return nil
	}
	for i := 0; i < len(alist); i++ {
		if myanswer[alist[i]] == 1 {
			myanswer[alist[i]] = 2
		} else {
			istrue = false
			myanswer[alist[i]] = 3
		}
	}
	v.IsTrue = istrue
	v.MyAnswer = myanswer
	v.Qid = qid
	v.Uid = uid
	repository.CreateCommitRecord(v)
	return nil
}

// 提交题目
func CommitAnswer(alist []int32, qid uint, uid int64) (int, string) {
	pool.P.EmptyChan <- pool.NewTask(JudgeProblem, alist, qid, uid)
	return 200, "提交成功"
}

// 获取做题情况
func GetRecord(q string, uid int64) (int, *models.CommitRecord) {
	v := repository.GetRecord(q, uid)
	return 200, v
}
