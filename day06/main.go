package main

import (
	"fmt"
	_ "hello/goodbye-package"
	"hello/hello-package"

	"github.com/mostafasolati/goodbye"
)

var language string

// init function will be the function that will be used before main function | init قبل از فانکشن main اجرا می‌شود.
// init has a higher position than main کامپایلر این فانکشن را به‌صورت خودکار اجرا می‌کند.
func init() {
	language = "Go!"
}

func main() {
	fmt.Println("I love ", language)
	//importing internal package. To do so I HAVE to add go.mod with go mod init "اسم دلخواه" to my project that I want to use the package
	//Here I must add go mod init hello to use the hello package
	//Note that to use hello package I must obey Exporting rules (caoital first letters to use it)
	hello.SayHello("Amiral")
	//goodbye.SayGoodbye("Amirali")
	//using open-source projects
	goodbye.SayGoodbye("Amirali")
}


