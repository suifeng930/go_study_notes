package main

import (
	"log"
	"reflect"
)

type user struct {
	name string `field:"name" type:"varchar(50)"`
	age  int    `field:"age" type:"int"`
}

func main() {

	var u user

	log.Println("可用 反射提供 struct tag  ，还能自动分解， 其常用")
	t := reflect.TypeOf(u)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		log.Printf("%s:%s %s \n", f.Name, f.Tag.Get("field"), f.Tag.Get("type"))

	}

}
