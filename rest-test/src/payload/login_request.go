package payload

type LoginRequest struct {
	PhoneNumberOrEmail string `json:"phone_number_or_email" validate:"required"`
	Password           string `json:"password" validate:"required"`
}
