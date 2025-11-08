package main

import (
	"fmt"
)

func main() {
	fmt.Println("----------------------------------------------------")

	fmt.Println("self exercise")
	scores := []int{95, 82, 74, 59, 68, 100, 45}
	for _, num := range scores {
		grade := getGrade(num)
		fmt.Printf("Score: %d, Grade: %s\n", num, grade)
	}

	fmt.Println("----------------------------------------------------")
	fmt.Println("Nested loop to show Multiplication Table")
	nested_for_loop()
	fmt.Println("----------------------------------------------------")
	fmt.Println("reverse loop")
	reverse_loop()
	fmt.Println("----------------------------------------------------")
	fmt.Println("infinity loop")
	//for stopping the infinit loop with out engagement I used break
	infinit_loop()
	fmt.Println("----------------------------------------------------")
	fmt.Println("continue in for")
	//continue life span is in the scope that is used
	loop_continue()
	fmt.Println()
	fmt.Println("----------------------------------------------------")
	fmt.Println("do loop with for")
	do_with_for()
	fmt.Println("----------------------------------------------------")
	fmt.Println("do while loop with for")
	do_while()
	fmt.Println("----------------------------------------------------")
	fmt.Println("goto use")
	go_to()
	fmt.Println("----------------------------------------------------")
	fmt.Println("function syntax")
	function_syntax()
}

func getGrade(Score int) string {
	switch {
	case Score >= 90 && Score <= 100:
		return "A"
	case Score >= 80 && Score <= 89:
		return "B"
	case Score >= 70 && Score <= 79:
		return "C"
	case Score >= 60 && Score <= 69:
		return "D"
	case Score < 60:
		return "F"
	default:
		return "invalid Score"
	}
}

func nested_for_loop() {
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			fmt.Print(i*j, "\t")
		}
		fmt.Println()
	}
}
func reverse_loop() {
	for i := 8; i >= 1; i-- {
		for j := 0; j < i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

func infinit_loop() {
	for {
		fmt.Println("Hello, World!")
		break
	}
	fmt.Println("loop finished")

}

func loop_continue() {
	for i := 1; i <= 5; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Print(i, " ")
	}
}

func do_with_for() {
	var i int
	for {
		i++
		if i > 5 {
			break
		}
		fmt.Print(i, "")
	}
}

func do_while() {
	var i int
	for i < 5 {
		i++
		fmt.Print(i, " ")
	}
}

func go_to() {
	goto B
	fmt.Println("A")
B:
	fmt.Println("B")
}

func function_syntax() {
	fmt.Println(`functionName(parameter1 type, parameter2 type) returnType`)
}
