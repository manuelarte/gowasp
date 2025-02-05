package errors

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

var _ error = new(PasswordNotValid)

type PasswordNotValid struct {
	Message string
}

func (p PasswordNotValid) Error() string {
	return p.Message
}
