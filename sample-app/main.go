package main

import "fmt"

// var a int
// var b = 1
// var c int = 1

var a, b, c int = 1, 2, 3

var (
	back_quote string = `아리랑\n
	아리랑\n
	  아라리요`
	double_quote string = "test\ntest\n"
)

const (
	d float32 = 1.0
	e float32 = 2.2
)

func var_test() {
	a_inside := 100
	fmt.Println(a_inside)
	fmt.Println(a, b, c)
	fmt.Println(d, e)
	fmt.Println(back_quote, double_quote)
}

func conversion() {
	var i int = 100
	var u uint = uint(i)
	var f float32 = float32(i)
	println(f, u)

	str := "ABC"
	bytes := []byte(str)
	str2 := string(bytes)
	println(bytes)
	println(str2)
}

func main() {
	var_test()
	conversion()
}
