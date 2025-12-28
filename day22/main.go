/*package main

import "fmt"

type String string

func (s *String) Write(p []byte) (int, error) {
	*s += String(string(p))
	return len(p), nil
}

func main() {
	var str String
	str.Write([]byte("Hello World!"))
	fmt.Println(str)
}
*/

/*
package main

import (
	"fmt"
	"io"
)

type String string

func (s *String) Write(p []byte) (int, error) {
	*s += String(string(p))
	return len(p), nil
}

func SayHello(w io.Writer) {
	w.Write([]byte("Hello World"))
}

func main() {
	var str String
	SayHello(&str)

SayHello(os.Stdout)
}
*/

/*
package main

import (

	"io"
	"os"

)

	func SayHello(w io.Writer) {
		w.Write([]byte("Hello World"))
	}

	func main() {
		SayHello(os.Stdout)
	}
*/

/*
package main

import (

	"io"
	"os"

)

	func SayHello(w io.Writer) {
		w.Write([]byte("Hello World"))
	}

	func main() {
		file, _ := os.Create("hello.txt")
		defer file.Close()
		SayHello(file)
	}
*/

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

func (s *String) Read(p []byte) (int, error) {
	n := copy(p, s.data[s.pos:])
	s.pos += n

	var err error
	if s.pos >= len(s.data) {
		err = io.EOF
	}

	return n, err
}

func (s *String) Write(p []byte) (int, error) {
	s.data += string(p)
	return len(p), nil
}

func main() {
	var str String

	str.Write([]byte("Hello World"))
	buffer := make([]byte, 20)

	n, err := str.Read(buffer)

	fmt.Println("Bytes read:", n)
	fmt.Println("Error:", err)
	fmt.Println("Buffer:", string(buffer[:n]))
}
*/

/*
package main

import (
	"fmt"
	"os"
)

func main() {
	buffer := make([]byte, 20)
	n, _ := os.Stdin.Read(buffer)

	fmt.Println("buffer:", string(buffer[:n]))
}
*/

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

func (s *String) Read(p []byte) (int, error) {
	if s.pos >= len(s.data) {
		return 0, io.EOF
	}

	n := copy(p, s.data[s.pos:])
	s.pos += n

	if s.pos >= len(s.data) {
		return n, io.EOF

	}
	return n, nil
}

func (s *String) Write(p []byte) (int, error) {
	s.data += string(p)
	return len(p), nil
}

func ReadWriteSomething(rw io.ReadWriter) string {
	_, _ = rw.Write([]byte("Go language is fun!"))

	if s, ok := rw.(*String); ok {
		s.pos = 0
	}

	buf := make([]byte, 64)
	n, _ := rw.Read(buf)

	return string(buf[:n])
}

func main() {
	var str String
	result := ReadWriteSomething(&str)
	fmt.Println(result)

}
*/
/*
package main

import (
	"bufio"
	"fmt"
	"io"
)

type String struct {
	data string
	pos  int
}

func (s *String) Read(p []byte) (int, error) {
	n := copy(p, s.data[s.pos:])
	s.pos += n

	var err error
	if s.pos >= len(s.data) {
		err = io.EOF
	}
	fmt.Println("String.Read called:", n, "bytes read")
	return n, err
}

func (s *String) Write(p []byte) (int, error) {
	s.data += string(p)
	fmt.Println("String.Write called:", string(p))
	return len(p), nil
}

func main() {
	rw := &String{}
	//r := bufio.NewReader(rw)
	w := bufio.NewWriter(rw)

	io.WriteString(w, "Hello ")
	io.WriteString(w, "this ")
	io.WriteString(w, "is a sample ")
	io.WriteString(w, "text")
	w.Flush()
	fmt.Println()

	buffer := make([]byte, 3)
	var err error
	var n int
	for err != io.EOF {
		n, err = rw.Read(buffer)
		fmt.Println(string(buffer[:n]))
	}
}
*/

package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	data := []byte{65, 66, 67}
	reader := bytes.NewReader(data)

	var buff bytes.Buffer

	io.Copy(&buff, reader)

	io.WriteString(&buff, "\nEnd of Chapter 8")
	fmt.Println(buff.String())
}
