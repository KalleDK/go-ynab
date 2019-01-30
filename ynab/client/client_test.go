package client

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/kalledk/go-ynab/ynab"
	"github.com/kalledk/go-ynab/ynab/budget"
	//"net/url"
)

func TestAccountID(t *testing.T) {
	at, err := ynab.NewAccessToken(os.Getenv("YNAB_TOKEN"))
	if err != nil {
		t.Fatalf("invalid token")
	}
	c := NewClient(at)
	fmt.Println(c)
}

func TestError(t *testing.T) {
	at, err := ynab.NewAccessToken(os.Getenv("YNAB_TOKEN"))
	if err != nil {
		t.Fatalf("invalid token")
	}
	c := NewClient(at)

	path := "user"
	var userResponse ynab.UserResponse
	err = c.Do(http.MethodGet, path, &userResponse, nil)
	if err != nil {
		t.Fatalf("invalid response %v", err)
	}
	fmt.Println(userResponse.Data.User)
}

func TestBudgets(t *testing.T) {
	at, err := ynab.NewAccessToken(os.Getenv("YNAB_TOKEN"))
	if err != nil {
		t.Fatalf("invalid token")
	}
	c := NewClient(at)

	path := "budgets"
	var budgetResponse budget.SummaryResponse
	err = c.Get(path, &budgetResponse)
	if err != nil {
		t.Fatalf("invalid response %v", err)
	}

	json, err := json.MarshalIndent(budgetResponse.Data.Budgets, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	_ = json

	//fmt.Println(string(json))
	//fmt.Println(budgetResponse.Data.Budgets)
}
