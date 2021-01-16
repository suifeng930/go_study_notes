## go语言学习笔记01章

------

### 1.1 特征

------

Go从零开始，没有历史包袱，在汲取众多经验教训后，可从头规划一个规则严谨、条理简单的世界。

Go语言简单的本质： 语言规则严谨没有歧义，更没有什么黑魔法变异用法，任何人写出的代码都基本一致，放弃部分“灵活”和“自由”，换来更好的维护性。保留指针，但默认阻止指针运算。将切片和字典作为内置类型，从运行时从层面进行优化。

#### 并发模型

------

Goroutine是Go最显著的特征。 它用类协程的方式来处理并发单元，却又在运行时层面做了更深度的优化处理。者使得语法上的并发编程变得极为容易。无须处理回调，无须关注执行绪切换，仅一个关键字，简单而自然。

搭配channel，实现CSP模型，将并发单元间的数据耦合拆解下来，各司其职，这对所有纠结于内存共享，锁粒度的开发人员都是一个可期盼的解脱。

#### 静态链接

------

将运行时、依赖库直接打包到可执行文件内部，简化了部署和发布操作，无须事先安装运行环境和下载诸多第三方库。这种简单方式，对于编写系统软件有着极大好处，因为裤依赖一直都是个麻烦。

### 1.2 简介

------

#### 源文件

------

源码文件使用UTF-8编码，对Unicode支持良好。每个源文件都属于包的一部分，在文件头部用`package`	声明所属包名称。

以`xxx.go`作为文件扩展名，语句结束分号会被默认省略，支持C样式注释。入口函数`main`没有参数，且必须放在main包中。

```go
package main
import ( "fmt")
func main(){
		fmt.Println("Hello,world!!!")
}
// 请删除未使用的导入，否则编译器会将其当作错误。
```

可以直接运行，也可以编译为可执行文件。

```sh
$ go run main.go
hello,world!!!
$ go build mian.go  # 编译文件
```

#### 变量

------

使用`var`定义变量，支持**类型推断**。基础数据类型划分清晰明确，有助于编写跨平台应用。编译器确保变量总是被初始化为零值，避免出现意外状况。

------

```go
package main
func main(){
		var x int32
		var s="hello ,world"
		println(x,s)
}
```

在函数内部，还可以省略`var`关键字，使用更简单的定义模式。

```go
package main

func main(){
		x :=100 //注意： 赋值符号不同
		println(x)
}

```

------

编译器将未使用的局部变量定义当作错误。

------



#### 表达式

------

Go仅有三种流控制语句，与大多数语言相比，都可称得上简单。

------

**if**

```go
func main(){
		x :=100
		if x>0{
				println("x")
		}else if x<0 {
				println("-x")
		}else{
				println("0")
		}
}
```

**switch**

```go
func main(){
		x :=100
		
		switch	{
				case  x>0 :
						println("x")
				case x<0 :
						println("-x")
				default:
						println("0")
		}
}
```

**for**

```go
func main(){
		for i :=0; i<5; i++ {
					println(i)
		}
		
		for i :=4; i>=0; i--{
				  println(i)
		}
}
```

```go
func main(){
			x :=0
			
			for x<5{ //        相当于 while（x<5） {...}
					println(x)
					x++
			}
}
```

```go
func main(){
		x :=4 
		
		for {  // 相当于 while (true) {...}  类似死循环
				println(x)
        x--
        
        if x<0{
        		break     // 一定要注意 跳出循环的逻辑，不然就是死循环
        }
		}
}
```

在迭代遍历时，`for...range`除元素外，还可以返回索引。

```go
func main(){
		x :=[] int(100,101,102)
		
		for i,n :=range x{
			println(i,":",n)
		}
}
//输出
// 0:100
// 1:101
// 2:102
```

#### 函数

------

函数可以定义多个返回值，甚至对其命名

------

```go
func main(){
			a,b :=10,2        //定义多个变量
			c,err :=div(a,b)  // 接受多返回值
			
			fmt.println(c,err)
}
```

函数是第一类型，可作为参数或返回值

```go
func test(x int) func(){   // 返回函数类型
		return func()	{        // 匿名函数
						println(x)     // 闭包
		}
}

func main(){
		x :=100
		
		f :=test(x)
		
		f()
}
```

用`defer`定义延时调用，无论函数是否出错，它都确保结束前被调用

