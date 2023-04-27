package request

type InquiryRequest struct {
	Method          string           `json:"method"`
	Nominal         string           `json:"nominalTopUp"`
	Fee             string           `json:"feeTopUp"`
	TransactionCode string           `json:"transactionCode"`
	PhoneNumb       string           `json:"phoneNumber"`
	Card            DebitCardRequest `json:"cardInfo"`
	CodeNumb        string           `json:"codeNumber"`
	BankCode        string           `json:"bankCode"`
}
