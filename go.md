# 优势

+ 部署简单
  + 可直接编译成机器码
  + 不依赖其他库
  + 直接运行即可部署
+ 静态类型语言
  + 编译的时候能检查出大多数问题
+ 语言层面的并发
  + 天生基因支持并发
  + 充分利用多核
+ 强大的标准库
  + runtime系统调度机制
  + 高效的GC垃圾回收
  + 丰富的标准库
+ 简单易学
  + 只有25个关键字
  + 支持内嵌C语法
  + 面向对象特征（封装、继承、多态）
  + 跨平台



# 开始前

要先定义package



# fmt包

+ `fmt.`

  + Print、Println

    + ("hello, world")
    + ("a", "b", "c")：Print多个字符串间没空格，Println有空格

  + Printf

    + ("hello, world")

    + ```go
      var a = "aaa"
      var b = "bbb"
      fmt.Printf("a = %v b = %v", a)
      
      /*
      %v：输出值
      %d：输出十进制数
      %b：输出二进制数
      %T：输出数据类型
      */
      ```



# 注释

```go
// 单行注释

/*
	多行注释
*/
```





# 数据类型

## 基本数据类型

+ 整型

  + ```
    uint8：0-255（2^8-1) == byte
    uint16：0-65535
    uint32
    uint64
    int8：-128-127
    int16
    int32 == rune
    int64
    int 会在32位或64平台下分别取对应位数的长度
    ```

+ 浮点型

  + ```
    float32
    float64
    complex64：32位实数和虚数
    complex128
    ```

+ 布尔型

  + `bool`

+ 字符串

  + `string`

## 复合数据类型

+ 数组
+ 切片
+ 结构体
+ 函数
+ map
+ 通道channel
+ 接口

```go
func main() {
    a := 1
    fmt.Printf("a的类型为：%T\n". a)	// a的类型为int
    unsafe.Sizeof(a)	// 返回a的内存占用字节数
}
```

## 强转类型

可能出现精度丢失

```go
var a1 int8 = 8
var a2 int16 = 10
fmt.Printf(int(16)a1 + a2)
```

## 数字字面量语法

%v：原样输出

%d：十进制输出

%b：二进制输出

%o：八进制输出

%x：十六进制输出

```go
v := 0b1001
v2 := 0o377
v3 := 0xac
```





# 运算符

```
逻辑运算符
&&：与
||：或
!：非

位运算符
&：按位与
|：按位或
^：按位异或
<<：左移
>>：右移

&变量：返回变量存储地址
*变量：这个变量是指针变量
var v1 = 4
var ptr *int
ptr = &v1
fmt.Println(*ptr)	// 4
fmt.Println(ptr)	// 0xc00008c058
v1 = 10
fmt.Println(*ptr)	// 10


```



# 函数

```go
func function_name([parameter list]) [return_types] {
    body
}

func sum(a int, b int) int {
    return a + b
}

// 如果参数都一样，可以只在最后一个变量写类型
func sum(a, b int) int {
    return a + b
}

func sum(a, b int) {	// 没有返回值可以省略返回值类型
    fmt.Println(a+b)
}

func swap(a, b int) (int, int) {	// 可以返回多个值
    return b, a
}

func sum(a int, b int,
         transform func(int) int) int {	// 函数可以作为一个函数的参数
    return transform(a) + transform(b)
}
func main() {
    square := func(x int) int {
        return x * x
    }
    fmt.Println(sum(1, 2, square))
}
```



# 变量

+ 定义位置

  + 函数内定义的变量是局部变量

  + 函数外定义的变量是全局变量

+ 全局变量和局部变量名可以相同，函数内的局部变量会被优先考虑

```go
var d = 2	// 函数外定义变量则必须要用var
func main() {
    a := 1	// 在函数内定义变量可以直接用 :=，并在定义同时赋值
    var b = 2
    var c int // 如果定义时不赋值，则需要指定类型
    fmp.Println(a)
    
    var (
    	username string
        age int
        sex string
    )
    
    var name, age = getUserinfo()
    // _为匿名变量，不会分配内存空间，不存在重复声明
    var name2, _ = getUserinfo()
}

func getUserinfo() (string, int) {
    return "zhangsan", 20
}
```

## 命名规则

+ 由数字、字母、下划线组成，首个字符不能为数字
+ 驼峰命名：age、maxAge（表示私有）、MaxAge（表示公有）、特有名词全大写（如DNS）



# 常量

```go
// 常量定义时必须赋值
const pi = 3.1415926

// 一次声明多个常量
const (
	A = "AAA"
    B = "BBB"
    C	// 没赋值则与上一行相同
)
```

## iota

iota是go的常量计数器，只能在常量的表达式中使用

iota在const关键字出现时被重置为0，const中每新增一行常量声明，将使iota计数一次（iota可理解为const语句块中的行索引）

