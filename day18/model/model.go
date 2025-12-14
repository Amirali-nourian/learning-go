//Encapsulation

package model

import "fmt"

type User struct {
	FirstName string
	LastName  string
	fullName  string
}

func (u *User) SetFullName() {
	u.fullName = u.FirstName + " " + u.LastName
}

func (u *User) GetFullName() string {
	return u.fullName
}

//---------------------------------------------------------------------------------bank-------------------------------------------------------------------

type BankAccount struct {
	Owner   string
	balance float64
}

func NewAccount(owner string, balance float64) *BankAccount {
	return &BankAccount{
		Owner:   owner,
		balance: balance,
	}
}

func (b *BankAccount) Deposit(amount float64) {
	b.balance += amount

}

func (b *BankAccount) Withdraw(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("invalid amount")
	} else if b.balance < amount {
		return fmt.Errorf("insufficient balance")
	}
	b.balance -= amount
	fmt.Println("withdraw was successful!")
	return nil
}

func (b BankAccount) GetBalance() float64 {
	return b.balance
}
