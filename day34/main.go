/*package main

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
*/

/*
package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("assets"))

	http.Handle("/static/", http.StripPrefix("/static", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	log.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}
*/

/*
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	tmpl := template.Must(template.ParseFiles("login.html"))

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Error in form processing", http.StatusBadRequest)
			return
		}

		user := r.FormValue("username")
		pass := r.FormValue("password")

		if user == "admin" && pass == "1234" {
			fmt.Fprintf(w, "<h1>خوش آمدید، مدیر سیستم! ✅</h1>")
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "<h1>نام کاربری یا رمز عبور اشتباه است ⛔</h1>")
		}
	})
	log.Println("Server running on :8080/login")
	http.ListenAndServe(":8080", nil)
}
*/

/*
package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

func contactHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmpl, err := template.ParseFiles("contact.html")
		if err != nil {
			http.Error(w, "Error loading form!", http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, nil)
		return
	}
	if r.Method == http.MethodPost {
		email := r.FormValue("email")
		subject := r.FormValue("subject")
		message := r.FormValue("message")

		if !strings.Contains(email, "@") {
			fmt.Fprintf(w, "invalid Email")
			return
		}
		fmt.Fprintf(w, "your message with subject [%s] has been received!", subject)

		_ = message
		return
	}

	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

func main() {
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))

	http.HandleFunc("/contact", contactHandler)

	fmt.Println("Server running on http://localhost:8080/contact")
	http.ListenAndServe(":8080", nil)
}
*/

/*
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

// ۱. متغیر سراسری برای قالب (لود شدن فقط یک‌بار در رم)
var contactTmpl *template.Template

func init() {
	// استفاده از Must باعث می‌شود اگر فایل پیدا نشد، برنامه سریعاً Panic کند
	// بهتر است در شروع برنامه بفهمیم فایل نیست تا اینکه در زمان اجرا به ارور بخوریم
	contactTmpl = template.Must(template.ParseFiles("contact.html"))
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// استفاده از قالب لود شده در رم (بسیار سریع)
		contactTmpl.Execute(w, nil)
		return
	}

	if r.Method == http.MethodPost {
		// محدود کردن حجم ورودی برای امنیت حافظه
		r.Body = http.MaxBytesReader(w, r.Body, 1048576) // حداکثر ۱ مگابایت

		email := r.FormValue("email")
		subject := r.FormValue("subject")
		message := r.FormValue("message")

		if !strings.Contains(email, "@") {
			// ارسال کد ۴۰۰ (Bad Request) به جای متن ساده
			http.Error(w, "Email format is invalid", http.StatusBadRequest)
			return
		}

		// خروجی تمیز
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, "<h3>Success!</h3><p>Message about <b>%s</b> received.</p>", subject)

		log.Printf("New Contact: %s - Content Length: %d", email, len(message))
		return
	}

	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

func main() {
	// هندل کردن فایل‌های استاتیک
	fs := http.FileServer(http.Dir("./assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	http.HandleFunc("/contact", contactHandler)

	log.Println("Server running on http://localhost:8080/contact")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
*/

/*
package main

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
	"time"
)

type LinkCheckRequest struct {
	URL string `json: "url"`
}

func checkLinkHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req LinkCheckRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	slog.Info("Starting SEO status check", "target_url", req.URL)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Link check initialized successfully"}`))
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	mux := http.NewServeMux()

	mux.HandleFunc("/api/v1/seo/check", checkLinkHandler)

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	slog.Info("SEO Backend API is starting", "port", 8080)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		slog.Error("Server encountered a fatal error", "error", err)
	}
}
*/

/*
package main

import (
	"log/slog"
	"net/http"
	"os"
)

// APIRouter ساختار روتر اختصاصی ما
type APIRouter struct {
	logger *slog.Logger
}

// ServeHTTP با این متد، استراکت ما رسماً به یک روتر تبدیل می‌شود
func (router *APIRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// ۱. کارهایی که باید برای همه درخواست‌ها انجام شود (متمرکز)
	w.Header().Set("Content-Type", "application/json")
	router.logger.Info("Request received", "path", r.URL.Path)

	// ۲. مسیریابی دستی با استفاده از switch
	switch r.URL.Path {
	case "/users":
		router.handleUsers(w, r)
	case "/status":
		router.handleStatus(w, r)
	default:
		// ۳. هندلینگ متمرکز برای مسیرهای نامعتبر (خطای 404)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error": "path not found"}`))
	}
}

// handleUsers هندلر مدیریت کاربران
func (router *APIRouter) handleUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(`{"error": "method not allowed"}`))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"users": ["ali", "sara"]}`))
}

// handleStatus هندلر بررسی وضعیت سیستم
func (router *APIRouter) handleStatus(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "running"}`))
}

func main() {
	jsonLogger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// نمونه‌سازی از روتر دست‌ساز خودمان
	myRouter := &APIRouter{
		logger: jsonLogger,
	}

	srv := &http.Server{
		Addr:    ":8080",
		Handler: myRouter, // روتر اختصاصی را مستقیماً به سرور می‌دهیم
	}

	jsonLogger.Info("Server is running on port 8080")
	srv.ListenAndServe()
}
*/

/*
package main

import (
	"fmt"
	"net/http"
)

type mymux struct{}

func (m *mymux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/a":
		fmt.Fprint(w, "This is A route")
	case "/b":
		fmt.Fprint(w, "This is B route")
	case "/c":
		fmt.Fprint(w, "This is C route")
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "My Custom 404 Page")
	}
}

func main() {
	mux := new(mymux)
	http.ListenAndServe("localhost:5050", mux)
}
*/

/*
package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	publicMux := http.NewServeMux()
	publicMux.HandleFunc("/api/users", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"cpu": "12%", "ram": "45MB"}`))
	})

	adminMux := http.NewServeMux()
	adminMux.HandleFunc("/admin/metrics", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"cpu": "12%", "ram": "45MB"}`))
	})

	publicServer := &http.Server{
		Addr:         ":8080",
		Handler:      publicMux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	adminServer := &http.Server{
		Addr:         ":9090",
		Handler:      adminMux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		log.Println("Public Server is starting on port 8080...")
		if err := publicServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Public server failed: %v", err)
		}
	}()

	go func() {
		log.Println("Admin Server is starting on port 9090...")
		if err := adminServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Admin server failed: %v", err)
		}
	}()
	select {}
}
*/

package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	webMux := http.NewServeMux()
	webMux.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to homepage"))
	})

	uploadMux := http.NewServeMux()
	uploadMux.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("File upload system"))
	})

	webServer := &http.Server{
		Addr:         ":8080",
		Handler:      webMux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	uploadServer := &http.Server{
		Addr:         ":8001",
		Handler:      uploadMux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		log.Println("Web Server is starting on port :8080...")
		if err := webServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Web Sever failed: %v", err)
		}
	}()

	go func() {
		log.Println("Upload Server is starting on port :8001...")
		if err := uploadServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Upload Server failed: %v", err)
		}
	}()

	select {}
}