```go
const a = iota // a=0
const (
	b = iota	// b=0
    c	// c=1
    d	// d=2
    _
    e	// e=4
    f = 100
    g = iota // g=6
)

const (
	n1, n2 = iota+1, iota+2	// 1 2
    n3, n4	// 2 3
    n5, n6	// 3 4
)
```





# 分支结构

## if

```go
func main() {
    a := 10
    if a > 10 {	// 条件不用加括号
        fmt.Println("a大于10")
    } else if a > 5 {
        fmt.Println("a大于5")
    } else {
        fmt.Println("a不大于5")
    }
    
    if b:= 3; a > 10 {	// if里可以初始化变量，只能在if-else if-else语句块内使用，用分号隔开条件
        fmt.Println("a大于10")
    }
}
```

## switch...case

```go
a := 10
switch a {
	case 1:	// case 只会执行一个，所以可以省略break
    fmt.Println("a==1")
    case 5:
    fmt.Println("a==5")
    case 10:
    fmt.Println("a==10")
    default:
    fmt.Println(a)
}
```

## 循环（只有for）

``` go
for i := 1; i < 5; i++ {
	fmt.Println(i)
}

// 模拟while
i := 1
for i < 5 {
    fmt.Println(i)
    i++
}
```









# 数组

```go
var arrayName [arraySize]dataType

a := [5]int{1, 2, 3, 4, 5}
```



# 切片/动态数组

```go
var sliceName []type
var sliceName []type = make([]type, length)
var sliceName []type = make([]type, length, capacity)

len(sliceName)	// 长度
cap(sliceName)	// 容量

a := make([]int, 0)
a = append(a, 1, ,2 ,3 ,4)

a := []int{1, 2, 3, 4}	// 用字面值创建

a[0] = 5	// 修改值
```



# Map

```go
var m = map[key类型]val类型
m := make(map[string]int)

m := map[string]int {
    "a" : 1,
    "b" : 2
}
fmt.Println(m)	// map[a:1 b:2]


```



# 结构体、指针

```go
type struct_variable_type struct {
   member definition
   member definition
   ...
   member definition
}

type Point struct {
    x int
    y int
}
func main() {
    p := Point{1, 2}
    fmt.Println(p)	// {1, 2}
    p.x = 3	// 修改值
    
    q := p
    p.x = 3	// 默认按值复制
    fmt.Println(p)	// {3, 2}
    fmt.Println(q)	// {1, 2}
    
    q := &p
    p.x = 3
    fmt.Println(p)	// {3, 2}
    fmt.Println(*q)	// {3, 2}
    
    q.y = 10	// 可以省略*
    fmt.Println(*q)	// {3, 10}
}
```

## 结构的方法

```go
type Point struct {
    x int
    y int
}

func (p Point) SetPoint(x, y int) {	// 方法名推荐首字母大写
    p.x = x
    p.y = y
}

func main() {
    p := Point{1, 2}
    p.SetPoint(3, 4)
    fmt.Println(p) // {1 2}。没用指针，值传递，方法在内部又复制了一份结构，不影响原结构
}

// 大部分情况方法要加*
func (p *Point) SetPoint(x, y int) {	// 方法名推荐首字母大写
    p.x = x
    p.y = y
}
func main() {
    p := Point{1, 2}
    p.SetPoint(3, 4)
    fmt.Println(p) // {3 4}
}
```



# 接口

```go
type interface_name interface {
    Print() // 规定必须要实现的方法，无需func关键字
}

type Shape interface {
    Print() // 规定必须要实现的方法，无需func关键字
}

// 结构要实现接口，定义同名的方法即可
type Rectangle struct{}
type Triangle struct{}

func (r Rectangle) Print() {
    fmt.Println("矩形")
}
func (r Triangle) Print() {
    fmt.Println("三角形")
}

func main() {
    var s Shape
    s = Rectangle{}
    s.Print()
    s = Triangle{}
    s.Print()
    
    printShape(Rectangle{})
}

func printShape(s Shape) {
    s.Print()
}
```



# 错误处理

```go
n, err := fmt.Println("dd")
if err != nil {
    // 执行正常的代码
} else {
    // 执行异常代码
}
```



# 并发

## Goroutines

```go
func func1() {
    time.Sleep(500 * time.Millisecond)
    fmt.Println("函数1")
}
func func2() {
    fmt.Println("函数2")
}

func main() {
    go func2()	// 开启一个线程执行func2
    func1()
}
```

## Channels——线程间通信

```go
func main() {
    ch := make(chan string) // chan关键字 管道里的数据类型
    go func2(ch)
    res1 := <- ch // 接收管道数据（发送、接收管道数据会阻塞线程）
    
    go func1(ch)
    res2 := <- ch
}

func func1(ch chan string) {
    ch <- "函数1"	// 给管道发送数据
}
func func2(ch chan string) {
    ch <- "函数2"
}
```



# 命令

```bash
go mod init 项目名 # 初始化项目
go mod tidy # 导入需要依赖，清除无用依赖

go get gorm.io/gorm # 下载第三方库到本地，并导入go.mod，更新go.sum校验文件
```

