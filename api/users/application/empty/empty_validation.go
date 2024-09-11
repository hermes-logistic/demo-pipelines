package users_empty_application

import (
	users_domain "go-api/api/users/domain"
	structs_helpers "go-api/helpers/structs"
)

func EmptyValidation(u *users_domain.Users) bool {
	newUser := users_domain.Users{
		Username: u.Username,
		Password: u.Password,
	}

	validate := structs_helpers.StructValidator

	response := structs_helpers.ValidateStruct(validate, newUser)

	if !response {
		return false
	} else {
		return true
	}
}
