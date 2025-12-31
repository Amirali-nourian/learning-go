/*package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
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
	rw := &String{data: "this is a sample text with some words"}

	scanner := bufio.NewScanner(rw)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	scanner = bufio.NewScanner(os.Stdin)
	fmt.Print("Enter your full name: ")
	scanner.Scan()
	fmt.Println("Hello", scanner.Text())
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
	reader := strings.NewReader("This is a sample text")
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Println(scanner.Text())

	}
}
*/

/*
package main

import (
	"fmt"
	"strings"
)

func main() {
	writer := new(strings.Builder)
	writer.WriteString("Hello World\n")
	fmt.Fprint(writer, "This is second line")
	fmt.Println(writer.String())
}
*/

/*
package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	buf := new(bytes.Buffer)
	buf.WriteString("This is the first line\n")
	buf.Write([]byte("This is second line"))
	fmt.Fprint(buf, "This is third line\n")
	io.WriteString(buf, "This is fourth line")
	bs, _ := io.ReadAll(buf)
	fmt.Println("Waht", string(bs))

}
*/

package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	buff1 := bytes.NewBuffer([]byte("buffer 1"))
	buff2 := bytes.NewBuffer([]byte("buffer 2"))
	fmt.Println(buff1.String())
	fmt.Println(buff2.String())

	reader := bytes.NewReader([]byte("Hello World"))
	io.Copy(os.Stdout, reader)
}



// force update commit message
