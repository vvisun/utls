package ecs

import (
	"log"
	"testing"
)

type comp_1 struct {
}

type comp_2 struct {
}

func (slf *comp_2) printA(a int) {
	log.Printf("comp_2.printA: %v", a)
}

func TestAll(t *testing.T) {
	e1 := NewEntity()

	var ttt *comp_1 = nil
	AddComponent(e1, ttt)
	log.Printf("1 是否comp_1: %v", ExistComponent[comp_1](e1))
	log.Printf("是否comp_2: %v", ExistComponent[comp_2](e1))

	RemoveComponent[comp_1](e1)
	AddComponent(e1, &comp_2{})
	log.Printf("2 是否comp_1: %v", ExistComponent[comp_1](e1))
	log.Printf("是否comp_2: %v", ExistComponent[comp_2](e1))

	RemoveComponent[comp_2](e1)
	AddComponent(e1, &comp_1{})
	log.Printf("3 是否comp_1: %v", ExistComponent[comp_1](e1))
	log.Printf("是否comp_2: %v", ExistComponent[comp_2](e1))

	RemoveComponent[comp_1](e1)
	log.Printf("4 是否comp_1: %v", ExistComponent[comp_1](e1))
	log.Printf("是否comp_2: %v", ExistComponent[comp_2](e1))

	RemoveComponent[comp_1](e1)

	aaa := &comp_1{}
	AddComponent(e1, aaa)
	log.Printf("5 是否comp_1: %v", ExistComponent[comp_1](e1))
	log.Printf("是否comp_2: %v", ExistComponent[comp_2](e1))

	RemoveComponent[comp_1](e1)
	log.Printf("6 是否comp_1: %v", ExistComponent[comp_1](e1))
	log.Printf("是否comp_2: %v", ExistComponent[comp_2](e1))

	bbb := AddComponent(e1, &comp_2{})
	bbb.(*comp_2).printA(2312)
	GetComponent[comp_2](e1).printA(9999)

	AddComponent(e1, &comp_2{})
	AddComponent(e1, &comp_2{})
	AddComponent(e1, &comp_2{})
}
