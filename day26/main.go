/*package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	sourceFile, err := os.Open("Source.txt")
	if err != nil {
		fmt.Println("Error opening source:", err)
		return
	}
	defer sourceFile.Close()

	destFile, err := os.Create("Destination.txt")
	if err != nil {
		fmt.Println("Error creating destination:", err)
		return
	}
	defer destFile.Close()

	bytesCopied, err := io.Copy(destFile, sourceFile)
	if err != nil {
		fmt.Println("Error copying file:", err)
		return
	}

	fmt.Printf("Copied %d bytes successfully.\n", bytesCopied)
}
*/

/*
package main

import (
	_ "embed"
	"fmt"
)

var Version string

var logoBytes []byte

func main() {
	fmt.Println("Application Version:", Version)
	fmt.Printf("Logo size in bytes: %d\n", len(logoBytes))
}
*/

/*
package main

import (
	_ "embed"
	"fmt"
	"io"
	"os"
)

//go:embed config.txt
var configContent string

func main() {

	note, err := os.Create("note.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer note.Close()

	_, err = note.WriteString("This is a text to be written in the note.txt!!")
	if err != nil {
		fmt.Println("Error writing to note.txt:", err)
		return
	}

	_, err = note.Seek(0, 0)
	if err != nil {
		panic(err)
	}

	backup_note, err := os.Create("backup_note.txt")
	if err != nil {
		fmt.Println("Error creating backup_note.txt file:", err)
		return
	}
	defer backup_note.Close()

	_, err = io.Copy(backup_note, note)
	if err != nil {
		fmt.Println("Error copying contents:", err)
		return
	}

	fmt.Println("Embedded config content:")
	fmt.Println(configContent)
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

		fmt.Println("Error creating file:", err)
		return
	}

	defer file.Close()

	file.Write([]byte("Hello World!\n"))
	file.WriteString("A new line\n")
	io.WriteString(file, "Third line!\n")
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
	err := os.WriteFile("newfile.txt", []byte("This is The End!"), 0777)
	if err != nil {
		fmt.Println("Error creating file:", err)
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
	file, err := os.Open("myfile.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	buff, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(buff))
}
