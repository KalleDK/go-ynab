package transaction

type ClearingStatus string

const (
	// ClearingStatusUncleared identifies an uncleared transaction
	ClearingStatusUncleared ClearingStatus = "uncleared"
	// ClearingStatusCleared identifies a cleared transaction
	ClearingStatusCleared ClearingStatus = "cleared"
	// ClearingStatusReconciled identifies a reconciled transaction
	ClearingStatusReconciled ClearingStatus = "reconciled"
)

type FlagColor string

const (
	// FlagColorRed identifies a transaction flagged red
	FlagColorRed FlagColor = "red"
	// FlagColorOrange identifies a transaction flagged orange
	FlagColorOrange FlagColor = "orange"
	// FlagColorYellow identifies a transaction flagged yellow
	FlagColorYellow FlagColor = "yellow"
	// FlagColorGreen identifies a transaction flagged green
	FlagColorGreen FlagColor = "green"
	// FlagColorBlue identifies a transaction flagged blue
	FlagColorBlue FlagColor = "blue"
	// FlagColorPurple identifies a transaction flagged purple
	FlagColorPurple FlagColor = "purple"
)

type Type string

const (
	// TypeTransaction identifies a hybrid transaction as transaction
	TypeTransaction Type = "transaction"
	// TypeSubTransaction identifies a hybrid transaction as sub-transaction
	TypeSubTransaction Type = "subtransaction"
)
