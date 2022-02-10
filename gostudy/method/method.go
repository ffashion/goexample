package main

import "fmt"

type TYPE struct {
	num  int
	num2 float32
}

func (_type *TYPE) add(para TYPE) *TYPE {
	_type.num += _type.num + para.num
	_type.num2 += _type.num2 + para.num2
	return _type
}
func main() {
	a := TYPE{1, 1.0}
	b := TYPE{2, 1.0}
	a.add(b)
	fmt.Println(a.num, " ", a.num2)
}
