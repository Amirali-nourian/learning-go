/*package main

import (
	"fmt"
	"log"
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

	clientIP := r.RemoteAddr
	requestPath := r.URL.Path

	go f.sendAlertEmail(
		fmt.Sprintf("Client IP: %s | Path: %s", clientIP, requestPath),
	)

	fmt.Fprintf(w, "Server Name: %s\n", f.ServerName)
	fmt.Fprintf(w, "Uptime: %s\n", uptime)
	fmt.Fprintf(w, "Requested URL Path: %s\n", r.URL.Path)
}

func (f *FileMetadataHandler) sendAlertEmail(details string) {
	log.Printf("Sending alert email with details: %s\n", details)

	time.Sleep(2 * time.Second)

	log.Printf("Alert email sent successfully with details: %s\n", details)
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

/*
package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Customer is here...")
		go func() {
			time.Sleep(5 * time.Second)
			fmt.Println("Heavy work is done in background!")
		}()

		fmt.Fprintln(w, "see how fast I handled it??")
	})

	fmt.Println("Server on :8080")
	http.ListenAndServe(":8080", nil)
}
*/

/*
package main

import (
	"log"
	"net/http"
	"text/template"
)

type PageData struct {
	Title    string
	User     string
	IsActive bool
}

func main() {
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := PageData{
			Title:    "User Panel",
			User:     "Amirali",
			IsActive: true,
		}

		err := tmpl.Execute(w, data)
		if err != nil {
			log.Printf("Error executing template: %v", err)
			http.Error(w, "Internal Server Error", 500)
		}
	})
	log.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}
*/

package main

import (
	"html/template"
	"log"
	"net/http"
)

type Student struct {
	Name   string
	Scores []int
}

func main() {
	tmpl, err := template.ParseFiles("report.html")
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		test := Student{
			Name:   "Amirali",
			Scores: []int{20, 15, 9, 18, 5, 12},
		}

		err := tmpl.Execute(w, test)
		if err != nil {
			log.Printf("Error executing template: %v", err)
			http.Error(w, "Internal Server Error", 500)
		}

	})
	log.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}
