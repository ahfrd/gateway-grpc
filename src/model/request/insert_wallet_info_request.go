package request

type InsertWalletInfoRequest struct {
	PhoneNumber  string `json:"phoneNumber"`
	SecurityCode string `json:"securityCode"`
}