```go
package main

func test (a ,b int) {
		defer println("dispose...")    //常用来释放资源 、解除锁定、或执行一些清除操作
																	// 可定义多个defer、 按FILO 顺序执行
		println(a/b)
		
}

func main() {
		test(10,0)
}
```

输出：

```sh
$ go run test.gp

dispose

panic: runtime error: integer divide by zero

```

#### 数据

------

切片`slice`可实现类似动态数组的功能。

```go
package main

import(
			"fmt"
)

func main() {
		x :=make ([]int,0,5)      //创建容量为5的切片
		for i:=0; i<8; i++ {
				x=append(x,i)         //追加数据，当超出容量限制时，自动分配更大的存储空间。
		}
		
		fmt.println(x)
}
//输出：   [0 1 2 3 4 5 6 7]
```

将字典`map`类型内置，可直接从运行时层面获得性能优化。

```go
package main

import ( "fmt")

func main() {
			m :=make(map[string]int)               // 创建字典类型对象
			
			m["a"]=1                               // 添加或设置
			
			x,ok :=m["b"]                         // 使用ok-idiom 获取值，可知道 key/value 是否存在
			
			fmt.println(x,ok)
			
			delete(m,"a")                         //删除
}
// 所谓 ok-idiom 模式： 值在多返回值中用一个名为ok的布尔值来标示操作是否成功，因为多操作默认返回零值，所以须额外说明。
```

结构体`struct`可以匿名嵌入其他类型

```go
package main

import( "fmt")

type user struct{    //	结构体类型
		name string
		age byte
}

type manager struct {
		user            // 匿名嵌入其他类型
		title string
}

func main() {
			var m manager   
			
			m.name ="Tom"    //直接访问匿名字段的成员
			m.age = 29
			m.title="CTO"
			
			fmt.println(m)
}
//输出： {{Tom 29} CTO}
```

#### 方法

------

可以为当前包内的任意类型定义方法

```go
package main

type X int

func (x *X) inc() {            //名称前的参数称作，receiver, 作用类似 python self
  *x++ 
}

func main() {
  	var x X
 		x.inc()
  println(x)
}
```

还可以直接调用匿名字段的方法，这种方式可以实现与继承类似的功能。

```go
package main

import("fmt")

type user struct {
  	name string
  	age 	byte
}

func (u user) ToString() string {
  return fmt.Sprintf("%+v",u)
}

type manager struct{
  	user 
  	title string
}

func main(){
  	var m manager
  	m.name="Tom"
  	m.age=29
  	
  println(m.ToString())  //调用user.ToString() 
}
```

#### 接口

------

接口采用了 `duck type`方式，也就是说无须在实现类型上添加显示声明。

```go
package main

import("fmt")

type user struct{
  	name string
  	age byte
}

func (u user) Print()	{
  fmt.Printf("%+v\n",u)
}

type Printer interface {         // 接口类型
  Print()
}
	
func main() {
  	var u user
  	u.name="Tom"
  	u.age=29
  	
 		var p Printer=u               //只要包含接口所需的全部方法，即表示实现了该接口
  	
    p.Print()  
}

// 另外有空接口类型 interface{}  ,用途类似OOP里面的system.Object, 可以接受任意类型的对象
```

#### 并发

------

整个运行时完全并发化设计。凡你能看到的，几乎都在以`goroutine`方式运行.这是一种比普通协程或线程更加高效的并发设计，能轻松创建和运行成千上万多并发任务。

------

```go
package main

import(
	"fmt"
	"time"
)

func task(id int) {
  
  for i :=0; i<5; i++ {
    fmt.Printf("%d:%d\n",id,i)
    time.Sleep(time.Second)
  }
}

func main() {
  go task(1)       //创建goroutine
  go task(2)
  
  time.Sleep(time.Second*6)
}
```

通道`channel`与`goroutine` 搭配，实现用通讯代替内存共享的CSP模型

------

```go
package main

//消费者

func consumer(data chan int,done chan bool) {
  for x :=range data{                   //接受数据，直到通道被关闭
    println("recv:",x)
  }
  done <-true       //通知main(),消费结束
}

// 生产者
func producer( data chan int) {
  for i :=0; i<4;i++ {
    	data <-i           // 发送数据     
  }
  close(data)            //生成结束，关闭通道
}

func main() {
  done :=make(chan bool)    //用于接受消费结束信号
  data :=make(chan int)     //数据管道
  
  go consumer(data,done)    //启动消费者
  go producer(data)         //启动生产者
  
  <-done                    //阻塞，直到消费值发回结束信号。
  
}
```


