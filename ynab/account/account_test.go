package account

import (
	"log"
	"os"
	"reflect"
	"testing"

	"github.com/KalleDK/go-ynab/ynab"
	"github.com/KalleDK/go-ynab/ynab/budget"
	"github.com/KalleDK/go-ynab/ynab/client"
)

func MakeBudgetClient() *budget.Client {
	accessToken, err := ynab.NewAccessToken(os.Getenv("YNAB_TOKEN"))
	if err != nil {
		log.Fatalf("invalid token %v", err)
	}

	budgetID, err := budget.NewID(os.Getenv("YNAB_BUDGET"))
	if err != nil {
		log.Fatalf("invalid budgetid %v", err)
	}

	return budget.NewClient(budgetID, client.NewClient(accessToken))
}

func TestGetAccount(t *testing.T) {

	c := MakeBudgetClient()
	accountID, err := NewID(os.Getenv("YNAB_ACCOUNT"))
	if err != nil {
		log.Fatalf("invalid token %v", err)
	}

	validAccount := Account{
		accountID,
		"Fælleskonto",
		"creditCard",
		true,
		50000,
		0,
		50000,
		false,
		false,
		"",
	}

	type args struct {
		client APIGetter
		id     ID
	}
	tests := []struct {
		name    string
		args    args
		wantAcc Account
		wantErr bool
	}{
		{"GetValidAccount", args{c, accountID}, validAccount, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAcc, err := GetAccount(tt.args.client, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotAcc, tt.wantAcc) {
				t.Errorf("GetAccount() = %v, want %v", gotAcc, tt.wantAcc)
			}
		})
	}
}
