/*

package main

import (
	"fmt"
	"time"
)

func baker(shelf chan string) {
	for i := 1; i <= 5; i++ {
		fmt.Printf("👨‍🍳 Baking bread [%d]...\n", i)

		// 1. Try to put on shelf (This blocks if buffer is full)
		// ۱. تلاش برای گذاشتن در قفسه (اگر پر باشد اینجا قفل می‌شود)
		shelf <- fmt.Sprintf("Bread %d", i)

		// 2. If we passed the line above, it means success
		// ۲. اگر از خط بالا رد شدیم، یعنی موفق شدیم
		fmt.Printf("✅ Placed bread [%d] on shelf\n", i)
	}
}

func main() {
	// Buffered channel with capacity 3
	c := make(chan string, 3)

	go baker(c)

	// Give baker time to fill the shelf
	// فرصت به نانوا برای پر کردن قفسه
	fmt.Println("... Customer is walking to the shop (3 seconds) ...")
	time.Sleep(3 * time.Second)

	fmt.Println("\n--- Customer Arrived ---")

	for i := 1; i <= 5; i++ {
		// 3. Take from shelf (Unblocks the baker)
		// ۳. برداشتن از قفسه (باعث می‌شود نانوا آزاد شود)
		bread := <-c

		fmt.Printf("😋 Customer bought [%s]\n", bread)
		time.Sleep(1 * time.Second)
	}
}


*/

/*
package main

import "fmt"

//directional channel

func SendData(c chan<- string) {
	c <- "Hello"
}

func readData(c <-chan string) {
	msg := <-c
	fmt.Println(msg)
}

func main() {
	c := make(chan string)

	go SendData(c)

	readData(c)

}

*/

/*
package main

import "fmt"

func Throw(c chan<- string) {
	fmt.Println("Hand:Throw the ball...")
	c <- "Ball"
}

func Catch(c <-chan string) {
	item := <- c

	fmt.Println("Glove: Caught the ", item)
}

func main() {
	c := make(chan string)

	go Throw(c)

	Catch(c)
}
*/

/*
package main

import "fmt"

func ping(c chan<- string) { //send
	fmt.Println("player1 throw the ball....")

	c <- "Ball"
}

func pong(c <-chan string) { //receive
	item := <-c

	fmt.Println("player2 Received the", item)
}

func main() {
	c := make(chan string)

	go ping(c)

	pong(c)
}
*/

/*
package main

import "fmt"

func main() {

	c := make(chan string, 2)

	c <- "Hello"
	c <- "World!"

	close(c)

	for msg := range c {
		fmt.Println(msg)
	}
}
*/

/*

package main

import (
	"fmt"
	"sync"
)

func computer(jobs chan<- string) { //sender
	for i := 1; i <= 10; i++ {
		msg := fmt.Sprintf("File %d", i)

		jobs <- msg

		fmt.Println("Computer sent:", msg)
	}
	close(jobs)
}

func printer(jobs <-chan string, wg *sync.WaitGroup) { //receiver
	defer wg.Done()

	for val := range jobs {
		fmt.Printf("🖨️ Printing %v...\n", val)
	}

}

func main() {
	var wg sync.WaitGroup

	c := make(chan string)

	wg.Add(1)

	go computer(c)

	go printer(c, &wg)

	wg.Wait()

	fmt.Println("All documents printed.")
}
*/

package main

import (
	"fmt"
	"sync"
)

func orderTaker(rawOrders chan<- string) { //sender
	for i := 1; i <= 5; i++ {
		holder := fmt.Sprintf("burger %v", i)

		rawOrders <- holder
	}
	close(rawOrders)
}

func chef(rawOrders <-chan string, cookedFood chan<- string) { //receiver
	for order := range rawOrders {
		holder := fmt.Sprintf("Cooked %v", order)

		cookedFood <- holder

	}
	close(cookedFood)
}

func waiter(cookedFood <-chan string, wg *sync.WaitGroup) { //reciever
	for server := range cookedFood {
		fmt.Printf("Serving: %v\n", server)
	}
	defer wg.Done()
}

func main() {
	rawChan := make(chan string)
	cookedChan := make(chan string)

	var wg sync.WaitGroup

	wg.Add(1)

	go orderTaker(rawChan)
	go chef(rawChan, cookedChan)

	go waiter(cookedChan, &wg)

	wg.Wait()

	fmt.Println("Restaurant Closed.")
}

// force update commit message
