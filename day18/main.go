/*//Encapsulation

package main

import (
	"day18/model"
)

func main() {
	test := model.User{
		FirstName: "Amirali",
		LastName:  "Nourian",
		//fullName:  "Jhon L",
	}

	test.SetFullName()
	println(test.GetFullName())
}
*/

/*
package main

import "fmt"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	width, height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * 3.14
}

func PrintArea(s Shape) {
	fmt.Println("Area is:", s.Area())
}

func main() {
	rectangle := Rectangle{
		width:  5,
		height: 10,
	}

	circle := Circle{
		Radius: 7,
	}

	PrintArea(rectangle)
	PrintArea(circle)
}
*/

/*
//interface
package main

import "fmt"

type Payer interface {
	Pay(amount int)
}

type Wallet struct {
	Cash int
}

func (w Wallet) Pay(amount int) {
	fmt.Println("Paying", amount, "from Wallet")
}

type CreditCard struct {
	Number string
}

func (c CreditCard) Pay(amount int) {
	fmt.Println("Paying", amount, "with Credit Card:", c.Number)
}

func ProcessPayment(p Payer, amount int) {
	p.Pay(amount)
}

func main() {
	myWallet := Wallet{
		Cash: 100,
	}

	myCard := CreditCard{
		Number: "1234-5678",
	}

	ProcessPayment(myWallet, 50)
	ProcessPayment(myCard, 1000)
}
*/

/*
// composition
package main

import "fmt"

type SingerInfo struct {
	Name string
	Age  int
}

type Album struct {
	Title string
	SingerInfo
}

func main() {
	myAlbum := Album{
		Title: "Jane Javani",
		SingerInfo: SingerInfo{
			Name: "Ebrahim Hamedi",
			Age:  74,
		},
	}
	fmt.Println("Album:", myAlbum.Title)
	println("Singer:", myAlbum.Name)
}
*/

/*
// Structs & Methods
package main

import "fmt"

type Cars struct {
	Brand string
	Speed int
}

func (c Cars) Drive(amount int) {
	c.Speed += amount
	fmt.Printf("Driving %s at speed: %d km/h\n", c.Brand, c.Speed)
}

func main() {
	myCar := Cars{
		Brand: "Mercedes-Benz",
		Speed: 0,
	}

	myCar.Drive(180)
	myCar.Drive(100)
}
*/

/*
// Encapsulation
package main

import "fmt"

type BankAccount struct {
	Owner   string
	balance int
}

func NewAccount(owner string, amount int) *BankAccount {
	return &BankAccount{
		Owner:   owner,
		balance: amount,
	}
}

func (b *BankAccount) GetBalance() int {
	return b.balance
}

func main() {
	myAccount := BankAccount{
		Owner:   "Amirali",
		balance: 3000000,
	}
	fmt.Println(myAccount.GetBalance())
}
*/

/*
//Composition

package main

import "fmt"

type Processor struct {
	Cores int
}

func (p Processor) Compute() {
	fmt.Println("Processing...")
}

type Computer struct {
	Model string
	Processor
}

func main() {
	pc := Computer{
		Model:     "Corsair",
		Processor: Processor{Cores: 16},
	}
	pc.Compute()
}
*/

/*
// interface
package main

import "fmt"

type Printer interface {
	Print()
}

type Epson struct{}

func (e Epson) Print() {
	fmt.Println("Printing from Epson...")
}

func main() {
	var p Printer
	p = Epson{}
	p.Print()
}
*/

/*
package main

import "fmt"

type SmartPhone struct {
	Model   string
	Battery int
}

func (p SmartPhone) ShowInfo() {
	p.Model = "Fake Model"
	fmt.Printf("Inside Value Method: %s - Battery: %d%%\n", p.Model, p.Battery)
}

func (p *SmartPhone) Charge() {
	p.Battery = 100
	fmt.Println("Inside Pointer Method: Phone Charged to 100%!")
}

func main() {
	myPhone := SmartPhone{
		Model:   "Iphone 17 pro max",
		Battery: 28,
	}
	fmt.Println("--- Step 1: Using Value Receiver ---")
	myPhone.ShowInfo()

	fmt.Println("Original Phone Model:", myPhone.Model)

	fmt.Println("\n--- Step 2: Using Pointer Receiver ---")
	fmt.Println("Current Battery:", myPhone.Battery)

	myPhone.Charge()
	fmt.Println("Original Phone Battery:", myPhone.Battery)
}
*/

/*
package main

import "fmt"

type Post struct {
	Title string
	Likes int
}

func (p *Post) AddLike() {
	p.Likes++
}

func (p Post) ShowInfo() {
	fmt.Printf("Title: %s | Likes: %d", p.Title, p.Likes)
}

func main() {
	post := Post{
		Title: "Go Lang",
		Likes: 0,
	}

	for i := 1; i <= 3; i++ {
		post.AddLike()
	}

	post.ShowInfo()
}
*/

package main

import (
	"day18/model"
	"fmt"
)

func main() {
	myAcc := model.NewAccount("Amirali", 1000)

	myAcc.Deposit(500)
	err := myAcc.Withdraw(2000)
	if err != nil {
		fmt.Println("Error:", err)
	}
	myAcc.Withdraw(300)
	fmt.Println(myAcc.GetBalance())
}

// force update commit message
