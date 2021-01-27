package main

import (
	"log"
	"reflect"
)

type user struct {
	name string
	age  int
}

type manager struct {
	user
	title string
}

func main() {

	//   对于匿名字段，可用多级索引（按定义顺序）直接访问
	var n manager

	t := reflect.TypeOf(n)

	name, _ := t.FieldByName("name") //  按名称查找
	log.Println(name.Name, name.Type)

	age := t.FieldByIndex([]int{0, 1}) // 按多级索引查找
	log.Println(age.Name, age.Type)
}
