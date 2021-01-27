package main

import (
	"log"
	"reflect"
)

func main() {

	x := 100

	tx := reflect.TypeOf(x)
	tp := reflect.TypeOf(&x)

	log.Println("tx  -->:", tx)
	log.Println("tp  -->:", tp)
	log.Println("tx equal tp  -->:", tx == tp)
	log.Println("tx.Kind() -->:", tx.Kind())
	log.Println("tp.Kind() -->:", tp.Kind())
	log.Println("tx equal tp.Elem() -->:", tx == tp.Elem())
	// tp.Elem()   返回指针  数组、 切片 字典 或 channel  的基类型

	log.Println("tp.Elem()   返回指针  数组、 切片 字典 或 channel  的基类型")
	xc :=reflect.TypeOf(map[string]int{}).Elem()
	xm :=reflect.TypeOf([]int32{}).Elem()
	log.Println(xc)
	log.Println(xm)

}
