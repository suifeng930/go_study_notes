package main

// 比如原本打算修改全局变量，结果变成重新定义同名局部变量。

var x = 100 // 全局变量

func main() {

	println(&x, x) // 全局变量

	x := "abc" // 重新定义和初始化同名局部变量 x

	println(&x, x)
}
