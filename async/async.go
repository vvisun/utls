package async

import (
	"github.com/vvisun/utls/leaflog"
)

func pcall(fn func()) {
	defer func() {
		if err := recover(); err != nil {
			leaflog.Error("aync/pcall: Error=%v", err)
		}
	}()

	fn()
}

func Run(fn func()) {
	go pcall(fn)
}
