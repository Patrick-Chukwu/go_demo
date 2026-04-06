package main

import (
	"errors"
	"fmt"
)


func main() {
	fmt.Println("Hello Test file")
	temperatureInt := 88
	temperatureFloat := float64(temperatureInt)

	fmt.Println(temperatureFloat)
	fmt.Printf("%.2f\n", temperatureFloat )

	s := fmt.Sprintf("I am %s for this things", "way too old")
	fmt.Println(s)

	if height := 18 ; height >= 18 {
		fmt.Println("You're old enough")
	} else {
		fmt.Println("You're not old enough")

	

	}

	result1 := sub(6,9)
	fmt.Println(result1)
	mul, div, err := calc(81, 9)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("multiplication is %v and division is %v\n", mul, div)
	}

	trucks := truck{
	bedSize : "Large",
	car : car {
		make : "Toyota",
		year : 2026,
	},
}

fmt.Println(trucks.make)
fmt.Println(trucks.year)
fmt.Println(trucks.bedSize)

r := rect{
	length : 5,
	breadth:  6,
}

fmt.Println(r.area())

}

func sub(x int, y int) int {
	z := int(y) - int(x)
	return z
}

func calc(x,y int) (mul, div int, err error) {
	if y == 0 {
		return 0,0, errors.New("Can't divide by zero")
	}

	mul = x *y
	div = x / y

	return mul, div, nil
}


//  embed struct practice

type car struct {
	make string
	year int
}

type truck struct {
	car
	bedSize string
}


type rect struct {
	length int
	breadth int
}

func (r rect) area() int {
	return r.length * r.breadth
}

