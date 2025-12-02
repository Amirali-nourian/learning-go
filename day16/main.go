/*package main

import (
	"fmt"
	"strconv"
)

func main() {
	var (
		n   int
		err error
		str = "Hello"
	)

	n, err = strconv.Atoi(str)

	if err != nil {
		fmt.Println("your string does not contain number:", err)
		return
	} else {
		fmt.Println("Successfully converted:", n)
	}
}
*/

/*
package main

import (
	"errors"
	"fmt"
)

func checkPassword(password string) error {
	if password != "something" {
		return errors.New("you are not authorizd!")
	}
	return nil
}

func main() {
	err := checkPassword("test")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("welcome")
	}
}
*/

/*

package main

import (
	"errors"
	"fmt"
)

func Divide(input1, input2 float64) (float64, error) {

	if input2 == 0 {
		return 0, errors.New("division by zero is not allowed")
	}
	return input1 / input2, nil
}

func main() {
	res, err := Divide(10, 2)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result", res)
	}

	res2, err2 := Divide(10, 0)

	if err2 != nil {
		fmt.Println("Error:", err2)
	} else {
		fmt.Println("Result", res2)
	}
}
*/

/*

package main

import (
	"errors"
	"fmt"
)

func CheckEntry(age int) (string, error) {
	if age < 0 {
		return "", errors.New("age cannot be negative")
	} else if age < 18 {
		return "", errors.New("you must be at least 18 years old")
	}
	return "Welcome to the club!", nil
}

func main() {
	msg, err := CheckEntry(-5)
	if err != nil {
		fmt.Println("Test -5:", err)
	} else {
		fmt.Println("Test -5:", msg)
	}

	msg2, err2 := CheckEntry(15)
	if err2 != nil {
		fmt.Println("Test 15:", err2)
	} else {
		fmt.Println("Test 15:", msg2)
	}

	msg3, err3 := CheckEntry(25)
	if err3 != nil {
		fmt.Println("Test 25:", err3)
	} else {
		fmt.Println("Test 25:", msg3)
	}
}
*/

/*
package main

import (
	"errors"
	"fmt"
)

func ValidateGrade(score int) (string, error) {
	if score < 0 {
		return "", errors.New("score cannot be negative")
	} else if score > 100 {
		return "", errors.New("score cannot be greater than 100")
	}

	return "Score accepted", nil
}

func main() {
	score, err := ValidateGrade(85)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(score)
	}

	score2, err2 := ValidateGrade(-10)
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println(score2)
	}

	score3, err3 := ValidateGrade(120)
	if err3 != nil {
		fmt.Println(err3)
	} else {
		fmt.Println(score3)
	}
}
*/

/*
//custom Errors

package main

import "fmt"

type UserValidation struct {
	Code    int
	Message string
}

func (u UserValidation) Error() string {
	return fmt.Sprintf("Error %d: %s", u.Code, u.Message)
}

func main() {
	var myErr error = UserValidation{Code: 404, Message: "Not Found!"}
	fmt.Println(myErr)
}
*/

/*
package main

import "fmt"

type PaymentError struct {
	Reason string
	Amount int
	User   string
}

func (p PaymentError) Error() string {
	return fmt.Sprintf("Payment failed for %v: %v (Amount: %v)", p.User, p.Reason, p.Amount)
}

func Pay(user string, amount int) error {
	if amount > 1000 {
		return PaymentError{
			Reason: "Insufficient funds",
			Amount: 1000,
			User:   user,
		}
	}
	fmt.Println("withdrawl successful!")
	return nil
}

func main() {
	err := Pay("Ali", 5000)

	if err != nil {
		fmt.Println(err)
	}
}
*/

/*

package main

import (
	"errors"
	"fmt"
)

type UserValidation struct {
	Message string
	Email   string
	Code    int
}

func (u *UserValidation) Error() string {
	return u.Message
}

func validationUser() error {
	return &UserValidation{
		Message: "Email already exist",
		Email:   "amirali83nori@gmail.com",
		Code:    1234,
	}
}

func main() {
	err := validationUser()

	var valErr *UserValidation

	if errors.As(err, &valErr) {

		fmt.Println("Error Message:", valErr.Message)
		fmt.Println("User Email:", valErr.Email)
		fmt.Println("Error Code:", valErr.Code)
	} else {
		fmt.Println("Error type was not UserValidation")
	}
}
*/

/*
package main

import (
	"errors"
	"fmt"
)

type UploadError struct {
	Filename string
	Size     int
}

func (u *UploadError) Error() string {
	return fmt.Sprintf("upload error: file %s is too large (%dMB)", u.Filename, u.Size)
}

func UploadFile(name string, size int) error {
	if size > 20 {
		return &UploadError{
			Filename: name,
			Size:     size,
		}
	}
	return nil
}

func main() {
	err := UploadFile("movie.mp4", 50)

	var uploadErr *UploadError

	if errors.As(err, &uploadErr) {
		fmt.Println("⚠️ Custom Error Detected!")
		fmt.Println("File Name:", uploadErr.Filename)
		fmt.Println("Size:", uploadErr.Size)
	} else if err != nil {
		fmt.Println("Generic Error:", err)
	} else {
		fmt.Println("Success!")
	}
}
*/

package main

import (
	"errors"
	"fmt"
)

type ValidationError struct {
	Field  string
	Reason string
}

type DatabaseError struct {
	Code    int
	Message string
}

func (v *ValidationError) Error() string {
	return fmt.Sprintf("Validation Error -> Field: %s, Reason: %s", v.Field, v.Reason)
}

func (d *DatabaseError) Error() string {
	return fmt.Sprintf("Database Error -> Code: %d, Message: %s", d.Code, d.Message)
}

func PublishArticle(title string, content string) error {
	if title == "" {
		return &ValidationError{
			Field:  "Title",
			Reason: "cannot be empty",
		}
	}
	if len(content) < 10 {
		return &ValidationError{
			Field:  "Content",
			Reason: "length too short!",
		}
	}
	if title == "server_down" {
		return &DatabaseError{
			Code:    500,
			Message: "connection failed",
		}
	}
	return nil
}

func main() {
	scenarios := []struct {
		Title   string
		Content string
	}{
		{"", "Some content"},
		{"My Article", "Hi"},
		{"server_down", "Valid content here"},
		{"Go Lang", "This is a great language"},
	}
	for _, s := range scenarios {
		fmt.Printf("--- Publishing: '%s' ---\n", s.Title)

		err := PublishArticle(s.Title, s.Content)

		var valErr *ValidationError
		var dbErr *DatabaseError
		if errors.As(err, &valErr) {
			fmt.Printf("❌ VALIDATION FAILED! Please fix the '%s'. (%s)\n", valErr.Field, valErr.Reason)
		} else if errors.As(err, &dbErr) {
			fmt.Printf("🔥 CRITICAL ERROR! Call admin. Code: %d\n", dbErr.Code)
		} else if err != nil {
			fmt.Println("Unknown Error:", err)
		} else {
			fmt.Println("✅ Article published successfully.")
		}
		fmt.Println()
	}
}
