package main

import (
	"log"
	"reflect"
)

type X int
type Y int
func main() {

	var a ,b X =100,122
	var c Y =300
	ta := reflect.TypeOf(a)
	tb := reflect.TypeOf(b)
	tc := reflect.TypeOf(c)

	// 在类型判断上要注意： type --->   真实静态类型      kind  --->  基础数据类型（底层数据类型）
	log.Println("ta==tb  :",ta==tb)
	log.Println("ta==tc  :",ta==tc)
	log.Println("ta.Kind()==tc.Kind()  :",ta.Kind()==tc.Kind())
}
