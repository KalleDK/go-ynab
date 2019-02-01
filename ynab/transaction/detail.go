package transaction

type Detail struct {
	Summary
	AccountName     string            `json:"account_name"`
	PayeeName       *string           `json:"payee_name"`
	CategoryName    *string           `json:"category_name"`
	SubTransactions []*SubTransaction `json:"subtransactions,omitempty"`
}
