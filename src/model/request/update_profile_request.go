package request

type UpdateProfileRequest struct {
	PhoneNumber string `json:"phoneNumber"`
	Email       string `json:"email"`
	Name        string `json:"name"`
}
