package main

import (
	"fmt"
)

type Expense struct {
	Title  string
	Amount int
}

var expenses []Expense

func addexpense() {
	var t string
	var a int
	fmt.Println("Title:")
	fmt.Scanln(&t)
	fmt.Println("Amount:")
	fmt.Scanln(&a)

	rxp := Expense{t, a}
	expenses = append(expenses, rxp)
	fmt.Println("expenses added")
}

func showexpense() {

	if len(expenses) == 0 {
		fmt.Println("No found")
		return
	}
	for i, e := range expenses {
		fmt.Println(i+1, e.Title, e.Amount)
	}
}
func totalexp() {
	total := 0
	for _, e := range expenses {
		total += e.Amount
	}
	fmt.Println(total)

}
func main() {
	for {
		fmt.Println("Expense tracker")
		fmt.Println("What you want")
		fmt.Println("add expense")
		fmt.Println("total amount")
		fmt.Println("show expense")
		fmt.Println("enter: add, show, total, e")

		var s string
		fmt.Scanln(&s)
		if s == "add" {
			addexpense()
		} else if s == "show" {
			showexpense()
		} else if s == "total" {
			totalexp()
		} else if s == "e" {
			return
		} else {

			fmt.Println("Not found")
			return
		}

	}
}
