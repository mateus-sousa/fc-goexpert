package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Account struct {
	Number  int `json:"n"`
	Balance int `json:"b"`
}

func main() {
	account := Account{
		Number:  1,
		Balance: 100,
	}
	res, err := json.Marshal(account)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(res))
	err = json.NewEncoder(os.Stdout).Encode(account)
	if err != nil {
		fmt.Println(err)
	}
	jsonPure := []byte(`{"n": 2, "b": 200}`)
	var account2 Account
	err = json.Unmarshal(jsonPure, &account2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(account2.Balance)
}
