package contracts

type ErrorResponse struct {
	Message       string      `json:"message"`
	Code          int         `json:"code"`
	Data          interface{} `json:"data"`
	TransactionId string      `json:"transactionId"`
}
