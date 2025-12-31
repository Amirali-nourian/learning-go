/*package main

import (
	"fmt"
	"time"
)

func mockServer(c chan<- string, name string, delay time.Duration) {
	time.Sleep(delay)
	c <- fmt.Sprintf("%s replied", name)

}

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go mockServer(c1, "USA server1", 1*time.Second)
	go mockServer(c2, "EU server2", 3*time.Second)

	select {
	case msg := <-c1:
		fmt.Printf("Winner: %s", msg)
	case msg2 := <-c2:
		fmt.Printf("Winner: %s", msg2)

	}
}
*/

//test

/*
package main

import (
	"fmt"
	"time"
)

func activateSensor(c chan<- string, location string, delay time.Duration) {
	time.Sleep(delay)

	c <- fmt.Sprintf("Alarm! Motion detected at %s", location)
}

func main() {
	DoorChan := make(chan string)
	WindowChan := make(chan string)

	go activateSensor(DoorChan, "Door location", 2*time.Second)
	go activateSensor(WindowChan, "Window location", 5*time.Second)

	select {
	case msg1 := <-DoorChan:
		fmt.Println("Warning:", msg1)
	case msg2 := <-WindowChan:
		fmt.Println("Warning:", msg2)
	}
}
*/

/*
package main

import (
	"fmt"
	"time"
)

func student(answerChan chan<- string, delay time.Duration, answer string) {
	time.Sleep(delay)

	answerChan <- fmt.Sprintf("Student chose %s", answer)

}

func main() {
	c := make(chan string)

	go student(c, 4*time.Second, "Option B")

	select {
	case message := <-c:
		fmt.Println("Answer accepted: ", message)

	case <-time.After(3 * time.Second):
		fmt.Println("Time's up! You failed.")
	}
}
*/

/*

package main

import (
	"fmt"
	"sync"
)

type BankAccount struct {
	balance int
	mu      sync.Mutex
}

func (b *BankAccount) Deposit(amount int) {
	b.mu.Lock()

	b.balance += amount

	b.mu.Unlock()
}

func (b *BankAccount) GetBalance() {
	b.mu.Lock()

	fmt.Println("total of account is:", b.balance)

	b.mu.Unlock()

}

func main() {
	account := BankAccount{
		balance: 0,
	}

	var wg sync.WaitGroup

	for i := 1; i <= 1000; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			account.Deposit(1)
		}()
	}

	wg.Wait()
	account.GetBalance()
}
*/

package main

import (
	"fmt"
	"strings"
	"sync"
)

type Post struct {
	Content string
	Likes   int
	mu      sync.RWMutex
}

func (p *Post) ReadPost() {

	p.mu.RLock()

	holder := p.Content
	fmt.Println("Content is:", holder)

	holder2 := p.Likes
	fmt.Println("Likes of content is:", holder2)

	p.mu.RUnlock()
}

func (p *Post) LikePost() {
	p.mu.Lock()

	p.Likes++

	p.mu.Unlock()
}

func main() {

	post := Post{
		Content: `Hi,
		This is my travel to Egypt.`,
		Likes: 2000,
	}

	var wg sync.WaitGroup

	wg.Add(5)

	for i := 1; i <= 5; i++ {
		go func() {
			go post.LikePost()
			wg.Done()
		}()

	}

	wg.Add(100)

	for i := 1; i <= 100; i++ {
		go func() {
			post.ReadPost()
			wg.Done()
		}()
	}

	wg.Wait()

	separator := strings.Repeat("=", 50)
	fmt.Println("\n" + separator)
	fmt.Println("📊 FINAL RESULT:")
	fmt.Println(separator)
	fmt.Printf("Content: %s\n", post.Content)
	fmt.Printf("Final Likes: %d\n", post.Likes)
	fmt.Println(separator)
}


// force update commit message
