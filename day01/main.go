package main

import "fmt"

func main() {
	//fmt use
	fmt.Println("Hello, World!")
	//multiple var declaration without type
	var (
		age  = 25
		name = "Amirali"
	)
	fmt.Println("My name is, ", name, "I'm ", age, " years old!")

	//multiple var declaration with type
	var (
		age_2  int    = 25
		name_2 string = "Amirali"
	)
	fmt.Println("My name is, ", name_2, "I'm ", age_2, " years old!")

	//variable declaration with out var
	//this way of declaring you can only reach it in that function you created the var (not global)
	name_3 := "Amirali"
	age_3 := 25
	fmt.Println("My name is, ", name_3, "I'm ", age_3, " years old!")

	//shortest way to declare var
	fName, My_age := "Amirali", 25
	fmt.Println("My name is, ", fName, "I'm ", My_age, " years old!")

	//finding max and min with builtin function
	a, b := 20, 30
	fmt.Println("max is: ", max(a, b))
	fmt.Println("min is: ", min(a, b))

	//variable casting
	a2, b2 := 20, 18.75
	fmt.Println(a2, b2)
	fmt.Println(float32(a2), int8(b2))

	//iota and constant declaration
	//iota only use in const
	//iota starts from 0
	const (
		stat_1 = iota + 1
		stat_2
		stat_3
		stat_4
	)
	fmt.Println("stat_4: ", stat_4)

	//Underline in here means Blank Identifier
	//means const is declared but no name
	const (
		reason_1 = iota + 1
		_
		_
		reason_4
	)
	fmt.Println("reason_4: ", reason_4)

}



// force update commit message
