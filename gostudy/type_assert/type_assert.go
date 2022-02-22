package main

import (
	"fmt"
)

func main() {

	//只有接口才有类型断言
	var b interface{} = 2

	//断言b为int 断言成功把值给value
	value, ok := b.(int)
	fmt.Print(value, ok, "\n")

	v, k := b.(string)
	if k {
		fmt.Println(v, k)
	}

	//type 只能用于switch
	switch b.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	}
}
