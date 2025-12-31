/*

//Empty interface / any

package main

import "fmt"

type Empty interface{}

func PrintAnything(v any) {
	fmt.Println("Received:", v)
}

func main() {
	PrintAnything(100)
	PrintAnything("Hello")
	PrintAnything(true)
	PrintAnything(3.14)

}


*/

/*
package main

import "fmt"

func main() {
	var anySlice []any
	anySlice = append(anySlice, 1, "Golang", true)

	for _, v := range anySlice {
		fmt.Println(v)
	}
}
*/

/*

package main

import "fmt"

func Process(data any) {
	valInt, IsInt := data.(int)

	if IsInt {
		fmt.Println("It's a number:", valInt*2)
		return
	}

	valStr, isStr := data.(string)
	if isStr {
		fmt.Println("This is a text:", valStr)
		return
	}

	fmt.Println("Unknown type!")
}

func main() {
	Process(10)
	Process("Test")
	Process(true)
}
*/

/*

package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func (u User) String() string {
	return fmt.Sprintf("👤 Name: %s | 🎂 Age: %d years old", u.Name, u.Age)
}

func main() {
	u := User{
		Name: "Amirali",
		Age:  25,
	}
	fmt.Println(u)
}
*/
/*
package main

import "fmt"

type Item struct {
	Name  string
	Price int
}

func (i Item) String() string {
	return fmt.Sprintf("📦 Product: [%s] ......... [%d] Toman", i.Name, i.Price)
}

func main() {
	var prodct []Item
	prodct = append(prodct, Item{
		Name:  "Laptop",
		Price: 90000000,
	}, Item{
		Name:  "Mouse",
		Price: 900000,
	})
	for _, val := range prodct {
		fmt.Println(val)
	}
}
*/

/*
//season-8

//io.write
package main

import "fmt"

type CounterWriter struct {
	Count int
}

func (cw *CounterWriter) Write(p []byte) (n int, err error) {
	length := len(p)
	cw.Count += length

	return length, nil
}

func main() {
	cw := &CounterWriter{}

	fmt.Fprintf(cw, "Hello Go")
	fmt.Fprintf(cw, "World")

	fmt.Println("Total Bytes Written:", cw.Count)
}
*/

/*

//io.reader
package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	source := strings.NewReader("Hello, World!")

	buffer := make([]byte, 4)

	for {
		n, err := source.Read(buffer)
		fmt.Printf("Read %d bytes: %q\n", n, buffer[:n])
		if err == io.EOF {
			fmt.Println("--- End Of File ---")
			break
		}
	}
}
*/

/*
package main

import (
	"io"
	"os"
	"strings"
)

func main() {
	source := strings.NewReader("This is a stream of data.\n")

	io.Copy(os.Stdout, source)
}
*/

/*

package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	txtholder := strings.NewReader("Go is Great")

	buffer := make([]byte, 2)

	for {
		n, err := txtholder.Read(buffer)

		for i := 0; i < n; i++ {
			if buffer[i] == ' ' {
				buffer[i] = '-'
			}
		}
		fmt.Print(string(buffer[:n]))
		if err == io.EOF {
			break
		}
	}
	fmt.Println()
}
*/

/*
package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	source := strings.NewReader("The quick brown fox jumps over the lazy dog.")
	buffer := make([]byte, 3)
	counter := 0
	for {
		n, err := source.Read(buffer)

		for i := 0; i < n; i++ {
			if buffer[i] == 'e' {
				counter++
			}

		}
		if err == io.EOF {
			break
		}
	}
	fmt.Println("total of e(s):", counter)
}
*/

/*
package main

import (
	"fmt"
)

type ConsoleWriter struct {
}

func (cw ConsoleWriter) Write(p []byte) (n int, err error) {
	fmt.Println(string(p))
	return len(p), nil
}

func main() {
	writer := ConsoleWriter{}

	writer.Write([]byte("Hello Go Interface!"))


	fmt.Println()

	fmt.Fprintln(writer, "This is printed via Fprintln using our custom writer!")
}
*/

package main

import (
	"fmt"
	"io"
)

type MyStringReader struct {
	content  string
	position int
}

func (r *MyStringReader) Read(p []byte) (n int, err error) {
	if r.position >= len(r.content) {
		return 0, io.EOF
	}
	remainingData := r.content[r.position:]
	n = copy(p, []byte(remainingData))
	r.position += n
	return n, nil
}

func main() {
	reader := &MyStringReader{content: "Hello Go World!"}

	buffer := make([]byte, 4)

	for {

		n, err := reader.Read(buffer)
		if err != nil && err != io.EOF {
			fmt.Println("Error:", err)
			break
		}
		fmt.Printf("Read %d bytes: %q\n", n, buffer[:n])
		if err == io.EOF {
			fmt.Println("End of stream reached (EOF).")
			break
		}
	}
}

// force update commit message
