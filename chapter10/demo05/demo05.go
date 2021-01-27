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

	//只有获取结构体指针的基类型后，才能遍历它的字段

	var m manager
	t := reflect.TypeOf(m)

	if t.Kind() == reflect.Ptr {
		t = t.Elem() // 获取指针的基类型
	}
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		log.Printf("f.Name: %s,f.Type:%s,f.Offset:%d\n", f.Name, f.Type, f.Offset)

		if f.Anonymous { // 输出匿名字段结构
			for x :=0;x<f.Type.NumField();x++ {
				af :=f.Type.Field(x)
				log.Printf("af.Name: %s ;af.Type:%s\n",af.Name,af.Type)

			}

		}
	}

}
