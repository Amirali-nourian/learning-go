/*package main

import (
	"fmt"
	"sync"
	"time"
)

func Process(item string) {
	time.Sleep(1 * time.Second)
	fmt.Println("Processed:", item)
}

func main() {

	//روش غلط در زیر استفاده شده
	fmt.Println("start processing ...")
	go Process("Image.jpg")
	time.Sleep(2 * time.Second)
	fmt.Println("Main function continues...")
	fmt.Println("Main function finished.")

	//updated with using WaitGroup

	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)

		msg := fmt.Sprintf("SMS %d", i+1)
		go func(m string) {
			defer wg.Done()

			SendSMS(m)
		}(msg)

	}

	fmt.Println("Waiting for SMS tasks...")

	wg.Wait()
	fmt.Println("All tasks completed successfully!")

	var wtg sync.WaitGroup

	wtg.Add(2)

	go PrintNumbers(&wtg)
	go PrintLetters(&wtg)

	wtg.Done()
	wtg.Wait()
	fmt.Println("cointing is done")
}

//end of main

func SendSMS(msg string) {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Sending:", msg)
}

func PrintNumbers(wtg *sync.WaitGroup) {
	defer wtg.Done()

	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}
	time.Sleep(200 * time.Millisecond)
}

func PrintLetters(wtg *sync.WaitGroup) {
	defer wtg.Done()

	for i := 'A'; i <= 'E'; i++ {
		fmt.Println(string(i))
	}
	time.Sleep(200 * time.Millisecond)

}
*/

/*
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wtg sync.WaitGroup

	wtg.Add(2)

	go PrintNumbers(&wtg)
	go PrintLetters(&wtg)

	wtg.Wait()
	fmt.Println("counting is done")
}

func PrintNumbers(wtg *sync.WaitGroup) {
	defer wtg.Done()
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}
	time.Sleep(200 * time.Millisecond)
}

func PrintLetters(wtg *sync.WaitGroup) {
	defer wtg.Done()
	for i := 'A'; i <= 'E'; i++ {
		fmt.Println(string(i))
	}
	time.Sleep(200 * time.Millisecond)
}
*/

/*
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(3)

	go PrintNumbers(&wg)
	go PrintLetters(&wg)
	go PrintSymbols(&wg)

	wg.Wait()
	fmt.Println("All tasks completed successfully!")
}

func PrintNumbers(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}
	time.Sleep(200 * time.Millisecond)
}

func PrintLetters(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 'A'; i <= 'E'; i++ {
		fmt.Println(string(i))
	}
	time.Sleep(200 * time.Millisecond)

}

func PrintSymbols(wg *sync.WaitGroup) {
	defer wg.Done()

	for _, symbol := range []rune{'@', '#', '$', '%'} {
		fmt.Println(string(symbol))
	}
	time.Sleep(200 * time.Millisecond)

}
*/

/*

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	files := []string{"Movie.mp4", "", "Music.mp3", "Photo.jpg"}

	fmt.Println("--- Start Download Manager ---")

	for _, file := range files {

		if file == "" {
			fmt.Println("Error: Skipping invalid file link.")
			continue
		}

		wg.Add(1)
		go downloadFile(file, &wg)
	}
	wg.Wait()
	fmt.Println("--- All downloads finished successfully ---")
}

func downloadFile(url string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Downloading %s...\n", url)
	time.Sleep(1 * time.Second)
	fmt.Printf("%s Downloaded.\n", url)
}
*/

/*

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	urls := []string{"google.com", "github.com", "", "stackoverflow.com"}

	for _, url := range urls {
		if url == "" {
			fmt.Println("Skipping invalid URL")
			continue
		}
		wg.Add(1)
		go checkStatus(url, &wg)
	}
	wg.Wait()
	fmt.Println("Finished checking all sites")
}

func checkStatus(url string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Checking %s ...\n", url)
	time.Sleep(1 * time.Second)
	fmt.Println(url, "is UP!")
}
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)

	go HelloWorld(c)

	fmt.Println("Waiting...")

	message := <-c
	fmt.Println("Received:", message)
}

// channel
func HelloWorld(c chan string) {
	time.Sleep(5 * time.Second)
	c <- "Hello World!"
}

// force update commit message
