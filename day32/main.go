/*
package main

import (
	"fmt"
	"net/http"
	"sync"
)

type CounterHandler struct {
	mu    sync.Mutex
	count int
}

func (c *CounterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c.mu.Lock()
	c.count++
	currentCount := c.count
	c.mu.Unlock()

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, "Total Requests Processed: %d\n", currentCount)
}

func main() {
	counter := &CounterHandler{}

	http.Handle("/count", counter)

	fmt.Println("Server starting on :8080/count...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
*/

/*
package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type FileMetadataHandler struct {
	ServerName string
	StartTime  time.Time
	once       sync.Once
}

func (f *FileMetadataHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f.once.Do(func() {
		fmt.Println("✅ First request received!")
		fmt.Println("Server started at:", f.StartTime.Format("2006-01-02 15:04:05"))

	})
	uptime := time.Since(f.StartTime)

	fmt.Fprintf(w, "Server Name: %s\n", f.ServerName)
	fmt.Fprintf(w, "Uptime: %s\n", uptime)
	fmt.Fprintf(w, "Requested URL Path: %s\n", r.URL.Path)
}

func main() {
	handler := &FileMetadataHandler{
		ServerName: "Static File Server Simulator",
		StartTime:  time.Now(),
	}
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", handler)
}
*/

package main

import (
	"log"
	"net/smtp"
)

func main() {
	from := "system@cableon.ir"
	password := "secure_Password"
	to := []string{"admin@example.com"}
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	subject := "Subject: Server Critical Alert\n"
	body := "The backend server has reached 90% CPU usage."
	msg := []byte(subject + "\n" + body)

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, msg)
	if err != nil {
		log.Fatalf("Failed to send email: %s", err)
	}

	log.Println("Email sent successfullu.")
}
