package validators

const (
	titleMaxSize = 255
)

func ValidateTitle(title string) ValidationError {
	switch {
	case title == "":
		return EmptyField()
	case len(title) > titleMaxSize:
		return FieldTooLong(titleMaxSize)
	}
	return ValidationError{}
}
