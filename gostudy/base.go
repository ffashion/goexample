package main

import (
	"fmt"
)

//可以省略的东西一般放在后面 比如类型
func main() {
	//使用const
	const count int = 5

	//for
	//定义并赋值
	for i := 0; i <= count; i++ {

		fmt.Printf("%d count\n", i)
		if i <= 2 {
			continue
		}
		if i >= 3 {
			break
		}

	}
	// if
	if true {
		fmt.Printf("true\n")
	}

	// switch case
	var case_one = 1
	switch case_one {
	case 1:
		fmt.Printf("1\n")
		fallthrough
	case 2:
		fmt.Printf("2\n")
	case 3:
		fmt.Printf("3\n")
	}
}
