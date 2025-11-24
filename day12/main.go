/*package main

import (
	"fmt"
	"time"
)

//channel

func HelloWorld(c chan string) {
	time.Sleep(5 * time.Second)

	c <- "Hello World!"
}

func main() {
	c := make(chan string)

	go HelloWorld(c)

	fmt.Println("Waiting...")

	message := <-c

	fmt.Println("Received:", message)
}
*/

/*
package main

import "fmt"

func multiply(num1, num2 int, c chan int) {
	c <- num1 * num2
}

func main() {
	c := make(chan int)

	go multiply(5, 10, c)

	Result := <-c

	fmt.Println("Result is:", Result)
}
*/

/*
package main

import (
	"fmt"
	"time"
)

func RUN(name string, c chan string) {
	fmt.Printf("🏃 %s is waiting for the baton...\n", name)

	baton := <-c

	fmt.Printf("⚡ %s got the baton (%s) and started running!\n", name, baton)

	time.Sleep(1 * time.Second)

	fmt.Printf("✅ %s finished and passed the baton.\n", name)

	c <- baton
}

func main() {
	c := make(chan string)

	go RUN("Runner 1", c)

	go RUN("Runner 2", c)

	time.Sleep(100 * time.Millisecond)
	fmt.Println("🔫 STARTING RACE! Passing the baton to the first runner...")

	c <- "Gold Baton"

	finalBaton := <-c

	fmt.Println("🏁 Race finished! Final baton received:", finalBaton)
}
*/

/*
package main

import (
	"fmt"
	"time"
)

func makeCoffee(order string, tray chan string) {
	fmt.Println("☕ Barista started making", order)

	time.Sleep(1 * time.Second)

	fmt.Printf("✨ %s is ready!\n", order)

	tray <- order + " (Ready)"
}

func main() {
	tray := make(chan string)

	orders := []string{"Cappuccino", "Latte", "Espresso"}

	fmt.Println("--- Shop Open ---")

	for _, order := range orders {
		go makeCoffee(order, tray)
	}

	for i := 0; i < len(orders); i++ {
		serve := <-tray

		fmt.Println("Serving to customer:", serve)
	}

	fmt.Println("--- All orders served ---")
}
*/

/*package main

import "fmt"

func main() {
	c := make(chan int, 2) // ظرفیت فقط ۲ تاست

	c <- 1 // ✅ اوکی (بافر: ۱/۲)
	c <- 2 // ✅ اوکی (بافر: ۲/۲) - پر شد!

	go func() {
		c <- 3
	}()

	fmt.Println(<-c)
}
*/

package main

import (
	"fmt"
	"time"
)

func baker(shelf chan string) {
	for i := 1; i <= 5; i++ {
		fmt.Printf("👨‍🍳 Baking bread [%d]...\n", i)

		shelf <- fmt.Sprintf("Bread %d", i)

		fmt.Printf("✅ Placed bread [%d] on shelf\n", i)

	}
}

func main() {
	c := make(chan string, 3)

	go baker(c)

	fmt.Println("... Customer is walking to the shop (3 seconds) ...")

	time.Sleep(3 * time.Second)

	fmt.Println("\n--- Customer Arrived ---")

	for i := 1; i <= 5; i++ {
		bread := <-c
		fmt.Printf("😋 Customer bought [%v]\n", bread)
		time.Sleep(1 * time.Second)
	}
}
