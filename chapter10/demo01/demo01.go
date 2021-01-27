package main

import (
	"log"
	"reflect"
)

type X int

func main() {

	var a X = 10
	t := reflect.TypeOf(a)
	log.Println(t.Name(), t.Kind())
	//    TYPE  表示 真实类型 （静态类型）
	//    Kind  表示 基础类型 （底层类型）

}
