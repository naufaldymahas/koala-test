package payload

// for sex true = male, false = female
type RegisterRequest struct {
	CustomerName string `json:"customer_name" validate:"required"`
	Email        string `json:"email" validate:"required,email"`
	PhoneNumber  string `json:"phone_number" validate:"required,number,min=9,max=13"`
	DOB          string `json:"dob" validate:"required"`
	Sex          bool   `json:"sex"`
	Password     string `json:"password" validate:"required"`
}
