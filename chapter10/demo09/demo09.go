package main

import (
	"fmt"
	"log"
	"reflect"
)

type X int

func (X) String() string {
	return ""

}

func main() {

	log.Println("辅助判断方法 Implements()  ConvertibleTo()  AssignableTo() 都是运行期进行动态调用和赋值所必须的  ")
	var a X
	t :=reflect.TypeOf(a)

	st :=reflect.TypeOf((*fmt.Stringer)(nil)).Elem()
	fmt.Println(t.Implements(st))

	it :=reflect.TypeOf(0)
	fmt.Println(t.ConvertibleTo(it))

	fmt.Println(t.AssignableTo(st),t.AssignableTo(it))
}
