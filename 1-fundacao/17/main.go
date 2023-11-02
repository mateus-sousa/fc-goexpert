package main

import "fmt"

type Account struct {
	balance int
}

func (c Account) withDrawnSimulate(withdrawnValue int) {
	c.balance -= withdrawnValue
	fmt.Println(c.balance)
}

func (c *Account) withdrawn(withdrawnValue int) {
	c.balance -= withdrawnValue
	fmt.Println(c.balance)
}

func main() {
	myAccount := Account{
		balance: 500,
	}

	myAccount.withDrawnSimulate(200)
	fmt.Println(myAccount.balance)

	myAccount.withdrawn(200)
	fmt.Println(myAccount.balance)
}
