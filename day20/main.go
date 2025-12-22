/*
package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	reader := strings.NewReader("Hello, Go World!")

	if _, err := io.Copy(os.Stdout, reader); err != nil {
		fmt.Println("Error:", err)
	}
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
	reader := strings.NewReader("This is a test data.")

	data, err := io.ReadAll(reader)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
*/

/*

package main

import (
	"bufio"
	"os"
)

func main() {
	writer := bufio.NewWriter(os.Stdout)

	writer.WriteString("Hello ")
	writer.WriteString("Buffered ")
	writer.WriteString("World!\n")

	writer.Flush()
}
*/

/*
package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	input := "Go is amazing"
	scanner := bufio.NewScanner(strings.NewReader(input))

	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
*/

/*
package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

type UpperReader struct {
	src io.Reader
}

func (u *UpperReader) Read(p []byte) (n int, err error) {
	n, err = u.src.Read(p)

	for i := 0; i < n; i++ {
		p[i] = bytes.ToUpper([]byte{p[i]})[0]
	}
	return n, err
}

func main() {
	original := strings.NewReader("hello go developers")
	upper := &UpperReader{src: original}

	output, _ := io.ReadAll(upper)

	fmt.Println(string(output))
}
*/

/*
package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	source := strings.NewReader("Hello, this is a simple stream test!\n")

	copiedBytes, err := io.Copy(os.Stdout, source)
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("Copied %d bytes.\n", copiedBytes)
	}
}
*/

/*
package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	myName := strings.NewReader("Hello, My name is Amirali\n")

	cpybytes, err := io.Copy(os.Stdout, myName)
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("Copied %d bytes.\n", cpybytes)
	}
}
*/

/*
package main

import "fmt"

type MyReader struct{}

func (m *MyReader) Read(p []byte) (n int, err error) {
	if len(p) == 0 {
		return 0, nil
	}

	p[0] = 'A'
	return 1, nil
}

func main() {
	reader := MyReader{}
	buf := make([]byte, 1, 5)
	n, err := reader.Read(buf)
	fmt.Println("bytes read:", n)
	fmt.Println("content:", string(buf[:n]))
	fmt.Println("error:", err)
}
*/

/*
package main

import "fmt"

type ConsoleWriter struct{}

func (cw *ConsoleWriter) Write(p []byte) (n int, err error) {
	fmt.Printf("%s", string(p))
	return len(p), nil
}

func main() {
	cw := ConsoleWriter{}

	_, err := fmt.Fprintf(&cw, "Hello %s, today is a great day to code!", "Amirali")
	if err != nil {
		fmt.Println("\nError occurred:", err)
	}
}
*/

package main

import "fmt"

type ByteCounter struct {
	Total int
}

func (bc *ByteCounter) Write(p []byte) (n int, err error) {
	holder := len(p)
	bc.Total += holder
	return holder, nil
}

func main() {
	bc := ByteCounter{}

	fmt.Fprint(&bc, "Hello ")
	fmt.Fprint(&bc, "Amirali")
	fmt.Fprint(&bc, ", today is a great day to code!")

	fmt.Println("Total bytes written:", bc.Total)

}
