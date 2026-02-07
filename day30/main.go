/*
package main

import (
	"encoding/xml"
	"net/http"
)

type Page struct {
	Loc     string `xml:"loc"`
	LastMod string `xml:"lastmod"`
}

type UrlSet struct {
	XMLName xml.Name `xml:"urlset"`
	Pages   []Page   `xml:"url"`
}

func sitemapHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/xml")

	page1 := Page{
		Loc:     "https://www.google.com/search?q=google.com",
		LastMod: "2026-02-07",
	}

	page2 := Page{
		Loc:     "https://www.yahoo.com",
		LastMod: "2026-02-07",
	}

	urlSet := UrlSet{
		Pages: []Page{page1, page2},
	}

	w.Write([]byte(xml.Header))

	encoder := xml.NewEncoder(w)
	encoder.Indent("", "  ")
	encoder.Encode(urlSet)
}

func main() {
	http.HandleFunc("/sitemap", sitemapHandler)
	http.ListenAndServe(":8080", nil)
}
*/

/*
package main

import (
	"encoding/csv"
	"log"
	"os"
)

type Product struct {
	ID    string
	Name  string
	Price string
}

func main() {
	products := []Product{
		{ID: "1", Name: "Laptop", Price: "1200"},
		{ID: "2", Name: "Mouse", Price: "25"},
		{ID: "3", Name: "Monitor", Price: "300"},
	}

	file, err := os.Create("products.csv")
	if err != nil {
		log.Fatalf("Failed to create file: %v", err)
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	header := []string{"ID", "Name", "Price"}
	if err := writer.Write(header); err != nil {
		log.Fatalf("Error writing to header: %v", err)
	}

	for _, p := range products {
		record := []string{p.ID, p.Name, p.Price}
		if err := writer.Write(record); err != nil {
			log.Fatalf("Error writing record: %v", err)
		}
	}

	writer.Flush()
	if err := writer.Error(); err != nil {
		log.Fatalf("Error flushing data to file: %v", err)
	}

	log.Println("CSV file created successfully.")
}
*/

/*
package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"
)

type Transaction struct {
	UserID      string
	Amount      float64
	Date        time.Time
	Description string
}

func main() {
	transactions := []Transaction{

		{UserID: "1", Amount: 285000, Date: time.Now(), Description: "Groceries"},

		{UserID: "2", Amount: 25000, Date: time.Now(), Description: "Taxi"},

		{UserID: "3", Amount: 2000, Date: time.Now(), Description: "Coffee"},

		{UserID: "4", Amount: 5000, Date: time.Now(), Description: "Payment, for rent"}, // شامل ویرگول

		{UserID: "5", Amount: 85000, Date: time.Now(), Description: "Internet bill"},
	}

	file, err := os.OpenFile("transaction.csv", os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0644)
	if err != nil {
		log.Fatalf("Failed to create file: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	header := []string{"UserID", "Amount", "Date", "Description"}
	if err := writer.Write(header); err != nil {
		log.Fatalf("Failed to write header: %v", err)
	}

	for _, t := range transactions {
		record := []string{
			t.UserID,
			fmt.Sprintf("%.0f", t.Amount),
			t.Date.Format("2006-01-02 15:04:05"),
			t.Description,
		}

		if err := writer.Write(record); err != nil {
			log.Fatalf("Failed to write record: %v", err)
		}
	}

	writer.Flush()
	if err := writer.Error(); err != nil {
		log.Fatalf("Error flushing data: %v", err)
	}

	log.Println("CSV generation completed successfully.")
}
*/

package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("transaction.csv", os.O_RDONLY, 0644)
	if err != nil {
		log.Fatalf("Unable to open file: %v", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error reading CSV record: %v", err)
		}

		fmt.Printf("User: %s, Amount: %s, Desc: %s\n", record[0], record[1], record[3])
	}
}
