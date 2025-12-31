package main

import (
	"fmt"
)

func main() {
	var a, b number = 5, 7

	fmt.Println(add(a, b))

	//type alias
	var c, d int = 11, 6
	fmt.Println(Add(c, d))

	//struct
	laptop := Product{
		ID:          101,
		Name:        "Asus ROG Strix",
		Price:       65000000,
		IsAvailable: true,
		Tags:        []string{"Gaming", "Laptop", "High-Performance"},
	}

	var mouse Product
	mouse.Name = "Logitech Mouse"
	mouse.Price = 1500000

	fmt.Println("Product 1:", laptop)
	fmt.Println("Product 2:", mouse)

	laptop.Price = 67000000

	fmt.Println("Updated Price for", laptop.Name, "is", laptop.Price)

	if laptop.IsAvailable {
		fmt.Println("Ready to ship!")
	}

	//my_struct
	employees := []Employee{
		{
			FirstName:     "Ali",
			LastName:      "Rezaei",
			BaseSalary:    12000000,
			OvertimeHours: 10,
		},
		{
			FirstName:     "Sara",
			LastName:      "Mohammadi",
			BaseSalary:    15000000,
			OvertimeHours: 5,
		},
	}

	for _, emp := range employees {
		payment := emp.CalculatePayment()
		fmt.Printf("Rights for %s: %d Tomans\n", emp.FullNAME(), payment)
	}

	ali := Employee{FirstName: "Ali", BaseSalary: 1000}
	fmt.Println("Ali's salary Before raise:", ali.BaseSalary)
	ali.GiveRaise(500)
	fmt.Println("Ali's salary After raise:", ali.BaseSalary)

	//Method_deployment
	mylamp := Lamp{
		IsOn:  true,
		Color: "Blue",
	}
	mylamp.ReportStatus()

	//Game
	Scorpion := Player{
		Name:   "Scorpion",
		Health: 100,
	}

	Scorpion.ShowStatus()
	Scorpion.TakeDamage(20)
	Scorpion.ShowStatus()
}

//out of main

type number int

func add(a, b number) number {
	return a + b
}

//type alias

type numbers = int

func Add(a, b numbers) numbers {
	return a + b
}

type Product struct {
	ID          int
	Name        string
	Price       int
	IsAvailable bool
	Tags        []string
}

type Employee struct {
	FirstName     string
	LastName      string
	BaseSalary    int
	OvertimeHours int
}

func (e Employee) CalculatePayment() int {
	return (e.OvertimeHours * 200000) + e.BaseSalary
}

func (fullName Employee) FullNAME() string {
	return fullName.FirstName + " " + fullName.LastName
}

func (e *Employee) GiveRaise(amount int) {
	e.BaseSalary = e.BaseSalary + amount
	fmt.Println("Ali's salary Raise given inside method:", e.BaseSalary)
}

//Method deployment

type Lamp struct {
	IsOn  bool
	Color string
}

//Method deployment

func (l Lamp) ReportStatus() {
	if l.IsOn {
		fmt.Println("Lamp is ON and color is", l.Color)
	} else {
		fmt.Println("Lamp is OFF")
	}
}

//Game with pointer

type Player struct {
	Name   string
	Health int
}

func (p Player) ShowStatus() {
	fmt.Printf("Player %s is at %d HP.\n", p.Name, p.Health)
}

func (p *Player) TakeDamage(damage int) {
	p.Health -= damage
}

// force update commit message
