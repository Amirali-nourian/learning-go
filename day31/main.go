/*
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
*/

/*
package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.OpenFile("transaction.csv", os.O_RDONLY, 0777)
	if err != nil {
		log.Fatalf("Unable to open file: %v", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)

	_, err = reader.Read()
	if err != nil {
		log.Fatalf("Error reading header: %v", err)
	}

	var sum float64 = 0

	for {
		record, err := reader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error reading CSV record: %v", err)
		}

		amountStr := record[1]

		amount, err := strconv.ParseFloat(amountStr, 64)
		if err != nil {
			log.Fatalf("Error converting amount: %v", err)
		}
		sum += amount
	}

	fmt.Printf("Sum of Amount column: %.2f\n", sum)
}
*/

/*
package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("transactions.csv")
	if err != nil {
		log.Fatalf("Critical error: Unable to open file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	_, err = reader.Read()
	if err == io.EOF {
		log.Println("File is empty.")
		return
	}
	if err != nil {
		log.Fatalf("Error reading header: %v", err)
	}

	var sum float64 = 0.0
	lineCount := 1

	for {
		record, err := reader.Read()
		lineCount++

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("Skipping corrupted line %d: %v\n", lineCount, err)
			continue
		}

		if len(record) < 2 {
			log.Printf("Skipping line %d: insufficient columns\n", lineCount)
			continue
		}

		amountStr := record[1]
		amount, err := strconv.ParseFloat(amountStr, 64)
		if err != nil {
			log.Printf("Skipping line %d: invalid amount format '%s' (%v)\n", lineCount, amountStr, err)
			continue
		}

		sum += amount
	}

	fmt.Printf("Total processed sum: %.2f\n", sum)
}
*/

package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

type Data struct {
	ID      int
	Payload string
	Metrics map[string]float64
}

func main() {
	input := Data{
		ID:      1,
		Payload: "Internal System Message",
		Metrics: map[string]float64{"cpu": 0.5, "memory": 0.8},
	}

	var network bytes.Buffer

	encoder := gob.NewEncoder(&network)

	if err := encoder.Encode(input); err != nil {
		log.Fatalf("Gob encoding failed: %v", err)
	}

	fmt.Printf("Encoded binary size: %d bytes\n", network.Len())

	var output Data
	decoder := gob.NewDecoder(&network)

	if err := decoder.Decode(&output); err != nil {
		log.Fatalf("Gob decoding failed: %v", err)
	}

	fmt.Printf("Decoded Result: %+v\n", output)
}
