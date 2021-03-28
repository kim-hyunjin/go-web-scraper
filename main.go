package main

import (
	"fmt"

	"github.com/kim-hyunjin/go-scraper/banking"
	"github.com/kim-hyunjin/go-scraper/dict"
)

func main() {
	account := banking.NewAccount("hyunjin")
	account.Deposit(5000)
	fmt.Println(account.Owner(), account.Balance())
	account.Withdraw(3000)
	account.ChangeOwner("gildong")
	fmt.Println(account.Owner(), account.Balance())
	fmt.Println(account)
	// err := account.Withdraw(3000)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	dictionary := dict.Dictionary{"first": "First word"}
	dictionary["hello"] = "hello"
	definition, err := dictionary.Search("first")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}

	err = dictionary.Add("hello2", "greeting")
	if err != nil {
		fmt.Println(err)
	}
	definition, err = dictionary.Search("hello2")
	fmt.Println(definition)

	err = dictionary.Update("hello2", "Hola!")
	if err != nil {
		fmt.Println(err)
	}
	definition, _ = dictionary.Search("hello2")
	fmt.Println(definition)

	dictionary.Delete("hello2")

	err = dictionary.Update("hello2", "howdy!")
	if err != nil {
		fmt.Println(err)
	}
}