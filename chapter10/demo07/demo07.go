package main

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
)

type A int
type B struct {
	A
}

func (a A) av()  {}
func (a *A) ap() {}
func (b B) bv()  {}
func (b *B) bp() {}


type gfg struct {
	Prop string
}

func (f gfg) Geek1() string {
	return f.Prop
}

func (f gfg) Geek2() {
}


func main() {

	//
	log.Println(" 输出方法集： 一样区分基类型和指针类型")

	var b B
	t := reflect.TypeOf(&b)
	s := []reflect.Type{t, t.Elem()}

	for _, t := range s {
		log.Println(t, ":")

		for i := 0; i < t.NumMethod(); i++ {
			log.Println(" ", t.Method(i))
		}
	}

	//
	log.Println(" 反射能探知当前包或外包的非导出结构成员")
	var ss http.Server
	tt := reflect.TypeOf(ss)
	method := tt.NumMethod()

	log.Println("=====> ",method)
	for i := 0; i < tt.NumField(); i++ {
		log.Println(tt.Field(i).Name)

	}


	fooType := reflect.TypeOf(gfg{})
	for i := 0; i < fooType.NumMethod(); i++ {
		method := fooType.Method(i)
		fmt.Println(method.Name)
	}


}
