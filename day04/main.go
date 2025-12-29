package main

import "fmt"

const pi = 3.14

var num = 10

func main() {
	//remember to use a variable(s) to assign a return unction value
	c, _ := AreaAndCircumference(4)
	fmt.Println("Circumference of circle with radius of 4 is: ", c)
	_, a := AreaAndCircumference(7)
	fmt.Println("Area circle with radius of 7 is:", a)
	var num = 20
	fmt.Println(num)
	fmt.Println("------------------------------------")
	fmt.Println("anonymous function")
	add := func(a, b int) int {
		return a + b
	}
	fmt.Println(add(3, 2))
	fmt.Println("------------------------------------")
	fmt.Println("anonymous function without name")
	//Immediately Invoked Function Expression اجرای بلافاصله (بدون ذخیره در متغیر)
	func(name string) {
		fmt.Println("Hello", name)
	}("Amirali")
}

//if you set name(s) for return value in functions you don't need to call the names of it and simply use 'return'
//better not to use this way it's not efficient
func AreaAndCircumference(r float64) (circumference float64, area float64) {
	circumference = 2 * r * pi
	area = r * r * pi
	return
}


