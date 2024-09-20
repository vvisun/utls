package errutil

import (
	"errors"
	"runtime"
)

var err_map map[int32]error = make(map[int32]error)

func init() {
	AddErrTable(err_table)
}

func AddErrTable(tbl map[int32]string) {
	for code, str := range tbl {
		err_table[code] = str
		err_map[code] = errors.New(str)
	}
}

func Error(code int32) error {
	return err_map[code]
}

func ErrorString(code int32) string {
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
