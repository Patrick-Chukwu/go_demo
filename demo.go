package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x string = "Hello, pat"
	var num int8 = 100
	var y string
	var w bool
	var u uint16
	var a int8
	var b rune
	var c byte
	y = "Bring this back"

	fmt.Println(x)
	fmt.Println(num)
	fmt.Println(u,y,w,a, b,c)

	// map
	mp := map[string]bool{}
	mp["b"] = false
	value, ok := mp["b"]
	fmt.Println(value, ok)

	num1 := "hi"
	num2 := 2
	num3 := num1 + fmt.Sprint(num2)
	fmt.Println((num3))


	test := "1E"
	result,err := strconv.ParseInt(test,2,0)
	fmt.Println(result, err)


}