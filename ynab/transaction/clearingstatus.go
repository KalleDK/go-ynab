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
