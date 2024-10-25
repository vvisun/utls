package errutil

import (
	"errors"
	"runtime"

	"github.com/vvisun/utls/leaflog"
)

var err_map map[int32]error = make(map[int32]error)
var errMissConfig = errors.New("未知错误(没有配置错误表)")

func init() {
	for code, str := range err_table {
		err_map[code] = errors.New(str)
	}
}

func AddErrTable(tbl map[int32]string) {
	for code, str := range tbl {
		err_table[code] = str
		err_map[code] = errors.New(str)
	}
}

func Error(code int32) error {
	if _, ok := err_map[code]; !ok {
		leaflog.Debug("[error]没有配置错误表: %d\n", code)
		return errMissConfig
	}
	return err_map[code]
}

func ErrorString(code int32) string {
	if _, ok := err_table[code]; !ok {
		leaflog.Debug("[error]没有配置错误表: %d\n", code)
		return errMissConfig.Error()
	}
	return err_table[code]
}

func IsSucc(code int32) bool {
	return code == Succ
}

func IsFail(code int32) bool {
	return code != Succ
}

func GetCurrentGoroutineStack() string {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	return string(buf[:n])
}
