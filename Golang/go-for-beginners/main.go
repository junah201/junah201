package main

import (
	"fmt"
	"log"

	"github.com/junah201/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("Junah")
	account.Deposit(10)
	fmt.Println(account)
	err := account.Withdraw(20)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(account)
}