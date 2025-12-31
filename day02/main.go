package main

import "fmt"

func main() {
	long_string_with_prefered_structure()
	switch_use_example()
	advanced_switch_example()
	list_case_switch()
	impliment_while_with_for()
	//use fmt.Println in printing the function value
	fmt.Println(circle_Circumference_and_Area(5.0))
	fmt.Println(naked_return(3, 5))
}

func long_string_with_prefered_structure() {
	fmt.Println(`this is a long text
	based on 
	how
	I want
	to be shown`)
}

func switch_use_example() {
	//In an Expression Switch, the type of the value in each case must be compatible with or identical to the type of the variable preceding the switch.
	screw := "philips"
	switch screw {
	case "philips":
		fmt.Println("you need a philips screwdriver(positive head)")
	case "flathead":
		fmt.Println("you need a flathead screwdriver(negative head)")
	case "torx":
		fmt.Println("you need a torx screwdriver(star head)")
	default:
		fmt.Println("you don't need a screwdriver")
	}

}

func advanced_switch_example() {
	//this switch is a replacement of if else if
	//it's a Boolean Expression
	//Expressionless Switch
	age := 17
	switch {
	case age < 13:
		fmt.Println("child")
	case age >= 13 && age <= 17:
		fmt.Println("teen")
	default:
		fmt.Println("elder")
	}
}

func list_case_switch() {
	//it's the same as normal switch BUT you can use multiple case too.(even in switch_use_example() function you can use this too)
	day := 3
	fmt.Print("Today is: ")
	switch day {
	case 1:
		fmt.Println("Saturday")
	case 2:
		fmt.Println("Sunday")
	case 3:
		fmt.Println("Monday")
	case 4:
		fmt.Println("Tuesday")
	case 5:
		fmt.Println("Wednesday")
	case 6, 7:
		fmt.Println("Thursday or Friday")
	default:
		fmt.Println("Unknown Day!")
	}
}

func impliment_while_with_for() {
	number := 100.0
	target := 5.0

	fmt.Println("Starting value:", number)
	fmt.Println("Target value:", target)
	fmt.Println("---")

	for number >= target {
		fmt.Printf("Current number is %.2f, still >= %.2f\n", number, target)

		number = number / 2
	}

	fmt.Println("---")
	fmt.Printf("Loop finished. Final number is %.2f\n", number)
}

func circle_Circumference_and_Area(radius float64) (float64, float64) {
	const (
		pi = 3.14
	)
	return 2 * pi * radius, pi * radius * radius
}

func naked_return(a, b int) (result int) {
	result = ((a + b) * 2) * 5
	return
}



// force update commit message
