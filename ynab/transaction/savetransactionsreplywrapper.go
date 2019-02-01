package transaction

type SaveTransactionReplyWrapper struct {
	IDs                []ID     `json:"transaction_ids"`
	Transaction        Detail   `json:"transaction"`
	Transactions       []Detail `json:"transactions"`
	DuplicateImportIDs []string `json:"duplicate_import_ids"`
}

type SaveTransactionWrapper struct {
	Transaction  SaveTransaction   `json:"transaction,omitempty"`
	Transactions []SaveTransaction `json:"transactions,omitempty"`
}
