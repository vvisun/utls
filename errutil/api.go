package errutil

import (
	"errors"
	"runtime"

	"github.com/vvisun/utls/leaflog"
)

var err_map map[int32]error = make(map[int32]error)

func init() {
	for code, str := range err_table {
		err_table[code] = str
		err_map[code] = errors.New(str)
	}
}

func AddErrTable(tbl map[int32]string) {
	for code, str := range tbl {
		if code < 1000 {
			leaflog.Fatal("[error]错误码必须大于1000: %d - %s\n", code, str)
		}
		err_table[code] = str
		err_map[code] = errors.New(str)
	}
}

func Error(code int32) error {
	if _, ok := err_map[code]; !ok {
		leaflog.Fatal("[error]没有配置错误表: %d\n", code)
		return errors.New("未知错误(没有配置错误表)")
	}
	return err_map[code]
}

func ErrorString(code int32) string {
	if _, ok := err_table[code]; !ok {
		leaflog.Fatal("[error]没有配置错误表: %d\n", code)
		return "未知错误(没有配置错误表)"
	}
	return err_table[code]
}

func IsSucc(code int32) bool {
	return code == Succ
}

func GetCurrentGoroutineStack() string {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	return string(buf[:n])
}
