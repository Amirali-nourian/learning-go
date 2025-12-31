package main

import "fmt"

func main() {
	//array starts from 0
	//if you want an array with length of 5
	//it's 0 to 4
	//you can not call or use array[5] it's out of bound [0:5]

	var StudentScores [20]int

	StudentScores[5] = 19

	//I used for/range to print value insted I can use for/len
	for index, value := range StudentScores {
		fmt.Println("student score", index+1, ":", value)
	}

	fmt.Println("Printing array with a simple for/len")

	for i := 0; i < len(StudentScores); i++ {
		fmt.Println("student score", i+1, ":", StudentScores[i])

	}

	for i := range StudentScores {
		StudentScores[i] = 20
	}

	fmt.Println(StudentScores)
	//with using len() you can find out the length of array
	fmt.Println(len(StudentScores))

	num := [100]int{}
	for i := 0; i < 100; i++ {
		num[i] = i + 1
	}
	fmt.Println(num)
	fmt.Println("USING SLICE")

	//slice syntax
	//این اسلایس پشت صحنه 16 جای حافظه رزرو شده اما فقط 8 تاش فعال هست به همین علت هست که 8 تا خونه چاپ میشه نه 16 تا
	numbers := make([]int, 8, 16)

	fmt.Println("Length:", len(numbers))
	fmt.Println("capacity:", cap(numbers))
	//to add a value to slice you can use append BUT it will be added in the end of it.
	//append ممکن است باعث reallocation شود
	numbers = append(numbers, 2)
	fmt.Println(numbers)
	//it's still 16 cause it has capacity to add to slice
	fmt.Println("capacity:", cap(numbers))

	build()
}

func build() {
	s := make([]string, 0, 3)
	s = append(s, "A")
	s = append(s, "B")
	s = append(s, "C")
	s = append(s, "D") // ظرفیت تمام شد، پس آرایه‌ی جدید ساخته می‌شود
	fmt.Println(s)
}



// force update commit message
