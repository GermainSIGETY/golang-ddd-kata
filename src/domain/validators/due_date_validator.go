package validators

func ValidateDueDate(dueDatAsLong *int64) ValidationError {
	switch {
	case dueDatAsLong == nil:
		return EmptyField()
	case *dueDatAsLong < 0:
		return InvalidNumber()

	}
	return ValidationError{}
}
