package ecs

import (
	"log"
	"reflect"
	"sync/atomic"
)

var gEntityId int64 = 0

func genEntityId() int64 {
	atomic.AddInt64(&gEntityId, 1)
	return atomic.LoadInt64(&gEntityId)
}

type Entity struct {
	eid        int64                        // Entity ID
	components map[reflect.Type]interface{} // 组件
}

func NewEntity() *Entity {
	return &Entity{eid: genEntityId(), components: make(map[reflect.Type]interface{})}
}

func (e *Entity) EID() int64 {
	return e.eid
}

//----------------------------------------------------------

// 检查指针是否是空值
func isNil(comp interface{}) bool {
	v := reflect.ValueOf(comp)
	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

// 给Entity添加组件实例comp
func AddComponent(e *Entity, comp interface{}) interface{} {
	compType := reflect.TypeOf(comp)
	if compType == nil || compType.Kind() != reflect.Ptr || isNil(comp) {
		log.Fatal("请传指针")
		return nil
	}
	if _, ok := e.components[compType]; ok {
		log.Fatalf("不可重复添加相同类型的组件 %s", compType)
		return nil
	}
	if nil == comp {
		log.Fatal("comp为nil, 如果可以赋空, 请使用RemoveComponent")
	}
	e.components[compType] = comp
	return comp
}

// 给Entity更新组件
func ReplaceComponent(e *Entity, comp interface{}) interface{} {
	compType := reflect.TypeOf(comp)
	if compType == nil || compType.Kind() != reflect.Ptr {
		log.Fatal("请传指针")
		return nil
	}
	if nil == comp || isNil(comp) {
		log.Fatal("comp为nil, 如果可以赋空, 请使用RemoveComponent")
	}
	e.components[compType] = comp
	return comp
}

// 移除Entity上类型为T的组件
func RemoveComponent[T any](e *Entity) {
	compType := reflect.TypeOf((*T)(nil)) //.Elem()
	if compType == nil {
		log.Fatal("请传类")
		return
	}
	delete(e.components, compType)
}

// 判断Entity上是否存在类型为T的组件
func ExistComponent[T any](e *Entity) bool {
	compType := reflect.TypeOf((*T)(nil)) //.Elem()
	if compType == nil {
		log.Fatal("请传类")
		return false
	}
	if _, ok := e.components[compType]; ok {
		return true
	}
	return false
}

// 获取Entity上类型为T的组件
func GetComponent[T any](e *Entity) *T {
	compType := reflect.TypeOf((*T)(nil)) //.Elem()
	if compType == nil {
		log.Fatal("请传类")
		return nil
	}
	if obj, ok := e.components[compType]; ok {
		return obj.(*T)
	}
	return nil
}
