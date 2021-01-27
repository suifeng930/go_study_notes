package main

import (
	"log"
	"reflect"
)

func main() {

	//  通过实际对象获取类型外， 也可以直接构造一些基础复合类型
	a := reflect.ArrayOf(10, reflect.TypeOf(byte(0)))
	m := reflect.MapOf(reflect.TypeOf(""), reflect.TypeOf(0))
	log.Println("reflect.ArrayOf(10,reflect.TypeOf(byte(0)))  -->:", a)
	log.Println("reflect.MapOf(reflect.TypeOf(\"\"),reflect.TypeOf(0))  -->:", m)

}
