package validators

const (
	descriptionMaxSize = 255
)

func ValidateDescription(title string) ValidationError {
	if len(title) > descriptionMaxSize {
		return FieldTooLong(descriptionMaxSize)
	}
	return ValidationError{}
}
