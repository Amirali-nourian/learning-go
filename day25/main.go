/*package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.OpenFile("myfile.txt", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("error opening file:", err)
		return
	}
	//defer file.Close()

	data, err := file.WriteString("This is a test to see if it works!!")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	fmt.Printf("file contains %d bytes in file\n", data)

	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		fmt.Println("Error seeking file:", err)
		return
	}
	content, _ := io.ReadAll(file)
	fmt.Println("Content:", string(content))
	file.Close()

	err = os.Rename("myfile.txt", "newname.txt")
	if err != nil {
		fmt.Println("Error while renaming file:", err)
		return
	}

	err = os.Remove("newname.txt")
	if err != nil {
		fmt.Println("Error while removing file:", err)
		return
	}

}
*/

/*
package main

import (
	"io"
	"os"
)

func main() {
	file, err := os.Create("file1.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	file2, err := os.Create("file2.txt")

	if err != nil {
		panic(err)
	}
	defer file2.Close()

	src, err := os.Open("file1.txt")
	if err != nil {
		panic(err)
	}
	dst, err := os.Open("file2.txt")
	if err != nil {
		panic(err)
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
		fmt.Println("Error creating images directory:", err)
		return
	}
	err = os.MkdirAll("a/b/c/d", 0777)
	if err != nil {
		fmt.Println("Error creating nested directories:", err)
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

	err := os.Mkdir("Backup", 0755)
	if err != nil {
		fmt.Println("Error creating Backup directory:", err)
		return
	}

	if _, err := os.Stat("data.txt"); err != nil {
		fmt.Println("data.txt does not exist:", err)
		return
	}
	err = os.Rename("data.txt", "Backup/old_data.txt")
	if err != nil {
		fmt.Println("Error moving file:", err)
		return
	}

	if _, err := os.Stat("Backup/old_data.txt"); err == nil {
		fmt.Println("Backup/old_data.txt")
	} else {
		fmt.Println("Move failed:", err)
		return
	}

	temp, err := os.Create("temp.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}

	temp.Close()

	err = os.Remove("temp.txt")
	if err != nil {
		fmt.Println("Error removing temp file:", err)
		return
	}
	fmt.Println("Temp file created and removed successfully.")

}
*/

/*
package main

import (
	"fmt"
	"os"
)

func main() {
	files, err := os.ReadDir(".")
	if err != nil {
		fmt.Println("Can't read directory:", err)
		return
	}
	for _, file := range files {
		if file.IsDir() {
			fmt.Print("[folder]")
		} else {
			fmt.Print("[file]")
		}
		fmt.Print(file.Name())
		fmt.Println()
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
	err := os.MkdirAll("projects/go/src", 0755)
	if err != nil {
		fmt.Println(err)
	}

	enteries, err := os.ReadDir(".")
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	for _, entry := range enteries {
		if entry.IsDir() {
			fmt.Print("[Folder]")
		} else {
			fmt.Print("[File] ")
		}
		fmt.Println(entry.Name())
	}
}
*/

package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.MkdirAll("App/Data/Logs", 0755)
	if err != nil {
		fmt.Println("Error creating folders:", err)
		return
	}

	logFilePath := "App/Data/Logs/app.log"
	file, err := os.Create(logFilePath)
	if err != nil {
		fmt.Println("Error creating log file:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString("Application started successfully\n")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	etry, err := os.ReadDir("App/Data")
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	for _, entry := range etry {
		if entry.IsDir() && entry.Name() == "Logs" {
			fmt.Println("Directory:", entry.Name())
		}
	}
	info, err := os.Stat(logFilePath)
	if err != nil {
		fmt.Println("Error getting file info:", err)
		return
	}
	fmt.Println("app.log size:", info.Size(), "bytes")
}
