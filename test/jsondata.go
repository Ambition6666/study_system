package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.New(rand.NewSource(time.Now().Unix()))
	// 生成一个包含 0 到 9 的随机排列
	randomPerm := rand.Perm(100)

	fmt.Println("Random permutation:", randomPerm[:10])
}
