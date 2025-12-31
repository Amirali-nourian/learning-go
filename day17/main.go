/*
//error wrap
package main

import (
	"errors"
	"fmt"
)

var invalidUserErr = errors.New("User is invalid!")

func ValidateUser() error {
	return invalidUserErr
}

func RegisterUser() error {
	err := ValidateUser()
	if err != nil {
		return fmt.Errorf("Error during registeration: %w", err)
	}
	return nil
}

func main() {
	regErr := RegisterUser()

	fmt.Println("Full Error:", regErr)

	valErr := errors.Unwrap(regErr)

	fmt.Println("unwrapped error:", valErr)
}
*/

/*
package main

import (
	"errors"
	"fmt"
)

var ErrLowBattery = errors.New("battery is low")

func StartCar() error {
	return ErrLowBattery
}

func Drive() error {
	err := StartCar()
	if err != nil {
		return fmt.Errorf("cannot drive: ... %w", err)
	}
	return nil
}

func main() {
	err := Drive()

	fmt.Println("Full error:", err)

	unwraperr := errors.Unwrap(err)

	fmt.Println("unwrapped error:", unwraperr)
}
*/

/*
package main

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("Not found")

func GetUser() error {
	return fmt.Errorf("database query failed: %w", ErrNotFound)
}

func main() {

	err := GetUser()

	//روش غلط: مقایسه مستقیم (چون متن‌ها فرق دارند، کار نمی‌کند)
	if err == ErrNotFound {
		fmt.Println("Direct match found!")
	}

	if errors.Is(err, ErrNotFound) {
		fmt.Println("user not found detected!")
	}

}
*/

/*
package main

import (
	"errors"
	"fmt"
)

var ErrEngineOverheat = errors.New("engine overheated")

func CheckEngine() error {
	return ErrEngineOverheat
}

func StartCar() error {
	err := CheckEngine()

	if err != nil {
		return fmt.Errorf("start failed: ...%w", err)
	}
	return nil
}

func Drive() error {
	err := StartCar()

	if err != nil {
		return fmt.Errorf("driving aborted: ...%w", err)
	}
	return nil
}

func main() {
	err := Drive()

	if errors.Is(err, ErrEngineOverheat) {
		fmt.Println("Please check the radiator water!")
	}
}
*/

/*
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("running program...")

	time.Sleep(3 * time.Second)

	panic("Something bad happened!")

}
*/

/*
package main

import "fmt"

func Div(a, b int) {
	if b == 0 {
		panic("Cannot divide by zero!")
	} else {
		fmt.Println(a / b)
	}
}

func main() {
	Div(10, 2)
	Div(10, 0)
}
*/

/*
package main

import "fmt"

func main() {
	f()

	fmt.Println("Returned normally from f.")
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("calling g.")
	g(0)
	fmt.Println("Returned normally from g.")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}
*/

package main

import "fmt"

func main() {
	fmt.Println("Resturant is open")

	SafeOrder("Ali", "Pizza")

	SafeOrder("Hacker", "Bomb")

	SafeOrder("Sara", "Pasta")

	fmt.Println("🔴 Restaurant Closed (Successfully).")

}

func SafeOrder(name string, food string) {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Oops! Something went wrong:", r)
			fmt.Println("Don't worry, system recovered.")
		}
	}()
	takeOrder(name, food)
}

func takeOrder(name string, food string) {
	fmt.Printf("Taking order %s for %s...\n", food, name)

	if food == "Bomb" {
		panic("Illegal and dangerous order!")
	}
	fmt.Println("Order registered. Enjoy!")
}

// force update commit message
