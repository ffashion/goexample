package main

import (
	"fmt"
)

//
var a int = 3
var b float32 = 4.0
var c float64 = 5.0
var d bool = true

//定义struct类型
type mystruct_t struct {
	id int
	s  string
}

//创建map 方式2 使用make函数
var my_map1 = make(map[int]string)

//channel

func main() {

	//创建指针
	var m int = 3
	var p *int = &m
	fmt.Println(p)

	//赋值指针
	p = nil

	//创建数组
	var array0 = [2]int{1, 2}
	var array1 = [2]int{}

	fmt.Println(array0, array1)
	fmt.Println("array 截取", array0[0:1]);

	//slice 切片
	//切片是数组的抽象
	//
	var slice0 = make([]int, 3)

	//直接初始化切片
	var slice1 = []int{
		1, 2, 3,
	}
	//使用数组初始化切片 
	var array2 = []int{1,2,4,5}
	slice2 := array2[0:2]//闭0 开2

	//获取切片长度
	fmt.Println("len",len(slice2))

	//获取切片容量 
	fmt.Println("cap",cap(slice2))

	fmt.Println(slice0,slice1,slice2,array2)
	


	//赋值数组

	//方式1 赋值struct
	var _mystruct mystruct_t
	_mystruct.id = 1
	_mystruct.s = "hell0 world"

	//方式2 赋值struct
	var _mystruct1 = mystruct_t{
		1,
		"aaaa",
	}

	//访问struct
	fmt.Println(_mystruct.s, _mystruct.id)
	fmt.Println(_mystruct1)

	//创建map 下标为int 内容为string 这时候的map为nil 不能赋值
	var my_map0 map[int]string

	//创建map 并赋值
	var my_map1 = map[int]string{
		1: "helloworld",
	}

	//初始化map 设置最开始的大小为100, 如果不设置那么他的长度使用 len(my_map0)为0
	my_map0 = make(map[int]string, 100)
	my_map1 = make(map[int]string)

	//赋值map
	my_map0[0] = "this is zero"
	my_map0[1] = "this is one"
	my_map0[2] = "this is two"

	//访问map 对于map range 返回key和value
	for key, v := range my_map0 {
		fmt.Printf("my_map0[%d]=%s\n", key, v)
	}
	fmt.Println(my_map1)
}
