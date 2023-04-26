package request

type RegisterRequestBody struct {
	PhoneNumber string `json:"phoneNumber"`
	Password    string `json:"password"`
}
