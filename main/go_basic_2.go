package main

import "fmt"

func main() {
	fmt.Println("接口struct")

	//创建一个数据结构 也就相当于java中的对象 其中最后一个”,“不能省略
	hm_horse := Animal{
		name:  "heima",
		speak: "ni hao",
	}
	fmt.Println("hm_horse = ", hm_horse) //   hm_horse =  {heima ni hao}

	bm_horse := Animal{
		name: "baima",
	}
	fmt.Println(bm_horse) //   {baima }

	//可以先定义一个空的数据结构进行赋值 还可以 先赋值一部分之后 再进行赋值
	fly_bird := Animal{}
	fly_bird.name = "pink bird"
	//也可以省略结构中的key的部分 直接进行赋值 但是位置顺序必须和key的顺序对应
	walk_monkey := Animal{"mokey", "banana", 100}
	fmt.Println(walk_monkey)

	walk_monkey.add()
	fmt.Println("walk_monkey.age = ", walk_monkey.age)

	//构造器 ---> 可以根据根据类构造出类的数据结构 但是一般情况下 还是会使用内置的new()方法
	fish := newAnimal1()
	fish.add()
	fmt.Println("fish.age = ", fish.age)

	ant := newAnimal2()
	substract(&ant)
	fmt.Println("ant.age = ", ant.age)

	//go中内置有new()用于为一个数据结构分配内存 new()等价于x&{}
	//1.使用new()
	tiger := new(Animal)
	tiger.name = "老虎"
	tiger.speak = "嗷呜"
	tiger.age = 3
	fmt.Println("tiger = ", tiger)

	//2.使用x&{}
	lion := &Animal{
		"狮子",
		"lion lion",
		4,
	}
	fmt.Println("lion = ", lion)

	//composition ---> 表示在一个数据结构中嵌套另一个数据结构
	horse_animal := &horse{
		//字段的名称必须和数据结构完全相同
		Animal: &Animal{
			name:  "马",
			speak: "吁",
			age:   5,
		},
		color: "blue",
	}
	fmt.Println("马的名字为 = ", horse_animal.Animal.name)
	fmt.Println("调用了add的方法后，age = ", horse_animal.add)

	//重载

}

type Animal struct {
	name  string
	speak string
	age   int
}

//属于数据结构的函数 可以为数据结构定义的自己的函数
func (animal *Animal) add() {
	animal.age += 10
}

//传地址的方法
func substract(animal *Animal) {
	animal.age -= 1
}

//构造器
//引用类型构造器
func newAnimal1() *Animal {
	return &Animal{
		name:  "fish",
		speak: "bulk",
		age:   1,
	}
}

//非引用类型构造器
func newAnimal2() Animal {
	return Animal{
		name:  "ant",
		speak: "tiktok",
		age:   2,
	}
}

type horse struct {
	*Animal
	color string
}
