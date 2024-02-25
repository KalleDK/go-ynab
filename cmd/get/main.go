package main

import (
	"log"
	"os"

	"github.com/KalleDK/go-ynab/pkg/ynab"
)

func main() {
	token := ynab.Token(os.Getenv("TOKEN"))

	client := ynab.Config{
		Token: token,
	}.NewClient()

	budgets, err := client.GetBudgets()
	if err != nil {
		log.Fatal(err)
	}

	var budget_id ynab.BudgetID
	for _, b := range budgets {
		if b.Name == "Demo" {
			budget_id = b.ID
		}
	}
	if budget_id.IsEmpty() {
		log.Fatal("No budget found")
	}

	var account_id ynab.AccountID
	accounts, err := client.GetAccounts(budget_id)
	if err != nil {
		log.Fatal(err)
	}
	for _, a := range accounts {
		if a.Name == "Test" {
			account_id = a.ID
		}
	}
	if account_id.IsEmpty() {
		log.Fatal("No account found")
	}

	var payee_id ynab.PayeeID
	payees, err := client.GetPayees(budget_id)
	if err != nil {
		log.Fatal(err)
	}
	for _, p := range payees {
		if p.Name == "Coop" {
			payee_id = p.ID
		}
	}
	if payee_id.IsEmpty() {
		log.Fatal("No payee found")
	}

	var category_id ynab.CategoryID
	catGroups, err := client.GetCategories(budget_id)
	if err != nil {
		log.Fatal(err)
	}
	for _, cg := range catGroups {
		for _, c := range cg.Categories {
			if c.Name == "Groceries" {
				category_id = c.ID
			}
		}
	}
	if category_id.IsEmpty() {
		log.Fatal("No category found")
	}

	var transaction_id ynab.TransactionID
	transactions, err := client.GetAccountTransactions(budget_id, account_id)
	if err != nil {
		log.Fatal(err)
	}
	for _, t := range transactions {
		if t.PayeeID != nil && *t.PayeeID == payee_id {
			transaction_id = t.ID
		}
	}
	if transaction_id.IsEmpty() {
		log.Fatal("No transaction found")
	}

}
