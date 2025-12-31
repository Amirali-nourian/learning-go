package main

import "fmt"

func main() {
	applyFunc(5, double)
	applyFunc(5, square)

	//executor function
	ShowDetail(kebab)
	ShowDetail(pizza)
	ShowDetail(burgrer)

	//calc exercise
	calculate(10, 5, add)
	calculate(10, 5, subtract)

	//recrsive func
	fmt.Println(factorial(5))
}

func double(num int) int {
	return num * 2
}

func square(num int) int {
	return num * num
}

//return of f func(int) int is happening in func double , func square
func applyFunc(n int, f func(int) int) {
	result := f(n)
	fmt.Println(result)
}

func pizza() (string, int, string) {
	return "Pizza", 350000, "Cheese + Peproni"
}

func burgrer() (string, int, string) {
	return "Burger", 185000, "Meat + Tomato + Cheese"
}

func kebab() (string, int, string) {
	return "Kebab", 200000, "Meat + Rice"
}

func ShowDetail(f func() (string, int, string)) {
	name, price, ingredients := f()
	fmt.Println("Show food details:")
	fmt.Println("Name:", name)
	fmt.Println("Price:", price)
	fmt.Println("Ingredients:", ingredients)
	fmt.Println()
}

//calc exercise

func add(num1, num2 int) int {
	return num1 + num2
}

func subtract(num1, num2 int) int {
	return num1 - num2
}

func calculate(a int, b int, operation func(int, int) int) {
	result1 := operation(a, b)
	fmt.Println("Result:", result1)
}

//---------calc exercise

//recrsive func

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

//---------recrsive func



// force update commit message
