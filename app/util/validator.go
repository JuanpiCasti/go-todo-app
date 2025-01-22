package util

type Validator interface {
	Validate(text string) *ValidationError
	AddValidations(validations ...Validation)
}

type Validation interface {
	Validate(text string) *ValidationError
}

type validator struct {
	validations []Validation
}

type ValidationError struct {
	Message string
}

func (e *ValidationError) Error() string {
	return e.Message
}

func NewValidator() Validator {
	return &validator{}
}

func NewPasswordValidator() Validator {
	v := NewValidator()
	v.AddValidations(&MinLengthValidator{min: 8}, &HasNumberValidator{})
	return v
}

func (v *validator) Validate(text string) *ValidationError {
	for _, validation := range v.validations {
		if err := validation.Validate(text); err != nil {
			return err
		}
	}
	return nil
}

func (v *validator) AddValidations(validations ...Validation) {
	for _, validation := range validations {
		v.validations = append(v.validations, validation)
	}
}

type MinLengthValidator struct {
	min int
}

func (v *MinLengthValidator) Validate(text string) *ValidationError {
	if len(text) < v.min {
		return &ValidationError{Message: "text is too short"}
	}
	return nil
}

type HasNumberValidator struct{}

func (v *HasNumberValidator) Validate(text string) *ValidationError {
	for _, char := range text {
		if char >= '0' && char <= '9' {
			return nil
		}
	}
	return &ValidationError{Message: "text must contain at least one number"}
}
