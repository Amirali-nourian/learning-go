/*package main

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
		Addr:         ":8000",
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
		log.Println("Web Server is starting on port :8000...")
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
*/

/*
package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"
)

func profileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", 405)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"name": "Amirali", "role": "Senior Developer"}`))
}

func metricsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", 405)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"active_goroutines": 12, "memory_usage": "15MB"}`))

}

func main() {
	profilemux := http.NewServeMux()
	profilemux.HandleFunc("/profile", profileHandler)

	monitorMux := http.NewServeMux()
	monitorMux.HandleFunc("/metrics", metricsHandler)

	profileServer := &http.Server{
		Addr:         ":8080",
		Handler:      profilemux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	monitorServer := &http.Server{
		Addr:         ":9090",
		Handler:      monitorMux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	go func() {
		logger.Info("Profile server is starting on port :8080")
		if err := profileServer.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("Profile Server  failed: %v", err)
		}
	}()

	go func() {
		logger.Info("Monitor server is starting on port :9090")
		if err := monitorServer.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("Monitor Server failed: %v", err)
		}
	}()

	select {}
}
*/

/*
package main

import (
	"fmt"
	"net/http"
)

type Middleware func(http.Handler) http.Handler

func Before(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("run berfore", r.URL.Path)
		handler.ServeHTTP(w, r)
	})
}

func After(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)
		fmt.Println(r.URL.Path, "run after")
	})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})

	middlewares := []Middleware{
		Before,
		After,
	}

	var finalHandler http.Handler = mux
	for _, middleware := range middlewares {
		finalHandler = middleware(finalHandler)
	}

	http.ListenAndServe("localhost:5050", finalHandler)
}
*/

package main

import (
	"log/slog"
	"net/http"
	"os"
	"time"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		slog.Info("Request started", "method", r.Method, "path", r.URL.Path)

		next.ServeHTTP(w, r)

		duration := time.Since(start)
		slog.Info("Request finished", "duration_ms", duration.Milliseconds())
	})
}

func targetHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(100 * time.Millisecond)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Hello from the core handler"}`))
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	mux := http.NewServeMux()

	finalHandler := http.HandlerFunc(targetHandler)

	wrappedHandler := loggingMiddleware(finalHandler)

	mux.Handle("/api/hello", wrappedHandler)

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	slog.Info("Server is listening with Middleware...", "port", 8080)
	srv.ListenAndServe()
}
