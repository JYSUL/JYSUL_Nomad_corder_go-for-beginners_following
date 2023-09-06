package main

import (
	"fmt"

	"github.com/JYSUL/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("JYSUL")
	fmt.Println(account)
	account.Deposit(50)
	fmt.Println(account.Balance())
	err := account.Withdraw(60)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(account.Balance())
}
