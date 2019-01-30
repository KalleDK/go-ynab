package account

import (
	"path"

	"github.com/kalledk/go-ynab/ynab/endpoint"
)

type Account struct {
	ID               ID     `json:"id"`
	Name             string `json:"name"`
	Type             Type   `json:"type"`
	OnBudget         bool   `json:"on_budget"`
	Balance          int64  `json:"balance"`
	ClearedBalance   int64  `json:"cleared_balance"`
	UnclearedBalance int64  `json:"uncleared_balance"`
	Closed           bool   `json:"closed"`
	Deleted          bool   `json:"deleted"`
	Note             string `json:"note,omitempty"`
}

const (
	apiPath = "accounts"
)

func get(client endpoint.Getter, subpath string, model interface{}) (err error) {
	accPath := path.Join(apiPath, subpath)
	return client.Get(accPath, model)
}

func GetAccount(client endpoint.Getter, id ID) (acc Account, err error) {
	var resp Response

	err = get(client, id.String(), &resp)
	if err != nil {
		return
	}

	acc = resp.Data.Account
	return
}
