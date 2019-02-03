package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/kalledk/go-ynab/ynab/account"
	"github.com/kalledk/go-ynab/ynab/api"
	"github.com/kalledk/go-ynab/ynab/budget"
	"github.com/kalledk/go-ynab/ynab/client"
)

func main() {
	accessToken, err := api.NewAccessToken(os.Getenv("YNAB_TOKEN"))
	if err != nil {
		log.Fatalf("invalid token, %v", err)
	}
	c := client.NewClient(accessToken)

	budgetID, err := budget.NewID(os.Getenv("YNAB_BUDGET"))
	if err != nil {
		log.Fatalf("invalid budget id, %v", err)
	}
	bc := c.Budgets().Budget(budgetID)

	accountID, err := account.NewID(os.Getenv("YNAB_ACCOUNT"))
	if err != nil {
		log.Fatalf("invalid account id, %v", err)
	}
	ac := bc.Accounts().Account(accountID)

	myaccount, err := ac.Get()
	if err != nil {
		log.Fatalf("invalid response %v", err)
	}

	json, err := json.MarshalIndent(myaccount, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(json))

}
