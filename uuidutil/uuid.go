package uuidutil

import (
	"sync/atomic"

	uuid "github.com/satori/go.uuid"
)

var gGid int64 = 0

// 内存级ID生成
func GenGid() int64 {
	atomic.AddInt64(&gGid, 1)
	return atomic.LoadInt64(&gGid)
}

// 分布式ID生成
func GenUUID() string {
	u2 := uuid.NewV4().String()
	return u2
}
