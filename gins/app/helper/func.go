package helper

import (
	"math/rand"
	"time"
)

func Random(max int) int {
	//将时间戳设置成种子数
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max)
}
