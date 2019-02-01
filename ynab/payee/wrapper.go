package payee

type Wrapper struct {
	Payee Payee `json:"payee"`
}

type WrapperList struct {
	Payees Payees `json:"payees"`
}
