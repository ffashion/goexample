package main

import (
	"fmt"
)

//体验下 go的 sb定义
func test0(a int, m, n string) string {
	fmt.Println(a)
	return m + n
}

//可变参数 需要是参数的最后一个参数
func test1(ret ...int) (int, int) {
	fmt.Print("test1 \n")
	return ret[0], ret[1]
}

func main() {
	test1(2, 3)
	fmt.Println(test0(1, "a", "b"))
}
