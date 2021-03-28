package main

import (
	"fmt"
	"log"

	"github.com/kim-hyunjin/go-scraper/banking"
)

func main() {
	account := banking.NewAccount("hyunjin")
	account.Deposit(5000)
	fmt.Println(account.Owner(), account.Balance())
	account.Withdraw(3000)
	account.ChangeOwner("gildong")
	fmt.Println(account.Owner(), account.Balance())
	fmt.Println(account)
	err := account.Withdraw(3000)
	if err != nil {
		log.Fatal(err)
	}
}