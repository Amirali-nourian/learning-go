/*
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

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("authMiddleware")

		if apiKey != "secret123" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)

			json.NewEncoder(w).Encode(map[string]string{
				"error": "Unauthorized access",
			})
			return
		}
		next.ServeHTTP(w, r)
	})
}

func targetHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write([]byte(`{"message":"You have access!"}`))
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	mux := http.NewServeMux()

	protectedHandler := http.HandlerFunc(targetHandler)

	mux.Handle("/api/v1/protected", authMiddleware(protectedHandler))

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	slog.Info("Server running", "port", 8080)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		slog.Error("Server fatal error", "error", err)
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

func addVersionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-App-Version", "1.0.0")
		next.ServeHTTP(w, r)
	})
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello from the main handler!"))
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	mux := http.NewServeMux()

	coreHandler := http.HandlerFunc(mainHandler)

	wrappedHandler := addVersionMiddleware(coreHandler)

	mux.Handle("/hello", wrappedHandler)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	logger.Info("Server is running on port 8080 with Middleware")
	srv.ListenAndServe()
}
*/

package main

import (
	"log/slog"
	"net/http"
	"os"
)

func loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-App-Version", "2.0.0")
		slog.Info("Incoming Request")
		if r.Method != http.MethodGet {
			http.Error(w, "Method Not Allowed - Only GET is accepted", http.StatusMethodNotAllowed)
			return
		}
		next.ServeHTTP(w, r)

	})
}

func homehandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from main handler"))
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	mux := http.NewServeMux()

	mainHandler := http.HandlerFunc(homehandler)

	mux.Handle("/", loggerMiddleware(mainHandler))

	slog.Info("Server is running on :8080")
	http.ListenAndServe(":8080", mux)

}
