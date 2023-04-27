package request

type DebitCardRequest struct {
	CardNumb    string `json:"nomorKartu"`
	Valid       string `json:"validUntil"`
	SecurityNum string `json:"securityNum"`
}
