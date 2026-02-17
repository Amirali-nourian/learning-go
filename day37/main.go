/*package main

import (
    "log/slog"
    "net/http"
    "os"
    "time"
)

func main() {
    logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

    client := &http.Client{
        Timeout: 20 * time.Second,
    }

    url := "https://jsonplaceholder.typicode.com/users/1"
    logger.Info("Sending GET request", "url", url)

    resp, err := client.Get(url)
    if err != nil {
        logger.Error("Failed to make HTTP request", "error", err)
        return
    }

    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        logger.Error("Received bad status code from server", "status", resp.StatusCode)
        return
    }

    logger.Info("Request was successful!", "status", resp.StatusCode)
}
*/

/*
package main

import (
    "log/slog"
    "net/http"
    "os"
    "time"
)

func main() {
    logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

    client := &http.Client{
        Timeout: 3 * time.Second,
    }

    url := "https://httpbin.org/get"
    logger.Info("Sending GET request", "url", url)

    resp, err := client.Get(url)
    if err != nil {
        logger.Error("Failed to make HTTP request", "error", err)
        return
    }

    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        logger.Error("Received bad status code from server", "status", resp.StatusCode)
        return
    }
    logger.Info("Request was successful!", "status", resp.StatusCode)

}
*/

/*
package main

import (
    "log/slog"
    "net/http"
    "os"
    "time"
)

func main() {
    logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

    client := &http.Client{
        Timeout: 2 * time.Second,
    }

    url := "https://cableon.ir"
    logger.Info("Sending GET request", "url", url)

    resp, err := client.Get(url)
    if err != nil {
        logger.Error("Failed to make HTTP request", "error", err)
        return
    }

    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        logger.Error("Received bad status code from server", "status", resp.StatusCode)
        return
    }

    logger.Info("Request was successful!", "status", resp.StatusCode)

}
*/

/*
package main

import (
    "log/slog"
    "net/http"
    "os"
    "time"
)

func main() {
    logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

    client := &http.Client{
        Timeout: 2 * time.Second,
    }

    url := "https://google.com"
    logger.Info("Sending GET request", "url", url)

    resp, err := client.Get(url)
    if err != nil {
        logger.Error("Failed to make HTTP request", "error", err)
        return
    }

    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        logger.Error("Received bad status code from server", "status", resp.StatusCode)
        return
    }
    logger.Info("Request was successful!", "status", resp.StatusCode)

}
*/
/*

package main

import (
    "bytes"
    "encoding/json"
    "log/slog"
    "net/http"
    "os"
    "time"
)

type UserData struct {
    Name string `json:"name"`
    Job  string `json:"job"`
}

func main() {
    logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

    user := UserData{
        Name: "Amirali",
        Job:  "Backend Developer",
    }

    jsonData, err := json.Marshal(user)
    if err != nil {
        logger.Error("Failed to marshal JSON", "error", err)
        return
    }
    bodyReader := bytes.NewBuffer(jsonData)

    url := "https://httpbin.org/post"
    req, err := http.NewRequest(http.MethodPost, url, bodyReader)
    if err != nil {
        logger.Error("Failed to create request", "error", err)
        return
    }

    req.Header.Set("Content-Type", "application/json")

    req.Header.Set("Authorization", "Bearer my-secret-token")

    client := &http.Client{
        Timeout: 5 * time.Second,
    }

    logger.Info("Sending POST request to external API...")
    resp, err := client.Do(req)
    if err != nil {
        logger.Error("Request failed", "error", err)
        return
    }

    defer resp.Body.Close()

    if resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusCreated {
        logger.Info("Data successfully posted!", "status", resp.StatusCode)
    } else {
        logger.Error("Server rejected our data", "status", resp.StatusCode)
    }
}
*/

package main

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
	"time"
)

type MessagePayload struct {
	Text string `json:"text"`
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	texter := MessagePayload{
		Text: "Heyyyyyyy BITCH!",
	}

	jsonchngr, err := json.Marshal(texter)
	if err != nil {
		logger.Error("Failed to marshal JSON", "error", err)
		return
	}

	bdyReader := bytes.NewBuffer(jsonchngr)

	url := "https://httpbin.org/post"

	req, err := http.NewRequest(http.MethodPut, url, bdyReader)
	if err != nil {
		logger.Error("Failed to create request", "error", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	req.Header.Set("X-Custom-Header", "Go-Master")

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	logger.Info("Sending PUT request to external API...")
	resp, err := client.Do(req)
	if err != nil {
		logger.Error("Request failed", "error", err)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusCreated {
		logger.Info("Data successfully posted!", "status", resp.StatusCode)
	} else {
		logger.Error("Server rejected our data", "status", resp.StatusCode)
	}
}
