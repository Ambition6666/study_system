package train

import (
	"context"
	"encoding/json"
	"studysystem/internal/repository"
	"studysystem/internal/service/pool"
	"studysystem/logs"
	"studysystem/models"
	"studysystem/vo"
	"sync"

	rpc "studysystem/clients"

	pri "studysystem/api/proto/private"
)

// 获取题目
func Get_problem_list(vid uint) (int, []models.Problem) {
	plist := repository.Get_problem_list(vid)
	return 200, plist
}

// 判断选择题目是否正确
func JudgeProblem(val ...any) []any {
	intf := val[0].([]any)
	alist := intf[0].([]int32)
	qid := intf[1].(uint)
	uid := intf[2].(int64)
	//res := intf[3].(chan any)
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
		n := new(vo.Test_res)
		n.Istrue = false
		n.QID = qid
		return []any{n, nil}
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
	n := new(vo.Test_res)
	n.Istrue = false
	n.QID = qid
	repository.CreateCommitRecord(v)
	return []any{n, nil}
}

// 获取oj题目信息

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

// 获取测试
func Get_test(l int, s int) (int, []models.Problem) {
	v := repository.Get__rand_problem_list(repository.Get__rand_problem(l, s, 1, 10))
	v = append(v, repository.Get__rand_problem_list(repository.Get__rand_problem(l, s, 2, 3))...)
	return 200, v
}

// 提交测试
func Commit_Test_answer(uid int64, list []json.RawMessage) (int, *vo.Test_res_s) {
	testres := new(vo.Test_res_s)
	res := make(chan any, 10)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		data := new(vo.Commit_answer_resquest)
		json.Unmarshal(list[i], data)
		t := pool.NewTask(JudgeProblem, data.Answer, data.Qid, uid)
		pool.P.EmptyChan <- t
		go func() {
			defer wg.Done()
			res <- t.Result
		}()
	}
	wg.Wait()
	close(res)
	//判断题目
	for val := range res {
		v, ok := val.(vo.Test_res)
		if !ok {
			logs.SugarLogger.Debugln("接口失败")
			return 500, nil
		}
		if v.Istrue {
			testres.Score += 4
			testres.Res = append(testres.Res, v)
		}
	}
	for i := 10; i < 13; i++ {
		// wg.Add(1)
		data := new(vo.Commit_code)
		err := json.Unmarshal(list[i], data)
		if err != nil {
			logs.SugarLogger.Debugln("解析消息错误", err)
		}
		q := repository.Get_problem(data.QID)
		res, err := rpc.ProCli.Judge(context.Background(), &pri.JudgeRequest{
			ProblemID: q.CodeID,
			Code:      []byte(data.Code),
			LangID:    data.LanguageID,
		})
		if err != nil {
			logs.SugarLogger.Debugln("判题错误:", err)
			return 500, nil
		}
		res1, err := rpc.ProCli.GetResult(context.Background(), &pri.GetResultRequest{
			JudgeID: res.JudgeID,
		})
		if err != nil {
			logs.SugarLogger.Debugln("判题错误:", err)
			return 500, nil
		}
		if res1.Result.Status == 10 {
			v := vo.Test_res{
				QID:    data.QID,
				Istrue: true,
			}
			testres.Score += 20
			testres.Res = append(testres.Res, v)
		} else {
			v := vo.Test_res{
				QID:    data.QID,
				Istrue: false,
			}
			testres.Res = append(testres.Res, v)
		}
	}
	return 200, testres
}
