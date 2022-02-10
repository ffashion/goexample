package main

import (
	"fmt"
)

//var 强调创建了一个变量
var one_var int = 11
var (
	count int = 3
	this  int = 11
)

//常量 可以省略标识符 string
const one_const int = 11
const (
	s string = "qqq"
	m int    = 100
	n        = 77 //省略类型
)

//特殊用法
//iota 是编译器实现的一个常量 可以在编译器自增 第一个iota为0
//这样可以实现枚举
const (
	ZHANG_ONE = iota
	ZHANG_TWO = iota
)

func main() {
	fmt.Println(s)
	fmt.Println(ZHANG_ONE)
}
