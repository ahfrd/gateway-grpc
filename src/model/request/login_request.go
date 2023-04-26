package request

type LoginRequestBody struct {
	PhoneNumber string `json:"phoneNumber"`
	Password    string `json:"password"`
}
