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


	test := "101"
	result,err := strconv.ParseInt(test,2,0)
	fmt.Println(result, err)
// comparison operator
	test1 := uint32(8)
	test2 := test1 == uint32(8)
	fmt.Println(test2)

// if statement
testIf := 2
if testIf < 3 {
    fmt.Println("run")
} else if testIf >5 {
    fmt.Println("jump")
} else {
    fmt.Println("sit")
}

// switch case
test_switch := 8
switch  {
case test_switch < 2:
	fmt.Println("This is less than 2")
	fallthrough
case test_switch < 5 :
	fmt.Println("This is less than 5")
	fallthrough
case test_switch > 3 && test_switch < 5:
	fmt.Println("This is greater than 3")
	fallthrough
case test_switch == 5:
	fmt.Println("This is five ")

default :
	fmt.Println("default")
	
}

test_string := "Hello patrick"
for _, char := range test_string { 
	fmt.Printf("%c\n", char)
}

// while loop
test_while := 5
for test_while < 10 {
	fmt.Println("While loop")
	test_while++
}

array := [2][2]int{{2 ,3},{4, 5}}
fmt.Printf("%T\n", array)
for _, value := range array {
	// fmt.Println(i,value)
	for i, details := range value {
		fmt.Println( i,details)
	}
}

array2 := [...][3]int{{3, 2, 1},{4, 5, 6},{5,6,7}}
fmt.Println(array2)
sl := array2[1:]
// sl[0] = [1, 0, 0]
fmt.Println(sl)
fmt.Println(sl[0])
// sl[0] = []int{3,5,8}
fmt.Println(sl[0])


}   