package main

import "fmt"

type transaction struct {
	ID             int
	Account        string
	Balance        float32
	ReceivedAmount float32
	ReceivedFrom   string
	SendAmount     float32
	SendTo         string
}

var transactions = []transaction{}

func getMultipleBalance(id int) []transaction {
	return transactions
}

func getBalance(id int) (string, float32) {
	var account string
	var balance float32
	for _, t := range transactions {
		if t.ID == id {
			account = t.Account
			balance = t.Balance
		}
	}
	return account, balance
}

func setDeposit(id int, amount float32) bool {
	if amount < 1 {
		return false
	}
	trx := make([]transaction, 0)
	trx = append(trx, transaction{ID: id, ReceivedAmount: amount})
	return true
}

func setWithdrawal(id int, amount float32) bool {
	if amount < 1 {
		return false
	}
	trx := make([]transaction, 0)
	trx = append(trx, transaction{ID: id, SendAmount: amount})
	return true
}

func main() {
	t1 := getMultipleBalance(1)
	fmt.Println("Init", t1)
	t2 := setDeposit(1, 10000)
	fmt.Println("T2", t2)
	t3 := setWithdrawal(1, 5000)
	fmt.Println("T3", t3)
	account, balance := getBalance(1)
	fmt.Println("Account", account)
	fmt.Println("Balance", balance)
}
