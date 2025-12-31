/*

package main

import (
	"fmt"
	"io"
)

type String struct {
	data string
	pos  int
}

func (s *String) Read(p []byte) (n int, err error) {
	n = copy(p, s.data[s.pos:])
	s.pos += n

	if s.pos >= len(s.data) {
		err = io.EOF
	}
	return n, err
}

func (s *String) Write(p []byte) (n int, err error) {
	s.data += string(p)

	return len(p), nil
}

func main() {
	var str String

	str.Write([]byte("Hello World!"))
	buffer := make([]byte, 20)
	n, err := str.Read(buffer)
	fmt.Println("Bytes read:", n)
	fmt.Println(err)
	fmt.Println("Buffer:", string(buffer))
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
	reader := strings.NewReader("Teaching go with source book")

	fmt.Print("content copied to stdout:")
	io.Copy(os.Stdout, reader)
	fmt.Println()

	reader2 := strings.NewReader("new data to read")
	data, err := io.ReadAll(reader2)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Data read by ReadAll: %s\n", data)
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
	myName := strings.NewReader("Amirali Nourian")
	io.Copy(os.Stdout, myName)
	myName.Seek(0, io.SeekStart)
	data, err := io.ReadAll(myName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("\nData read:", string(data))
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
	input := "Line 1\nLine 2\nLine 3"
	scanner := bufio.NewScanner(strings.NewReader(input))

	for scanner.Scan() {
		fmt.Println("Scanned Line:", scanner.Text())
	}
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
	GroceryList := `bread
	onion
	burger
	letus
	tomato
	pickle`
	scanner := bufio.NewScanner(strings.NewReader(GroceryList))

	for counter := 1; scanner.Scan(); counter++ {
		fmt.Println("Line", counter, scanner.Text())
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
	reader := strings.NewReader("Golang is awesome!")
	fmt.Println("CopyN output: ")
	io.CopyN(os.Stdout, reader, 5)
	fmt.Println("\n-------------------")

	source := strings.NewReader("Learning TeeReader from the book.")

	tee := io.TeeReader(source, os.Stdout)
	fmt.Print("Reading from tee: ")
	data, _ := io.ReadAll(tee)
	fmt.Printf("\nData stored in variable: %s\n", string(data))
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
	LongText := strings.NewReader("This is a very very very very very very very very long TEXT!")
	fmt.Print("CopyN output: ")
	io.CopyN(os.Stdout, LongText, 10)
	restofText := io.TeeReader(LongText, os.Stdout)
	fmt.Print("\nRest of the TEXT read by TEE: ")
	io.ReadAll(restofText)
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
	reader := strings.NewReader("Learning GO step by step")

	tee := io.TeeReader(reader, os.Stdout)

	io.ReadAll(tee)
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

func main() {
	reader := strings.NewReader("Learning Go is fun!")

	var buf bytes.Buffer

	tee := io.TeeReader(reader, &buf)

	data, err := io.ReadAll(tee)
	if err != nil {
		panic(err)
	}

	fmt.Println("Data read from TeeReader:")
	fmt.Println(string(data))

	fmt.Println("\nData copied into bytes.Buffer:")
	fmt.Println(buf.String())
}
*/

/*
package main

import (
	"bytes"
	"fmt"
)

func main() {
	var myBuffer bytes.Buffer

	myBuffer.WriteString("I am learning ")
	myBuffer.WriteString("bytes.Buffer ")
	myBuffer.WriteString("step by step.")

	content := myBuffer.String()

	fmt.Println("Inside Buffer:", content)
}
*/

/*
package main

import (
	"bytes"
	"fmt"
)

func main() {
	var myNameBuffer bytes.Buffer

	myNameBuffer.WriteString("I am Amirali ")
	myNameBuffer.WriteString("Nourian ")
	myNameBuffer.WriteString("I'm 25 years old.")

	printer := myNameBuffer.String()
	fmt.Println("content is:", printer)

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

func main() {
	reader := strings.NewReader("Go is very powerful!")

	var backup bytes.Buffer

	tee := io.TeeReader(reader, &backup)

	Data, _ := io.ReadAll(tee)

	fmt.Println("Read from TeeReader:", string(Data))

	fmt.Println("Inside our backup buffer:", backup.String())
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

func main() {
	reader := strings.NewReader("Exercises for Chapter 8 of the Reference Book")

	var buff bytes.Buffer

	tee := io.TeeReader(reader, &buff)

	data, _ := io.ReadAll(tee)

	fmt.Println("Read from TeeReader:", string(data))
	fmt.Println("Inside buffer:", buff.String())

	fmt.Println("length of the buffer:", buff.Len())
}
*/

/*
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	pr, pw := io.Pipe()

	go func() {
		defer pw.Close()
		fmt.Fprint(pw, "Data is flowing through the pipe...")
	}()

	fmt.Print("Main thread receiving: ")
	io.Copy(os.Stdout, pr)
	fmt.Println("\nDone.")
}
*/

/*
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	pr, pw := io.Pipe()

	go func() {
		defer pw.Close()
		fmt.Fprint(pw, " Go, Python, C++")
	}()
	fmt.Print("Main Thread receiving:")
	io.Copy(os.Stdout, pr)
	fmt.Println("\nDone")
}
*/

/*
package main

import (
	"fmt"
	"strings"
)

func main() {
	var builder strings.Builder

	builder.WriteString("Learning ")
	builder.WriteString("Go ")
	builder.WriteString("is ")
	builder.WriteString("awesome!")

	result := builder.String()

	fmt.Println(result)
}
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	var builder strings.Builder

	for i := 1; i <= 10; i++ {
		fmt.Fprintf(&builder, "- Number %d ", i)
	}
	result := builder.String()
	fmt.Println(result)

}

// force update commit message
