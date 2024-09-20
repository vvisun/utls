package datastruct

import "reflect"

// 获取结构体名字
func GetStructName(ptr interface{}) string {
	if t := reflect.TypeOf(ptr); t.Kind() == reflect.Ptr {
		return t.Elem().Name()
	} else {
		return t.Name()
	}
}
