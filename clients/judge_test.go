package clients

import (
	"context"
	"fmt"
	judge "studysystem/api/proto/judge"
	"testing"
)

func TestJudge(t *testing.T) {
	conn, err := InitJudgeGRPC()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	ctx := context.Background()

	// 提交代码
	res, err := JudgeCli.Judge(ctx, &judge.JudgeRequest{
		ProblemID: 1,
		Code:      []byte("a, b = map(int, input().split())\nprint(a + b)"),
		LangID:    3,
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("提交成功: %v\n", res.JudgeID)

	// 获取结果
	result, err := JudgeCli.GetResult(ctx, &judge.GetResultRequest{
		JudgeID: res.JudgeID,
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("result.Result: %+v\n", result.Result)
}
