package main

import "fmt"

func op(a string, b string) string {
	return a + b
}

func append(a string, b string, op func(string, string) string) string {
	return op(a, b)
}

func main() {
	//normal pointer
	a := 3
	fmt.Println(&a)

	//function pointer
	fmt.Println(append("aaa", "bbb", op))

}
