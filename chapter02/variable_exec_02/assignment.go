package main

// 1.简短模式并不总是重新定义变量，也可能是部分退化的赋值操作。
// 2.退化赋值的前提条件：最少有一个新变量被定义，且必须是同一作用域

func main() {

	x := 100

	println(&x)

	x, y := 200, "abc" //注意： x 退化为赋值操作， 仅有y是变量定义

	println(&x, x)
	println(&y, y)



	// 2.退化赋值的前提条件：最少有一个新变量被定义，且必须是同一作用域
	z :=100
	println(&z)

	//z :=200           //: no new variables on left side of :=

	println(&z,z)


	a :=100
	println(&a,a)
	{
		a,b :=200,"sa"  // 不同作用域，全部是新的局部变量定义
		println(&a,a,b)
	}

}
