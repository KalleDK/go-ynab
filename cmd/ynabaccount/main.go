package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/kalledk/go-ynab/ynab"
	"github.com/kalledk/go-ynab/ynab/account"
	"github.com/kalledk/go-ynab/ynab/budget"
	"github.com/kalledk/go-ynab/ynab/client"
)

func main() {
	accessToken, err := ynab.NewAccessToken(os.Getenv("YNAB_TOKEN"))
	if err != nil {
		log.Fatalf("invalid token, %v", err)
	}
	c := client.NewClient(accessToken)

	budgetID, err := budget.NewID(os.Getenv("YNAB_BUDGET"))
	if err != nil {
		log.Fatalf("invalid budget id, %v", err)
	}
	bc := budget.NewClient(budgetID, c)

	accountID, err := account.NewID(os.Getenv("YNAB_ACCOUNT"))
	if err != nil {
		log.Fatalf("invalid account id, %v", err)
	}

	myaccount, err := account.GetAccount(bc, accountID)
	if err != nil {
		log.Fatalf("invalid response %v", err)
	}

	json, err := json.MarshalIndent(myaccount, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(json))

}
