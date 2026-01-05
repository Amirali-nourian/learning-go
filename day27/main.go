/*package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.OpenFile("myfile.txt", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	defer file.Close()

	file.WriteString("Hello")

	file.Seek(2, io.SeekStart)
	buff := make([]byte, 3)
	file.Read(buff)
	fmt.Println(string(buff))
}
*/

/*
package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("myfile.txt", os.O_RDWR, 0777)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	defer file.Close()

	file.WriteAt([]byte("LL"), 2)
	buff := make([]byte, 2)
	file.ReadAt(buff, 2)
	fmt.Println(string(buff))
}
*/
/*
package main

import (
	"fmt"
	"os"
)


type info interface {
	Name() string
	Size() int
	Mode() FileMode
	ModTime() time.Time
	IsDir() bool
	Sys() any
}*/

/*

func main() {
	info, err := os.Stat("myfile.txt")
	if err != nil {
		fmt.Println("Error on stateing:", err)
		return
	}
	fmt.Println("File name:", info.Name())
	fmt.Println("Size in bytes:", info.Size())
	fmt.Println("Permissions:", info.Mode())
	fmt.Println("Last Modified:", info.ModTime())
	fmt.Println("Is Directory:", info.IsDir())
	fmt.Printf("System interface type: %T\n", info.Sys())
	fmt.Printf("System info: %+v\n\n", info.Sys())
}
*/

/*
package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Stat("invalidfilename.txt")
	if os.IsNotExist(err) {
		fmt.Println("file doesn't exist!")
	}
}
*/

/*
package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.Remove("myfile.txt")
	if err != nil {
		fmt.Println("Error removing the file:", err)
		return
	}
	fmt.Println("File removed successfully!")
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
	src, err := os.Open("file1.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	dst, err := os.Create("file2.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}

	io.Copy(dst, src)
}
*/

/*
package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.Mkdir("images", 0777)
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}
}
*/

/*
package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.MkdirAll("a/b/c/d", 0777)
	if err != nil {
		fmt.Println("Error creating directories:", err)
		return
	}
	fmt.Println("Directories created successfully!")
}
*/

/*
package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.ReadDir(".")
	if err != nil {
		fmt.Println("Error reading directories:", err)
		return
	}
	for _, val := range file {
		if val.IsDir() {
			fmt.Println("[Folder]", val.Name())
		} else {
			fmt.Println("[File]", val.Name())
		}
	}
}
*/

package main

import (
	"embed"
	_ "embed"
	"fmt"
)

//go:embed file1.txt
var file1 string

//go:embed file2.txt
var file2 string

//go:embed images
var folder embed.FS

func main() {
	fmt.Println(file1)
	fmt.Println(string(file2))
	file, err := folder.ReadFile("images/file3.txt")
	if err != nil {
		fmt.Println("Error finding file:", err)
		return
	}
	fmt.Println(string(file))

	files, err := folder.ReadDir("images")
	if err != nil {
		fmt.Println("Error reading direction:", err)
		return
	}
	fmt.Println("contents of images folder:")
	for _, f := range files {
		fmt.Println(f.Name())
	}
}
