package structs_helpers

import (
	"go-api/logs"

	"github.com/go-playground/validator/v10"
)

// StructValidator is an instance of a validator used for validating structures.
var StructValidator = validator.New()

// ValidateStruct validates the `s` structure using the `v` validator.
// Returns true if the structure is valid, otherwise returns false.
func ValidateStruct(v *validator.Validate, s interface{}) bool {
	// Validate the `s` structure using the `v` validator
	err := v.Struct(s)
	if err != nil {
		// If there are validation errors, log each error and return false
		for _, err := range err.(validator.ValidationErrors) {
			logs.Warning_Logger.Println(err.StructField(), err.Tag())
		}
		return false
	} else {
		// If there are no validation errors, return true
		return true
	}
}
