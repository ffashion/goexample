package main

// import path find
// $GOROOT/src/fmt  ==> $GOPATH/src/fmt
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
	fmt.Print("\n")
	//range for,  type can map, string, array, slice
	strs := []string{"aaa", "bbb", "ccc"}
	for i, s := range strs {
		fmt.Println(i, "->", s) //i = 0 ,1,2
	}
	fmt.Print("\n")
	//range for :map
	maps := map[string]string{"a": "apple", "b": "bad"}
	for k, v := range maps {
		fmt.Println(k, "->", v)
	}
	//for : while(1)
	for {
		fmt.Print("this is while(1)\n")
		break
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
