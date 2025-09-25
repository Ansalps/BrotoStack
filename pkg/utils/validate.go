package utils

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/go-playground/validator/v10"
)

func UsernamRegex(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`^[a-zA-z0-9_.]+$`)
	return re.MatchString(fl.Field().String())
}
func NoRepeatedPeriods(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`\.\.`)
	return !re.MatchString(fl.Field().String())
}

func Validate(s any) error{
	fmt.Println("just a check to check whether it is entering in this function")
	validate := validator.New(validator.WithRequiredStructEnabled())
	validate.RegisterValidation("noRepeatedPeriods", NoRepeatedPeriods)
	validate.RegisterValidation("usernameRegex", UsernamRegex)
	if err := validate.Struct(s); err != nil {
		validationErrors, ok := err.(validator.ValidationErrors)
		if !ok {
			fmt.Println("Error type assertion failed!")
			return errors.New("error on type assertion of error to validationError")
		}
		for _, validationErr := range validationErrors {
			switch validationErr.Tag() {
			case "required":
				return fmt.Errorf("The field '%s' is required ", validationErr.Field())
			case "min=1":
				return fmt.Errorf("The field '%s' should have a minimum length of '%s' ", validationErr.Field(),validationErr.Param())
			case "max=30":
				return fmt.Errorf("The field '%s' should only have a maximum length of '%s' ", validationErr.Field(),validationErr.Param())
			case "usernamRegex":
				return fmt.Errorf("The field '%s' should only include alphabets,numbers,underscore or periods ", validationErr.Field())
			case "startsnotwith":
				return fmt.Errorf("The field '%s' should not start with '%s' ", validationErr.Field(),validationErr.Param())
			case "endsnotwith=.":
				return fmt.Errorf("The field '%s' should not end with '%s' ", validationErr.Field(),validationErr.Param())
			case "noRepeatedPeriods":
				return fmt.Errorf("The field '%s' should not have repeated '%s' ", validationErr.Field(),validationErr.Param())
			case "email":
				return fmt.Errorf("The field '%s' is not a valid email ", validationErr.Field())
			case "eqfield":
				return fmt.Errorf("The field '%s' should be same as '%s' field ", validationErr.Field(),validationErr.Param())
			
			default:
				// Default error message for other validation failures
				fmt.Printf("Validation failed for '%s': %s\n", validationErr.Field(), validationErr.Error())
			}
		}
	}
	fmt.Println("is it reaching here?")
	return  nil

}
