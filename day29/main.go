/*package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func profilPic(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		html := `
<!DOCTYPE html>
		<html>
		<head>
			<meta charset="UTF-8">
			<title>Upload Profile Picture</title>
		</head>
		<body>
			<h2>Upload Profile Picture</h2>
			<form action="/profile-picture" method="POST" enctype="multipart/form-data">
				<input type="file" name="picture" accept="image/*">
				<br><br>
				<button type="submit">Upload</button>
			</form>
		</body>
		</html>
		`
		fmt.Fprint(w, html)
		return
	}
	if r.Method == http.MethodPost {
		err := r.ParseMultipartForm(5 << 20)
		if err != nil {
			http.Error(w, "File is too large", http.StatusBadRequest)
			return
		}
	}

	file, handler, err := r.FormFile("picture")
	if err != nil {
		http.Error(w, "Error reading file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	if handler.Size > 5<<20 {
		fmt.Fprint(w, "File is too large")
		return
	}

	err = os.MkdirAll("upload", os.ModePerm)
	if err != nil {
		http.Error(w, "Cannot create upload directory", http.StatusInternalServerError)
		return
	}
	dst, err := os.Create("uploads/" + handler.Filename)
	if err != nil {
		http.Error(w, "Cannot save file", http.StatusInternalServerError)
		return
	}

	defer dst.Close()
	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, "Error saving file", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "File uploaded successfully: %s", handler.Filename)
	return

	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

func main() {
	http.HandleFunc("/profile-picture", profilePictureHandler)

	fmt.Println("Server running on http://localhost:8080/profile-picture")
	http.ListenAndServe(":8080", nil)
}
*/

/*
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Prodcut struct {
	ID           int     `json:"product_id"`
	Name         string  `json:"name"`
	Price        float64 `json:"price"`
	Description  string  `json:"desc,omitempty"` //if it's empty don't show else show
	InternalCode string  `json:"-"` //never show
}

func productAPI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	myProduct := Prodcut{
		ID:           101,
		Name:         "Gaming Laptop",
		Price:        2500.50,
		InternalCode: "SECRET_123",
	}

	err := json.NewEncoder(w).Encode(myProduct)
	if err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/api/product", productAPI)
	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
*/

/*
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Username    string `json:"user_name"`
	Age         int8   `json:"age"`
	Password    string `json:"-"`
	PhoneNumber int64  `json:"phone_number,omitempty"`
}

func UserAPI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	myUser := User{
		Username: "Amirali",
		Age:      25,
		Password: "123124",
	}

	err := json.NewEncoder(w).Encode(myUser)
	if err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/api/me", UserAPI)
	fmt.Println("Server running on http://localhost:8080/api/me")
	http.ListenAndServe(":8080", nil)
}
*/

/*
package main

import (
	"encoding/xml"
	"fmt"
	"net/http"
)

type Employee struct {
	XMLName xml.Name `xml:"employee"`

	ID   int    `xml:"id,attr"`
	Name string `xml:"full_name"`
	Role string `xml:"role"`
}

func xmlHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/xml")

	emp := Employee{
		ID:   101,
		Name: "Amirali",
		Role: "Backend Developer",
	}

	w.Write([]byte(xml.Header))

	err := xml.NewEncoder(w).Encode(emp)
	if err != nil {
		http.Error(w, "Error encoding XML", http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/api/xml", xmlHandler)
	fmt.Println("XML Server on http://localhost:8080/api/xml")
	http.ListenAndServe(":8080", nil)
}
*/

package main

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"time"
)

type Page struct {
	Loc     string `xml:"loc"`
	LastMod string `xml:"lastmod"`
}

type UrlSet struct {
	XMLName xml.Name `xml:"urlset"`
	Urls    []Page   `xml:"url"`
}

func sitemapHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/xml")

	pages := []Page{
		{
			Loc:     "https://www.google.com/search?q=google.com",
			LastMod: time.Now().Format("2006-01-02"),
		},
		{
			Loc:     "https://www.yahoo.com",
			LastMod: time.Now().Format("2006-01-02"),
		},
	}

	urlset := UrlSet{
		Urls: pages,
	}

	xmlData, err := xml.MarshalIndent(urlset, "", "  ")
	if err != nil {
		http.Error(w, "XML Error", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(xml.Header))
	w.Write(xmlData)
}

func main() {
	http.HandleFunc("/sitemap", sitemapHandler)
	fmt.Println("XML Server on http://localhost:8080/sitemap")
	http.ListenAndServe(":8080", nil)
}
