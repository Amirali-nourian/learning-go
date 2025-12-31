package main

import "fmt"

func main() {
	using_slice()

	adding_value_to_slice_with_for()

	slicing_parting()

	copying_slice()

	ranger()

	fmt.Println(sum(1, 2, 3))
	fmt.Println(sum(5, 10, 15, 20, 25))

	//sending a slice to variadic func
	test := []int{10, 20, 30, 40, 50}
	fmt.Println(sum(test...))

	show("Amirali", "Football", "Music", "gaming")

	fmt.Println(mul(2, 3, 4))
	fmt.Println(mul(1.5, 10))
}

func adding_value_to_slice_with_for() {
	numbers := make([]int, 0, 4)
	for i := 0; i < 6; i++ {
		fmt.Println("len:", len(numbers), "cap:", cap(numbers), "numbers:", numbers)
		numbers = append(numbers, i)
	}
}

func using_slice() {
	//slice before and after appending
	cities := []string{"Tehran", "Shiraz", "Tabriz"}
	fmt.Println("Length before append: ", len(cities))
	fmt.Println("Capacity before append: ", cap(cities))
	cities = append(cities, "Mashhad")
	fmt.Println("Length after append: ", len(cities))
	fmt.Println("Capacity after append: ", cap(cities))
	fmt.Println(cities[2])

	part := cities[1:3]
	fmt.Println("Slicing of a slice", part)
	part[0] = "Qom"
	fmt.Println(cities)
}

func slicing_parting() {
	nums := []int{10, 20, 30, 40, 50}
	fmt.Println("values of slice before slicing: ", nums)
	partofslice := nums[1:3]
	fmt.Println("part of slice we parted: ", partofslice)
	partofslice[0] = 99
	fmt.Println("values of slice after slicing and changing value: ", nums)

}

func copying_slice() {
	nums := make([]int, 4)
	fmt.Println("slice before copying", nums)
	copy(nums, []int{1, 2, 3, 4, 5})
	//copy مقدارها را تا جایی می‌ریزد که جای خالی وجود داشته باشد
	//copy نمی‌تواند طول و ظرفیت اسلایس را افزایش دهد
	fmt.Println("slice after copying", nums)
}

func ranger() {
	var nums = []int{10, 20, 30, 40, 50}
	for range nums {
		fmt.Println("I love Go!")
	}
	fmt.Println("Values of slice is :", nums)
	clear(nums)
	fmt.Println("Values of slice after clearing it :", nums)

}

//variadic function

func sum(nums ...int) int {
	total := 0

	for _, v := range nums {
		total += v
	}
	return total
}

func show(name string, hobbies ...string) {
	fmt.Println(name, hobbies)

}

func mul(numbers ...float64) float64 {
	total := 1.0
	for _, value := range numbers {
		total *= value
	}
	return total
}



// force update commit message
