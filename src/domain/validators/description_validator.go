package validators

const (
	descriptionMaxSize = 255
)

func ValidateDescription(title *string) ValidationError {
	if title != nil && len(*title) > descriptionMaxSize {
		return FieldTooLong(descriptionMaxSize)
	}
	return ValidationError{}
}
