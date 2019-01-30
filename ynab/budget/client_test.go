package budget

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/KalleDK/go-ynab/ynab"
	"github.com/KalleDK/go-ynab/ynab/account"
	"github.com/KalleDK/go-ynab/ynab/client"
)

func TestBudgets(t *testing.T) {
	envToken := os.Getenv("YNAB_TOKEN")
	at, err := ynab.NewAccessToken(envToken)
	if err != nil {
		t.Fatalf("invalid token")
	}
	c := client.NewClient(at)

	path := "./budgets"
	var budgetResponse SummaryResponse
	err = c.Get(path, &budgetResponse)
	if err != nil {
		t.Fatalf("invalid response %v", err)
	}

	bc := NewClient(budgetResponse.Data.Budgets[0].ID, c)
	settResp, err := bc.GetSettings()
	if err != nil {
		t.Fatalf("invalid response %v", err)
	}

	json, err := json.MarshalIndent(settResp, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	_ = json

	//fmt.Println(string(json))
	//fmt.Println(budgetResponse.Data.Budgets)
}

func TestAccount(t *testing.T) {
	envToken := os.Getenv("YNAB_TOKEN")
	at, err := ynab.NewAccessToken(envToken)
	if err != nil {
		t.Fatalf("invalid token")
	}
	c := client.NewClient(at)

	budgetID, _ := NewID(os.Getenv("YNAB_BUDGET"))
	accID, _ := account.NewID(os.Getenv("YNAB_ACCOUNT"))

	bc := NewClient(budgetID, c)
	settResp, err := bc.GetAccount(accID)
	if err != nil {
		t.Fatalf("invalid response %v", err)
	}

	json, err := json.MarshalIndent(settResp, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	_ = json

	fmt.Println(string(json))
	t.Error("No")
}
