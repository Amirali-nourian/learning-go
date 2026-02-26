/*package main

import (
	"log"
	"log/slog"
	"net"
	"time"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	err := conn.SetDeadline(time.Now().Add(5 * time.Second))
	if err != nil {
		log.Fatalf("Failed to set deadline: %v", err)
		return
	}

	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			log.Fatalf("Connection closed or read error: %v", err)
			return
		}
		clientData := string(buffer[:n])
		log.Fatalf("Received: %s", clientData)

		_, err = conn.Write([]byte("Server received: " + clientData))
		if err != nil {
			log.Printf("Write error: %v", err)
			return
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	defer listener.Close()

	slog.Info("TCP Server is listening on port 8080...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: %v", err)
			continue
		}
		go handleConnection(conn)
	}
}
*/

/*
package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Println("Error starting server:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server is waiting on port 8080...")
	conn, err := listener.Accept()
	if err != nil {
		log.Println("Error accepting connection:", err)
		return
	}
	defer conn.Close()

	conn.Write([]byte("Hello from Server! Welcome!"))
}
*/
/*
package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	buffer := make([]byte, 1024)

	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Erorr reading:", err)
		return
	}
	fmt.Println("Message from server:", string(buffer[:n]))
}
*/

/*
package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	listener, err := net.Listen("tcp", "localhost:5060")
	if err != nil {
		log.Fatalf("Failed to bind to port: %v", err)
	}
	defer listener.Close()

	log.Println("Server started, ready to receive connections on port 5060...")
	go func() {
		clientCounter := 1
		for {
			conn, err := listener.Accept()
			if err != nil {
				select {
				case <-ctx.Done():
					return
				default:
					log.Printf("Error accepting connection: %v\n", err)
					continue
				}
			}
			go handleConnection(ctx, conn, clientCounter)
			clientCounter++
		}
	}()
	<-ctx.Done()
	log.Println("Shutdown signal received, gracefully stopping server...")
}

func handleConnection(ctx context.Context, conn net.Conn, id int) {
	log.Printf("Client [%d] connected from %s\n", id, conn.RemoteAddr().String())

	defer conn.Close()

	buffer := make([]byte, 1024)

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			if err == io.EOF {
				log.Printf("Client [%d] closed the connection.\n", id)
			} else {
				log.Printf("Client [%d] read error: %v\n", id, err)
			}
			return
		}
		fmt.Printf("Message from Client %d: %s\n", id, string(buffer[:n]))
	}
}
*/

/*
package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		panic("Can not create a listener, error:" + err.Error())
	}
	defer listener.Close()

	fmt.Println("Server started, ready to receive connction")

	var i = 1
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic("error in accepting connection, error: " + err.Error())
		}
		go handleConnection(conn, i)
		i++
	}
}

func handleConnection(conn net.Conn, i int) {
	fmt.Println("A new connection established, number assigned:", i)

	defer conn.Close()

	buffer := make([]byte, 1024)

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("error in reading the connection, error: " + err.Error())
			return
		}
		fmt.Printf("Client %d: %s\n", i, string(buffer[:n]))
	}
}
*/

/*
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	conn, err := net.DialTimeout("tcp", "localhost:8080", 3*time.Second)
	if err != nil {
		log.Fatalf("Failed to connect to server: %v\n", err)
	}
	defer conn.Close()

	log.Println("Connected to TCP server on localhost:8080")
	fmt.Println("Write something to send to the server (type 'exit' to quit):\n")

	go func() {
		serverScanner := bufio.NewScanner(conn)
		for serverScanner.Scan() {
			fmt.Println("Server:", serverScanner.Text())
			fmt.Print(">> ")
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)

	for {

		message := strings.TrimSpace(scanner.Text())

		if message == "exit" {
			log.Println("Closing connection gracefully...")
			break
		}

		if message == "" {
			continue
		}

		_, err = fmt.Fprintln(conn, message)
		if err != nil {
			log.Printf("Error sending data to server: %v\n", err)
			break
		}
	}

}
*/

/*
package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func main() {
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Server Started...")

	http.ListenAndServe("localhost:8080", nil)
}
*/

/*
package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/bye", http.StatusTemporaryRedirect)
	w.Write([]byte("you are now on /bye"))
}

func main() {
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Server started...")

	http.ListenAndServe(":8080", nil)

}
*/

package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w,
		"Host:\t", r.Host, "\n",
		"URI:\t", r.RequestURI, "\n",
		"Query:\t", r.URL.RawQuery, "\n",
		"Path:\t", r.URL.Path, "\n",
		"Name:\t", r.URL.Query().Get("name"), "\n",
		"Age:\t", r.URL.Query().Get("age"), "\n",
	)
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.ListenAndServe(":8080", nil)
}
