package tools

import (
	"math/rand"
	"time"
)

// 产生随机数
func Randnum(scope int) int {
	rand.New(rand.NewSource(time.Now().Unix()))
	i := rand.Intn(scope) + 1
	return i
}

// 产生一个在n范围内m个数的队列
func RandNumSlice(n int, m int) []int {
	rand.New(rand.NewSource(time.Now().Unix()))
	// 生成一个包含 0 到 9 的随机排列
	randomPerm := rand.Perm(n)
	return randomPerm[:m]
}
