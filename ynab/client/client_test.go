package client

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/kalledk/go-ynab/ynab/account"
	"github.com/kalledk/go-ynab/ynab/api"
	"github.com/kalledk/go-ynab/ynab/budget"
	"github.com/kalledk/go-ynab/ynab/payee"
	"github.com/kalledk/go-ynab/ynab/transaction"
)

func SprintJSON(model interface{}) string {
	data, err := json.MarshalIndent(model, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	return string(data)
}

func TestUser(t *testing.T) {
	token, _ := api.NewAccessToken(os.Getenv("YNAB_TOKEN"))
	c := NewClient(token)
	user, err := GetUser(c)
	fmt.Printf("Error = %v\n", err)
	fmt.Printf("User = %v\n", SprintJSON(user))

}

func TestBudgetSettings(t *testing.T) {
	token, _ := api.NewAccessToken(os.Getenv("YNAB_TOKEN"))
	c := NewClient(token)

	budgetID, _ := budget.NewID(os.Getenv("YNAB_BUDGET"))
	settings, err := GetBudgetSettings(c, budgetID)
	fmt.Printf("Error = %v\n", err)
	fmt.Printf("Settings = %v\n", SprintJSON(settings))
}

func TestBudgetTransaction(t *testing.T) {
	token, _ := api.NewAccessToken(os.Getenv("YNAB_TOKEN"))
	budgetID, _ := budget.NewID(os.Getenv("YNAB_BUDGET"))
	transactionID, _ := transaction.NewID(os.Getenv("YNAB_TRANSACTION"))
	client := NewClient(token)
	budgetClient := client.Budgets().Budget(budgetID)
	transaction, err := GetTransaction(budgetClient, transactionID)
	fmt.Printf("Error = %v\n", err)
	fmt.Printf("Transaction = %v\n", SprintJSON(transaction))
}

func TestBudgetAddTransaction(t *testing.T) {
	token, _ := api.NewAccessToken(os.Getenv("YNAB_TOKEN"))
	budgetID, _ := budget.NewID(os.Getenv("YNAB_BUDGET"))
	accountID, _ := account.NewID(os.Getenv("YNAB_ACCOUNT"))
	payeeID, _ := payee.NewID(os.Getenv("YNAB_PAYEE"))
	client := NewClient(token)
	transactionClient := client.Budgets().Budget(budgetID).Transactions()
	s := transaction.Transaction{
		AccountID: accountID,
		Date:      "2019-02-01",
		Amount:    12370,
		PayeeID:   payeeID,
		Memo:      "Serial 1234",
	}
	reply, err := transactionClient.Add(s)
	fmt.Printf("Error = %v\n", err)
	fmt.Printf("Transaction = %v\n", SprintJSON(reply))
}

func TestBudgetAddTransactions(t *testing.T) {
	token, _ := api.NewAccessToken(os.Getenv("YNAB_TOKEN"))
	budgetID, _ := budget.NewID(os.Getenv("YNAB_BUDGET"))
	accountID, _ := account.NewID(os.Getenv("YNAB_ACCOUNT"))
	payeeID, _ := payee.NewID(os.Getenv("YNAB_PAYEE"))
	client := NewClient(token)
	transactionClient := client.Budgets().Budget(budgetID).Transactions()
	s := []transaction.Transaction{
		transaction.Transaction{
			AccountID: accountID,
			Date:      "2019-02-01",
			Amount:    12390,
			PayeeID:   payeeID,
			Memo:      "Serial 1234",
		},
	}
	reply, err := transactionClient.AddList(s)
	fmt.Printf("Error = %v\n", err)
	fmt.Printf("Transaction = %v\n", SprintJSON(reply))
}

func TestBudgetPayees(t *testing.T) {
	token, _ := api.NewAccessToken(os.Getenv("YNAB_TOKEN"))
	budgetID, _ := budget.NewID(os.Getenv("YNAB_BUDGET"))

	client := NewClient(token)
	budgetClient := client.Budgets().Budget(budgetID)
	payeesClient := budgetClient.Payees()

	payees, err := payeesClient.Get()

	fmt.Printf("Error = %v\n", err)
	fmt.Printf("Payees = %v\n", SprintJSON(payees))
}
