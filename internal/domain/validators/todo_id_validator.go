package validators

func ValidateTodoId(ID int) ValidationError {
	if ID <= 0 {
		return InvalidNumber()
	}
	return ValidationError{}
}
