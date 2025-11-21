package main

import "fmt"

func main() {
	copy_slice()
	ranging()

	fmt.Println(sum(1, 2, 3))
	fmt.Println(sum(5, 10))

	maping_test()
}

func copy_slice() {
	nums := []int{1, 2, 3, 4}
	n := copy(nums, []int{5, 6})
	fmt.Println("n:", n, "nums:", nums)
}

func ranging() {
	sliceTest := make([]int, 5)

	copy(sliceTest, []int{10, 20, 30, 40, 50})
	for i := range sliceTest {
		fmt.Println("index is:", i, "value is:", sliceTest[i])
	}
	fmt.Println("Original slice value", sliceTest)
	for _, value := range sliceTest {
		value += value
		//This only changes the local copy 'value', NOT the slice itself
	}
	fmt.Println("no changes applied with changing value of slice (it's just a copy)", sliceTest)

	for i := range sliceTest {
		// در اینجا با استفاده از اندیس، مستقیماً مقدار داخل اسلایس را تغییر می‌دهیم
		sliceTest[i] = sliceTest[i] + 1

	}
	// Now the slice is actually modified
	fmt.Println("slice values has changed because of using index", sliceTest)

}

//variadic functions

func sum(numbers ...int) int {
	var sum int
	for _, n := range numbers {
		sum += n
	}
	return sum
}

func maping_test() {
	ages := map[string]int{}

	ages["Amirali"] = 25
	ages["Mehrsa"] = 19
	ages["Parmis"] = 12
	ages["koroush"] = 13

	fmt.Println(ages)

	delete(ages, "Parmis")
	//deleting a key/value from map
	fmt.Println(ages)

	for name, age := range ages {
		fmt.Println(name, ":", age)
	}

	age, ok := ages["Maryam"]
	if ok {
		fmt.Println("Maryam exists, age:", age)
	} else {
		fmt.Println("Maryam is not in map")
	}

	clear(ages) // مپ کاملاً خالی می‌شود
	fmt.Println("map cleared=", ages)

}
