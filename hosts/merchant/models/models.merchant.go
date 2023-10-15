package models

type InquiryItems struct{
	ID				int			`json:"id"`
	Name			string		`json:"name"`
	Type			string		`json:"type"`
	Price			int			`json:"price"`
	Quantity		int			`json:"quantity"`
}
type InquiryDiscounts struct{
	ID				int			`json:"id"`
	Name			string		`json:"name"`
	Type			string		`json:"type"`
	Percentage		int			`json:"percentage"`
}

type ResponseMerchantItems struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Data    []InquiryItems
}
type ResponseMerchantDiscounts struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Data    []InquiryDiscounts
}

type TransactionItems struct{
	ItemID			int			`json:"item_id"`
	Discount		string		`json:"discount"`
	Quantity		int			`json:"quantity"`
	CCNumber		string		`json:"cc_number"`
	CVV				string		`json:"cvv"`
	Amount			int			`json:"amount"`
}

type ResponseItems struct{
	ID			int			`json:"item_id"`
	Name		string		`json:"item_name"`
	Quantity	int			`json:"quantity"`
	CC			string		`json:"cc_number"`
	Code		[]string	`json:"code"`
}

type ResponseTransactionItems struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Data    ResponseItems
}