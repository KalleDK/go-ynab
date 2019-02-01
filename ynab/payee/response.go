package payee

type Response struct {
	Data Wrapper `json:"data"`
}

type ResponseList struct {
	Data WrapperList `json:"data"`
}
