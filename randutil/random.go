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
