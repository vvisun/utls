package randutil

import (
	"math/rand"
	"time"
)

func RandSeed() {
	rand.Seed(time.Now().UnixNano())
}

func RandInt64(minV int64, maxV int64) int64 {
	if minV == maxV {
		return minV
	}
	return minV + int64(rand.Intn(int(maxV-minV+1)))
}

func RandInt32(minV int32, maxV int32) int32 {
	if minV == maxV {
		return minV
	}
	return minV + int32(rand.Intn(int(maxV-minV+1)))
}

func RandInt(minV int, maxV int) int {
	if minV == maxV {
		return minV
	}
	return minV + rand.Intn(maxV-minV+1)
}

func Shuffle(deck []int32) []int32 {
	for i := len(deck) - 1; i > 0; i-- {
		// 生成一个[0, i]范围内的随机索引
		j := rand.Intn(i + 1)
		// 交换deck[i]和deck[j]
		deck[i], deck[j] = deck[j], deck[i]
	}
	return deck
}
