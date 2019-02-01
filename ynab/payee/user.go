package payee

type Payee struct {
	ID                ID     `json:"id"`
	Name              string `json:"name"`
	TransferAccountID string `json:"transfer_account_id"`
	Deleted           bool   `json:"deleted"`
}

type Payees = []Payee
