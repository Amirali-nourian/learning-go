/*package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("test.txt")

	if err != nil {
		fmt.Println("error creating file:", err)
		return
	}
	defer file.Close()

	fmt.Println("File created successfully!")
}
*/

/*package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("note.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	content := "Hello Go programming!"
	l, err := file.WriteString(content)

	if err != nil {
		fmt.Println("Error writing to file:", err)
		file.Close()
		return
	}

	fmt.Printf("%d bytes written successfully\n", l)
}
*/

/*package main

import (
	"fmt"
	"os"
)

func main() {
	var FileName string

	fmt.Println("Enter a name for file name:")
	fmt.Scanln(&FileName)

	file, err := os.Create(FileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()
	fmt.Printf("%s created successfully.", FileName)
}
*/

/*
package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("note.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	content := "Go is a powerful language for system programming."
	bytesWritten, err := file.WriteString(content)
	if err != nil {
		panic(err)
		return
	}
	fmt.Printf("Successfully wrote %d bytes to the file.\n", bytesWritten)

}
*/

/*
package main

import (
	"fmt"
	"os"
)

func main() {
	var name string
	var age uint8

	fmt.Println("Enter your name: ")
	fmt.Scanln(&name)

	fmt.Println("Enter your age: ")
	fmt.Scanln(&age)

	file, err := os.Create("user.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	content := fmt.Sprintf("Name: %s, Age: %d", name, age)

	bytesWritten, err := file.WriteString(content)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Printf("Successfully wrote %d bytes\n", bytesWritten)
	fmt.Println("Information saved successfully")

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
	file, err := os.Create("myfile.txt")

	if err != nil {
		fmt.Println("Error while creating file:", err)
		return
	}
	defer file.Close()

	file.Write([]byte("Hello World!\n"))
	file.WriteString("A new line\n")
	io.WriteString(file, "Third line\n")
	fmt.Fprint(file, "Fourth line", 1234)
}
*/

/*
package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("user.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	data := make([]byte, 100)

	count, err := file.Read(data)
	if err != nil {
		fmt.Println("Error rading file:", err)
		return
	}

	fmt.Printf("Read %d bytes: %s\n", count, string(data[:count]))
}
*/

/*
package main

import (
	"fmt"
	"os"
)

func main() {
	var fileName string

	fmt.Println("What file you want to open with?")

	fmt.Scanln(&fileName)

	file, err := os.Open(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("No such file found!", err)
		} else {
			fmt.Println("Error opening file:", err)
		}
		return
	}

	defer file.Close()

	//data := make([]byte, 100)

	count, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("Content:", string(count))
}
*/

/*
package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("File created successfully!")
	defer file.Close()

	_, err = file.WriteString("New log entry\n")
	if err != nil {
		fmt.Println("Error writing:", err)
	}
}
*/

/*
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("notes.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	fmt.Println("Enter a sentence:")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()

	_, err = file.WriteString(text + "\n")
	if err != nil {
		fmt.Println("Error writing:", err)
		return
	}

	fmt.Println("Saved successfully!")
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
	file, err := os.OpenFile("secret.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	file.WriteString("Hi")

	file.Seek(0, 0)

	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("Read content:", string(data))

	fileInfo, err := file.Stat()
	if err == nil {
		fmt.Println("File Name:", fileInfo.Name())
		fmt.Println("Size in bytes:", fileInfo.Size())
	}
}
*/

package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.OpenFile("data.txt", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	file.WriteString("GoProgramming")

	file.Seek(0, 0)

	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println("Content:", string(data))

	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("Error on getting status:", err)
		return
	} else {
		fmt.Println("File name:", fileInfo.Name())
		fmt.Println("Size in Bytes:", fileInfo.Size())
	}

	file.Seek(2, io.SeekStart)

	data, err = io.ReadAll(file)
	if err != nil {
		panic(err)
		return
	}
	fmt.Println("read from second byte:", string(data))

	file.Seek(-3, io.SeekEnd)

	data, err = io.ReadAll(file)
	if err != nil {
		panic(err)
		return
	}
	fmt.Println("read from last 3 bytes:", string(data))

}
