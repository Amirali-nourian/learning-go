/*package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("Error starting server:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server is waiting on port 8080...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting connection:", err)
			continue
		}

		go handleConnection(conn)

	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	conn.Write([]byte("Hello from Server! Welcome!\n"))

	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		msg := scanner.Text()
		fmt.Printf("Received: %s\n", msg)

		conn.Write([]byte("Message received: " + msg + "\n"))
	}
	if err := scanner.Err(); err != nil {
		log.Println("Connection error:", err)
	}

}
*/

/*
package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("Error starting server: ", err)
		return
	}

	defer listener.Close()

	fmt.Println("Server running on http://localhost:8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Error stablsihing connection: ", err)
			return
		}

		now := time.Now()

		body := "<html><body>"
		body += "<h1> ساعت رسمی </h1>"
		body += fmt.Sprintf("<b> امروز </b> %d/%d/%d", now.Year(), now.Month(), now.Day())
		body += "</br>"
		body += fmt.Sprintf("<b> ساعت </b> %d:%d", now.Hour(), now.Minute())
		body += "</html></body>"

		response := "HTTP/1.1 200 OK\r\n"
		response += "Content-Type: text/html; charset=UTF-8\r\n"
		response += "Connection: close\r\n"
		response += fmt.Sprintf("Content-Length: %d\r\n", len(body))
		response += "\r\n"
		response += body

		fmt.Fprint(conn, response)
		conn.Close()
	}
}
*/

package main

func main() {

}
