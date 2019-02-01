package transaction

type Type string

const (
	// TypeTransaction identifies a hybrid transaction as transaction
	TypeTransaction Type = "transaction"
	// TypeSubTransaction identifies a hybrid transaction as sub-transaction
	TypeSubTransaction Type = "subtransaction"
)
