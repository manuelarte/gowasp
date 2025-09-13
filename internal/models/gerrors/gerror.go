package gerrors

var _ error = new(PasswordNotValidError)

type PasswordNotValidError struct {
	Message string
}

func (p PasswordNotValidError) Error() string {
	return p.Message
}
