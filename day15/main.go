/*package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Map

	m.Store("num", 15)

	item, ok := m.Load("num")
	if ok {
		num := item.(int)
		fmt.Println("The number is:", num)
	}

	m.Store(7, "Friday")

	if item, ok := m.Load(7); ok {
		day := item.(string)
		fmt.Println("The day is:", day)
	}

	m.Delete("num")


}
*/

/*
package main

import (
	"fmt"
	"sync"
)

func main() {

	var config sync.Map

	var wg sync.WaitGroup

	wg.Add(3)

	//there is no need to add a goroutine parent for the other goroutine
	//نیازی به یک گوروتین مادر برای اعمال گوروتین های بی نام نیست
	go func() {
		config.Store("Theme", "Dark Mode")
		wg.Done()
	}()

	go func() {
		config.Store("volume", 85)
		wg.Done()
	}()

	go func() {
		config.Store("Wifi", true)
		wg.Done()
	}()

	wg.Wait()

	test, ok := config.Load("Theme")
	if ok {
		printer := test.(string)
		fmt.Println("theme is:", printer)
	}

	test, ok = config.Load("volume")
	if ok {
		printer2 := test.(int)
		fmt.Println("volume is:", printer2)
		if printer2 > 50 {
			fmt.Println("Volume is High:", printer2)
		}
	}

}
*/

/*
//Waiting for Result Pattern
package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)

		ch <- "Amirali-Nourian"

		fmt.Println("Goroutine finished!")
	}()

	fmt.Println("Waiting...")

	fullName := <-ch

	fmt.Println("My name is:", fullName)

}
*/

// Waiting for Result Pattern practice

/*
package main

import (
	"fmt"
	"math"
	"time"
)

func calculatePower(base int, power int, c chan int) {
	holder := math.Pow(float64(base), float64(power))

	c <- int(holder)
}

func main() {
	c := make(chan int)

	go calculatePower(2, 10, c)

	fmt.Println("Calculating...")

	time.Sleep(2 * time.Second)

	result := <-c

	fmt.Println("Result:", result)
}
*/

/*
//Waiting for Task
package main

import (
	"fmt"
	"sync"
)

type User struct {
	Name  string
	Email string
}

func main() {
	users := make(chan User)

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		user := <-users

		fmt.Printf("User %s registered with email %s\n", user.Name, user.Email)

		wg.Done()
	}()

	fmt.Println("Sending Task...")

	users <- User{Name: "Amirali", Email: "amirali83nori@gmail.com"}

	wg.Wait()
}
*/

/*
package main

import (
	"fmt"
	"sync"
)

func bot(names <-chan string, wg *sync.WaitGroup) {

	holder := <-names

	fmt.Printf("Hello %s, welcome to the group!", holder)
	wg.Done()

}

func main() {
	c := make(chan string)

	var wg sync.WaitGroup

	wg.Add(1)

	go bot(c, &wg)

	fmt.Println("Bot is ready...")

	c <- "Amirali"

	wg.Wait()

}
*/

/*
//pooling

package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan string)

	for i := 1; i <= 5; i++ {
		go func(workerID int) {
			for name := range c {
				fmt.Printf("Worker %d received: %s\n", workerID, name)
				time.Sleep(time.Second)
			}
		}(i)
	}

	names := []string{"A", "B", "C", "D", "E", "F"}
	for _, n := range names {
		c <- n
	}
	time.Sleep(3 * time.Second)
}
*/

/*
package main

import (
	"fmt"
	"sync"
	"time"
)

func Worker(ID int, Jobs <-chan string, wg *sync.WaitGroup) {

	for value := range Jobs {
		fmt.Printf("Worker %d started processing %v\n", ID, value)
		time.Sleep(1 * time.Second)
		fmt.Printf("Worker %d finished %v\n", ID, value)
	}
	wg.Done()
}

func main() {
	Jobs := make(chan string)

	var wg sync.WaitGroup

	wg.Add(3)

	for i := 1; i <= 3; i++ {
		go Worker(i, Jobs, &wg)

	}

	images := []string{"img 1", "img 2", "img 3", "img 4", "img 5"}
	for _, val := range images {
		Jobs <- val
	}
	close(Jobs)

	wg.Wait()
}
*/

/*
// drop
package main

import (
	"fmt"
	"time"
)

func baker(jobs <-chan string) {
	for job := range jobs {
		fmt.Printf("Baker started baking for %s...\n", job)

		time.Sleep(2 * time.Second)

		fmt.Printf("%s's bread is ready!\n", job)
	}
}

func main() {

	jobs := make(chan string, 3)

	for i := 1; i <= 10; i++ {
		customer := fmt.Sprintf("Customer ID:%v", i)

		select {
		case jobs <- customer:
			fmt.Println("✅", customer, "waiting in queue.")
		default:
			fmt.Println("⛔️ Bakery is full!", customer, "left (Dropped).")
		}
		time.Sleep(500 * time.Millisecond)
	}
	time.Sleep(4 * time.Second)
	fmt.Println("--- Shop Closed ---")
}
*/

/*
package main

import (
	"fmt"
	"time"
)

func processor(request <-chan int) {
	for val := range request {
		fmt.Println("Processing req ...", val)
		time.Sleep(500 * time.Millisecond)
		fmt.Printf("✅ Req %v Done\n", val)
	}
}

func main() {
	c := make(chan int, 5)

	go processor(c)

	for i := 1; i <= 20; i++ {
		select {
		case c <- i:
			fmt.Printf("📥 Req %v Accepted\n", i)
		default:
			fmt.Printf("⛔️ Req %v Dropped (Server Busy)\n", i)
		}
		time.Sleep(50 * time.Millisecond)
	}

	time.Sleep(3 * time.Second)
}
*/

package main

import (
	"fmt"
	"time"
)

func processTransaction(resultChan chan<- string) {
	time.Sleep(3 * time.Second)

	holder := fmt.Sprintln("Transaction Successful")

	resultChan <- holder
}

func main() {
	c := make(chan string)

	go processTransaction(c)

	select {
	case res := <-c:
		fmt.Println("✅ Success", res)
	case <-time.After(2 * time.Second):
		fmt.Println("❌ Timeout! Transaction Cancelled")
	}
}
