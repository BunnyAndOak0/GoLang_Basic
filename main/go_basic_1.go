package main

import (
	"fmt"
)

func main() {
	//必须先声明或进行变量赋值
	var x int
	x = 10
	fmt.Println("x = ", x)

	//声明和初始化合并
	var y int = 100
	fmt.Println("y = ", y)
	//或者
	z := 200
	fmt.Println("z = ", z)

	//:=在go中属于类型推断操作，包含了声明和变量赋值两个操作
	a := getAdd(x)
	fmt.Println("a = ", a)

	//但是需要注意的是:=使用完毕后不可以再进行:=的声明操作，只可以使用=进行赋值操作
	//以下会报错
	//b := 1000
	//b := 2000
	//fmt.Println(b)

	//但是如果:=的左边有任何一个新的变量，那么语法就是正确的，两者需要类型一致，
	//因为第一次推断出来的类型不能再改变，后续操作只允许进行简单的赋值操作
	name, age := "alice", 18
	fmt.Println("name = ", name, "age = ", age)

	name, sex := "100", "女"
	fmt.Println("name = ", name, "sex = ", sex)

	//函数定义
	//如果不需要所有的返回值，可以将不需要的返回值使用_来表示丢弃这段结果
	d, _ := getResult1()
	fmt.Println("d = ", d)

	//如果函数参数类型相同则不同的参数可以共享数据类型 ===> 划重点！！！参数！！！
	b, _ := getResult2(a, a)
	fmt.Println("如果函数的参数类型相同，则不同的参数（注意是参数！！！）可以共享数据类型，===》b = ", b)
}

func getAdd(x int) int {
	return x  * 2
}

func getResult1() (string, int) {
	return "hello", 100
}

func getResult2(a, b int) (int, int) {
	return 100, 200
}
